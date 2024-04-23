package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/menezmethod/st-server/src/st-gateway/internal/auth"
	"github.com/menezmethod/st-server/src/st-gateway/internal/interceptor"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/config"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/middleware"
	pbAuth "github.com/menezmethod/st-server/src/st-gateway/pkg/pb/auth"
	pbHelper "github.com/menezmethod/st-server/src/st-gateway/pkg/pb/helper"
	pbJournal "github.com/menezmethod/st-server/src/st-gateway/pkg/pb/journal"
	pbRecord "github.com/menezmethod/st-server/src/st-gateway/pkg/pb/record"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/utils"
)

func main() {
	logger := initLogger()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("Failed to sync zap logger", zap.Error(err))
		}
	}(logger)

	cfg := loadConfig(logger)
	authServiceClient := initAuthServiceClient(cfg, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handler := createGatewayMux(ctx, cfg, logger, authServiceClient)
	addPrometheusHandler(handler)

	startHTTPServer(cfg, logger)
}

func addPrometheusHandler(handler http.Handler) {
	http.Handle(utils.MetricsPath, promhttp.Handler())
	http.Handle(utils.DefaultRoute, handler)
}

func createGatewayMux(ctx context.Context, cfg *config.Config, logger *zap.Logger, authServiceClient pbAuth.AuthServiceClient) http.Handler {
	conn, err := grpc.DialContext(ctx, cfg.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(utils.LogFailedToDialService, zap.String("url", cfg.AuthSvcUrl), zap.Error(err))
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			logger.Error(utils.LogFailedToCloseConnection, zap.Error(err))
		}
	}(conn)

	mux := runtime.NewServeMux(
		runtime.WithMetadata(interceptor.AuthInterceptor(createAuthFunc(authServiceClient), logger)),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	registerServiceHandlers(ctx, mux, cfg, opts, logger)

	allowedOrigins := []string{utils.AllowedOrigins}
	corsMiddleware := middleware.CORSMiddleware(allowedOrigins)
	authMiddleware := middleware.AuthMiddleware(authServiceClient)

	handler := corsMiddleware(authMiddleware(mux))

	return handler
}

func initLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}
	return logger
}

func loadConfig(logger *zap.Logger) *config.Config {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}
	return &cfg
}

func initAuthServiceClient(cfg *config.Config, logger *zap.Logger) pbAuth.AuthServiceClient {
	authServiceClient, err := auth.InitServiceClient(cfg, logger)
	if err != nil {
		logger.Fatal("Failed to initialize auth service client", zap.Error(err))
	}
	return authServiceClient
}

func createAuthFunc(authServiceClient pbAuth.AuthServiceClient) func(string) (string, error) {
	return func(token string) (string, error) {
		res, err := authServiceClient.Validate(context.Background(), &pbAuth.ValidateRequest{Token: token})
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "AuthService.Validate: %v", err)
		}
		return strconv.FormatUint(res.UserId, 10), nil
	}
}

func registerServiceHandlers(ctx context.Context, mux *runtime.ServeMux, cfg *config.Config, opts []grpc.DialOption, logger *zap.Logger) {
	registerHandler(ctx, mux, cfg.AuthSvcUrl, opts, pbAuth.RegisterAuthServiceHandlerFromEndpoint, logger)
	registerHandler(ctx, mux, cfg.JournalSvcUrl, opts, pbJournal.RegisterJournalServiceHandlerFromEndpoint, logger)
	registerHandler(ctx, mux, cfg.HelperSvcUrl, opts, pbHelper.RegisterSTHelperHandlerFromEndpoint, logger)
	registerHandler(ctx, mux, cfg.JournalSvcUrl, opts, pbRecord.RegisterRecordServiceHandlerFromEndpoint, logger)
}

func registerHandler(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption, registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error, logger *zap.Logger) {
	if err := registerFunc(ctx, mux, endpoint, opts); err != nil {
		logger.Fatal(utils.LogFailedToRegister, zap.Error(err))
	}
}

func startHTTPServer(cfg *config.Config, logger *zap.Logger) {
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		logger.Fatal(utils.LogFailedToStartServer, zap.String("port", cfg.Port), zap.Error(err))
	}
}

package main

import (
	"log"
	"net"
	"net/http"

	"github.com/go-playground/validator/v10"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/menezmethod/st-server/src/st-auth-svc/configs"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/db"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/services"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/utils"
)

func main() {
	logger := initLogger()
	defer logger.Sync()

	c := loadConfig(logger)

	dbHandler := db.InitDB(c.DBUrl, logger)

	reg := initPrometheus()

	lis := initListener(c.Port, logger)
	logger.Info("Auth service listening on", zap.String("port", c.Port))

	grpcServer := initGRPCServer(dbHandler, logger, c)

	startPrometheusServer(reg)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", zap.Error(err))
	}
}

func initLogger() *zap.Logger {
	logger, err := zap.NewDevelopment() // Change to NewProduction in Prod
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	return logger
}

func loadConfig(logger *zap.Logger) configs.Config {
	c, err := configs.LoadConfig()
	if err != nil {
		logger.Fatal("Failed loading configs", zap.Error(err))
	}
	return c
}

func initPrometheus() *prometheus.Registry {
	reg := prometheus.NewRegistry()
	grpcMetrics := grpcprometheus.NewServerMetrics()
	reg.MustRegister(grpcMetrics)
	return reg
}

func initListener(port string, logger *zap.Logger) net.Listener {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("Failed to listen", zap.String("port", port), zap.Error(err))
	}
	return lis
}

func initGRPCServer(dbHandler db.DB, logger *zap.Logger, c configs.Config) *grpc.Server {
	grpcMetrics := grpcprometheus.NewServerMetrics()
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	grpcMetrics.InitializeMetrics(grpcServer)

	pb.RegisterAuthServiceServer(grpcServer, &services.Server{
		H:         dbHandler,
		Logger:    logger,
		Jwt:       initJwt(c),
		Validator: validator.New(),
	})

	return grpcServer
}

func initJwt(c configs.Config) utils.JwtWrapper {
	return utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "st-auth-svc",
		ExpirationHours: 24 * 365,
	}
}

func startPrometheusServer(reg *prometheus.Registry) {
	go func() {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		if err := http.ListenAndServe(":9093", nil); err != nil {
			log.Printf("can't initialize prometheus metrics port: %v\n", zap.Error(err))
		}
	}()
}

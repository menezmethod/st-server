package auth

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/menezmethod/st-server/src/st-gateway/pkg/config"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/pb/auth"
)

type ServiceClient struct {
	AuthServiceClient auth.AuthServiceClient
}

type Server struct {
	auth.UnimplementedAuthServiceServer
}

// InitServiceClient initializes a client for the AuthService using the provided configuration.
// It logs the initialization attempt and returns an error if the connection cannot be established.
func InitServiceClient(cfg *config.Config, logger *zap.Logger) (auth.AuthServiceClient, error) {
	var cc *grpc.ClientConn
	var err error

	logger.Info("Initializing gRPC service client for Auth service", zap.String("URL", cfg.AuthSvcUrl))

	for i := 0; i < 12; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cc, err = grpc.DialContext(ctx, cfg.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		cancel()

		if err == nil {
			logger.Info("Successfully connected to Auth service")
			return auth.NewAuthServiceClient(cc), nil
		}

		logger.Error("Failed to connect to Auth service, retrying...", zap.Int("attempt", i+1), zap.Error(err))

		time.Sleep(5 * time.Second)
	}

	logger.Error("Failed to connect to Auth service after retries", zap.Error(err))
	return nil, err
}

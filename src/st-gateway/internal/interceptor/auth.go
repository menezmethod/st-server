package interceptor

import (
	"context"
	"net/http"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

// AuthFunc is a function that authenticates the provided token and returns the user ID
type AuthFunc func(token string) (userID string, err error)

// AuthInterceptor creates a new HTTP request modifier for authentication
func AuthInterceptor(authFunc AuthFunc, logger *zap.Logger) func(ctx context.Context, req *http.Request) metadata.MD {
	return func(ctx context.Context, req *http.Request) metadata.MD {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			logger.Warn("No authorization header in request")
			return nil
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			logger.Warn("Invalid authorization header format")
			return nil
		}

		userID, err := authFunc(token)
		if err != nil {
			logger.Error("Failed to validate token", zap.Error(err))
			return nil
		}

		logger.Info("Token validated and user ID extracted", zap.String("userID", userID))
		return metadata.Pairs("user-id", userID)
	}
}

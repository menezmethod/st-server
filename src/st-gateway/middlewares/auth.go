package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/auth/pb"
	"strings"
)

type MiddlewareConfig struct {
	svc *auth.ServiceClient
}

func InitAuthMiddleware(svc *auth.ServiceClient) MiddlewareConfig {
	return MiddlewareConfig{svc}
}

func (c *MiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		abortWithStatusMessage(ctx, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		abortWithStatusMessage(ctx, http.StatusUnauthorized, "Authorization header must start with 'Bearer '")
		return
	}

	token := strings.TrimPrefix(authHeader, prefix)
	if token == "" {
		abortWithStatusMessage(ctx, http.StatusUnauthorized, "Token not provided")
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: wrapperspb.String(token), // Correctly converting string to *wrapperspb.StringValue
	})

	if err != nil || res.GetStatus() != http.StatusOK {
		abortWithStatusMessage(ctx, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	ctx.Set("userId", res.GetUserId())
	ctx.Next()
}

func abortWithStatusMessage(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}

package auth

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
)

type MiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) MiddlewareConfig {
	if svc == nil || svc.AuthServiceClient == nil {
		return MiddlewareConfig{}
	}
	return MiddlewareConfig{svc}
}
func (c *MiddlewareConfig) AuthRequired(ctx *gin.Context) {
	if c.svc == nil || c.svc.AuthServiceClient == nil {
		log.Println("ServiceClient or AuthServiceClient is nil")
		respondWithError(ctx, http.StatusInternalServerError, "Internal server error", "service configuration issue")
		return
	}

	authHeaderParts := strings.Split(ctx.Request.Header.Get("authorization"), "Bearer ")
	if len(authHeaderParts) < 2 {
		respondWithError(ctx, http.StatusUnauthorized, "Bearer token not provided", "unauthorized")
		return
	}

	token := authHeaderParts[1]
	res, err := c.svc.AuthServiceClient.Validate(ctx.Request.Context(), &pb.ValidateRequest{
		Token: token,
	})

	if err != nil || res.Status != http.StatusOK {
		errorMessage := "Invalid or expired token"
		if err != nil {
			errorMessage = err.Error()
		}
		respondWithError(ctx, http.StatusUnauthorized, "Failed to authenticate", errorMessage)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}

func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		var isAllowed bool
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func respondWithError(ctx *gin.Context, statusCode int, message, detail string) {
	errorResponse := gin.H{
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"level":     "ERROR",
		"message":   message,
		"error":     detail,
	}

	ctx.AbortWithStatusJSON(statusCode, errorResponse)
	ctx.Abort()
}

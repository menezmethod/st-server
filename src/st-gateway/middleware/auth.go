package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/auth"
	"strings"

	"st-gateway/pkg/auth/pb"
)

type Config struct {
	svc *auth.ServiceClient
}

func InitAuthMiddleware(svc *auth.ServiceClient) Config {
	return Config{svc}
}

func (c *Config) AuthRequired(ctx *gin.Context) {
	if ctx.Request.Header.Get("authorization") == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if len(strings.Split(ctx.Request.Header.Get("authorization"), "Bearer ")) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: strings.Split(ctx.Request.Header.Get("authorization"), "Bearer ")[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}

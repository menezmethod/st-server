package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"strings"

	"st-gateway/pkg/auth/pb"
)

type MiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) MiddlewareConfig {
	return MiddlewareConfig{svc}
}

func (c *MiddlewareConfig) AuthRequired(ctx *gin.Context) {
	if ctx.Request.Header.Get("authorization") == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if len(strings.Split(ctx.Request.Header.Get("authorization"), "Bearer ")) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: wrapperspb.String(strings.Split(ctx.Request.Header.Get("authorization"), "Bearer ")[1]),
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}

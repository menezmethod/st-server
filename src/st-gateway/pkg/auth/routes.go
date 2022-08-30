package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth/routes"

	"st-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerGroup := r.Group(fmt.Sprintf("v%v/auth", c.ApiVersion))
	routerGroup.POST("/register", svc.Register)
	routerGroup.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

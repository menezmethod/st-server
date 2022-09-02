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

	routerGroup := r.Group(fmt.Sprintf("v%v/", c.ApiVersion))
	routerGroup.POST("/auth/register", svc.Register)
	routerGroup.POST("/auth/login", svc.Login)
	//routerGroup.GET("/trader/:id", svc.Login)
	routerGroup.PATCH("/trader/:id", svc.UpdateUser)
	routerGroup.DELETE("/trader/:id", svc.DeleteUser)

	return svc
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) {
	routes.DeleteUser(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, svc.Client)
}

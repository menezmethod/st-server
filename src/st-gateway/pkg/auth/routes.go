package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth/routes"

	"st-gateway/pkg/config"
)

func RegisterAuthRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	routerGroup := r.Group(fmt.Sprintf("v%v/", c.ApiVersion))
	routerGroup.POST("/auth/register", svc.Register)
	routerGroup.POST("/auth/login", svc.Login)
	routerGroup.GET("/auth/me", svc.FindMe)
	routerGroup.GET("/users", svc.FindAllUsers)
	routerGroup.GET("/users/:id", svc.FindOneUser)
	routerGroup.PATCH("/users/:id", svc.UpdateUser)
	routerGroup.DELETE("/user/:id", svc.DeleteUser)
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

func (svc *ServiceClient) FindAllUsers(ctx *gin.Context) {
	routes.FindAllUsers(ctx, svc.Client)
}

func (svc *ServiceClient) FindOneUser(ctx *gin.Context) {
	routes.FindOneUser(ctx, svc.Client)
}

func (svc *ServiceClient) FindMe(ctx *gin.Context) {
	routes.Me(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateUser(ctx *gin.Context) {
	routes.UpdateUser(ctx, svc.Client)
}

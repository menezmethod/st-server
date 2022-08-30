package trade

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/trade/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group(fmt.Sprintf("v%v/trade", c.ApiVersion))
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateTrade)
	routes.GET("/:id", svc.FindOne)
	routes.DELETE("/:id", svc.Delete)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateTrade(ctx *gin.Context) {
	routes.CreateTrade(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}

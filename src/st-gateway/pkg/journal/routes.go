package journal

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerGroup := r.Group(fmt.Sprintf("v%v", c.ApiVersion))
	routerGroup.Use(cors.Default())
	routerGroup.Use(a.AuthRequired)
	//routerGroup.GET("/journals/", svc.AllJournals)
	routerGroup.POST("/journal/", svc.CreateJournal)
	routerGroup.PATCH("/journal/:id", svc.EditJournal)
	routerGroup.GET("/journal/:id", svc.FindOneJournal)
	routerGroup.DELETE("/journal/:id", svc.DeleteJournal)
	//routerGroup.GET("/trades/", svc.AllTrades)
	routerGroup.POST("/trade/", svc.CreateTrade)
	routerGroup.PATCH("/trade/:id", svc.EditTrade)
	routerGroup.GET("/trade/:id", svc.FindOneTrade)
	routerGroup.DELETE("/trade/:id", svc.DeleteTrade)
}

func (svc *ServiceClient) FindOneJournal(ctx *gin.Context) {
	routes.FineOneJournal(ctx, svc.Client)
}

func (svc *ServiceClient) CreateJournal(ctx *gin.Context) {
	routes.CreateJournal(ctx, svc.Client)
}

func (svc *ServiceClient) EditJournal(ctx *gin.Context) {
	routes.EditJournal(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteJournal(ctx *gin.Context) {
	routes.DeleteJournal(ctx, svc.Client)
}

func (svc *ServiceClient) FindOneTrade(ctx *gin.Context) {
	routes.FineOneTrade(ctx, svc.Client)
}

func (svc *ServiceClient) CreateTrade(ctx *gin.Context) {
	routes.CreateTrade(ctx, svc.Client)
}

func (svc *ServiceClient) EditTrade(ctx *gin.Context) {
	routes.EditTrade(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteTrade(ctx *gin.Context) {
	routes.DeleteTrade(ctx, svc.Client)
}

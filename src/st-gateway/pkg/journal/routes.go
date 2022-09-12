package journal

import (
	"fmt"
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

	routes := r.Group(fmt.Sprintf("v%v", c.ApiVersion))
	routes.Use(a.AuthRequired)
	//routes.GET("/journals/", svc.AllJournals)
	routes.POST("/journal/", svc.CreateJournal)
	routes.PATCH("/journal/:id", svc.EditJournal)
	routes.GET("/journal/:id", svc.FindOneJournal)
	routes.DELETE("/journal/:id", svc.DeleteJournal)
	//routes.GET("/trades/", svc.AllTrades)
	routes.POST("/trade/", svc.CreateTrade)
	routes.PATCH("/trade/:id", svc.EditTrade)
	routes.GET("/trade/:id", svc.FindOneTrade)
	routes.DELETE("/trade/:id", svc.DeleteTrade)
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

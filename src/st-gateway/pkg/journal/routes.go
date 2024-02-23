package journal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/middlewares"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/config"
	"st-gateway/pkg/journal/routes"
)

func RegisterJournalRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := middlewares.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routerJournal := r.Group(fmt.Sprintf("v%v", c.ApiVersion))
	routerJournal.GET("/journals", svc.FindAllJournals)
	routerJournal.POST("/journal/", svc.CreateJournal)
	routerJournal.PATCH("/journal/:id", svc.EditJournal)
	routerJournal.GET("/journal/:id", svc.FindOneJournal)
	routerJournal.DELETE("/journal/:id", svc.DeleteJournal)
	routerJournal.GET("/trades/", svc.FindAllTrades)
	routerJournal.POST("/trade/", svc.CreateTrade)
	routerJournal.PATCH("/trade/:id", svc.EditTrade)
	routerJournal.GET("/trade/:id", svc.FindOneTrade)
	routerJournal.DELETE("/trade/:id", svc.DeleteTrade)
	routerJournal.Use(a.AuthRequired)
	return svc
}

func (svc *ServiceClient) FindAllJournals(ctx *gin.Context) {
	routes.FindAllJournals(ctx, svc.Client)
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

func (svc *ServiceClient) FindAllTrades(ctx *gin.Context) {
	routes.FindAllTrades(ctx, svc.Client)
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

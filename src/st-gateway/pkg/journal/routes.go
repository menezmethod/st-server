package journal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/configs"
	"st-gateway/middleware"
	"st-gateway/pkg/auth"
	"st-gateway/pkg/journal/routes"
)

func RegisterJournalRoutes(r *gin.Engine, c *configs.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := middleware.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		JournalServiceClient: InitJournalServiceClient(c),
		TradeServiceClient:   InitTradeServiceClient(c),
	}

	routerJournal := r.Group(fmt.Sprintf("v%v", c.ApiVersion))
	routerJournal.GET("/journals", svc.FindAllJournals)
	routerJournal.POST("/journal/", svc.CreateJournal)
	routerJournal.PATCH("/journal/:id", svc.UpdateJournal)
	routerJournal.GET("/journal/:id", svc.FindOneJournal)
	routerJournal.DELETE("/journal/:id", svc.DeleteJournal)
	routerJournal.GET("/trades/", svc.FindAllTrades)
	routerJournal.POST("/trade/", svc.CreateTrade)
	routerJournal.PATCH("/trade/:id", svc.UpdateTrade)
	routerJournal.GET("/trade/:id", svc.FindOneTrade)
	routerJournal.DELETE("/trade/:id", svc.DeleteTrade)
	routerJournal.Use(a.AuthRequired)
	return svc
}

func (svc *ServiceClient) FindAllJournals(ctx *gin.Context) {
	routes.FindAllJournals(ctx, svc.JournalServiceClient)
}
func (svc *ServiceClient) FindOneJournal(ctx *gin.Context) {
	routes.FineOneJournal(ctx, svc.JournalServiceClient)
}

func (svc *ServiceClient) CreateJournal(ctx *gin.Context) {
	routes.CreateJournal(ctx, svc.JournalServiceClient)
}

func (svc *ServiceClient) UpdateJournal(ctx *gin.Context) {
	routes.UpdateJournal(ctx, svc.JournalServiceClient)
}

func (svc *ServiceClient) DeleteJournal(ctx *gin.Context) {
	routes.DeleteJournal(ctx, svc.JournalServiceClient)
}

func (svc *ServiceClient) FindAllTrades(ctx *gin.Context) {
	routes.FindAllTrades(ctx, svc.TradeServiceClient)
}

func (svc *ServiceClient) FindOneTrade(ctx *gin.Context) {
	routes.FineOneTrade(ctx, svc.TradeServiceClient)
}

func (svc *ServiceClient) CreateTrade(ctx *gin.Context) {
	routes.CreateTrade(ctx, svc.TradeServiceClient)
}

func (svc *ServiceClient) UpdateTrade(ctx *gin.Context) {
	routes.UpdateTrade(ctx, svc.TradeServiceClient)
}

func (svc *ServiceClient) DeleteTrade(ctx *gin.Context) {
	routes.DeleteTrade(ctx, svc.TradeServiceClient)
}

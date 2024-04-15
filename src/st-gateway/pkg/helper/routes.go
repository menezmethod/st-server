package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper/routes"
)

type endpoint struct {
	method, path string
	handler      func(ctx *gin.Context)
}

func RegisterHelperRoutes(r *gin.Engine, c *configs.Config) *ServiceClient {
	svc := &ServiceClient{
		STHelperClient: StHelperClient(c),
	}

	endpoints := []endpoint{
		{"POST", "/helper/analyze", svc.AnalyzeFinancialData},
		{"POST", "/helper/stock-data", svc.GetStockQuote},
		{"POST", "/helper/stock-hd", svc.GetHistoricalStockData},
	}

	group := r.Group(fmt.Sprintf("v%v/", c.ApiVersion))

	helperRoutes := group.Group("/")
	for _, e := range endpoints {
		switch e.method {
		case "POST":
			helperRoutes.POST(e.path, e.handler)
		}
	}

	return svc
}

func (svc *ServiceClient) AnalyzeFinancialData(ctx *gin.Context) {
	routes.AnalyzeFinancialData(ctx, svc.STHelperClient)
}

func (svc *ServiceClient) GetStockQuote(ctx *gin.Context) {
	routes.GetStockQuote(ctx, svc.STHelperClient)
}

func (svc *ServiceClient) GetHistoricalStockData(ctx *gin.Context) {
	routes.GetHistoricalStockData(ctx, svc.STHelperClient)
}

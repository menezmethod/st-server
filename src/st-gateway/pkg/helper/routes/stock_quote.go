package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper/pb"
)

type GetStockQuoteRequestBody struct {
	Ticker   string `json:"ticker"`
	Provider string `json:"provider"`
}

func GetStockQuote(ctx *gin.Context, c pb.STHelperClient) {
	b := GetStockQuoteRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	res, err := c.GetStockQuote(timeoutCtx, &pb.StockQuoteRequest{
		Ticker:   b.Ticker,
		Provider: b.Provider,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred", "details": err.Error()})
		return
	}

	stockData := map[string]interface{}{
		"stock_quote": map[string]interface{}{
			"symbol":               res.Symbol,
			"name":                 res.Name,
			"exchange":             res.Exchange,
			"lastPrice":            res.LastPrice,
			"open":                 res.Open,
			"high":                 res.High,
			"low":                  res.Low,
			"volume":               res.Volume,
			"prevClose":            res.PrevClose,
			"change":               res.Change,
			"changePercent":        res.ChangePercent,
			"yearHigh":             res.YearHigh,
			"yearLow":              res.YearLow,
			"marketCap":            res.MarketCap,
			"sharesOutstanding":    res.SharesOutstanding,
			"pe":                   res.Pe,
			"earningsAnnouncement": res.EarningsAnnouncement,
			"eps":                  res.Eps,
			"sector":               res.Sector,
			"industry":             res.Industry,
			"beta":                 res.Beta,
		},
	}

	ctx.JSON(http.StatusOK, stockData)
}

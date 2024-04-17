package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper/pb"
)

type GetHistoricalDataRequestBody struct {
	Symbol    string `json:"symbol"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Provider  string `json:"provider"`
}

func GetHistoricalStockData(ctx *gin.Context, c pb.STHelperClient) {
	b := GetHistoricalDataRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	res, err := c.GetHistoricalStockData(timeoutCtx, &pb.HistoricalStockDataRequest{
		Symbol:    b.Symbol,
		StartDate: b.StartDate,
		EndDate:   b.EndDate,
		Provider:  b.Provider,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred", "details": err.Error()})
		return
	}

	historicalData := map[string]interface{}{
		"historical_data": map[string]interface{}{
			"open":              res.Open,
			"high":              res.High,
			"low":               res.Low,
			"close":             res.Close,
			"volume":            res.Volume,
			"vwap":              res.Vwap,
			"adj_close":         res.AdjClose,
			"unadjusted_volume": res.UnadjustedVolume,
			"change":            res.Change,
			"change_percent":    res.ChangePercent,
		},
	}

	ctx.JSON(http.StatusOK, historicalData)
}

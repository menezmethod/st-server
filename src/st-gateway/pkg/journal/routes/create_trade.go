package routes

import (
	"context"
	"net/http"
	"st-gateway/pkg/util"

	"github.com/gin-gonic/gin"
	"st-gateway/pkg/journal/pb"
)

type CreateTradeRequestBody struct {
	Comments        string  `json:"comments"`
	CreatedBy       string  `json:"createdBy"`
	Direction       string  `json:"direction"`
	EntryPrice      float32 `json:"entryPrice"`
	ExitPrice       float32 `json:"exitPrice"`
	Journal         uint64  `json:"journal"`
	BaseInstrument  string  `json:"baseInstrument"`
	QuoteInstrument string  `json:"quoteInstrument"`
	Market          string  `json:"market"`
	Outcome         string  `json:"outcome"`
	Quantity        float32 `json:"quantity"`
	StopLoss        float32 `json:"stopLoss"`
	Strategy        string  `json:"strategy"`
	TakeProfit      float32 `json:"takeProfit"`
	TimeClosed      string  `json:"timeClosed"`
	TimeExecuted    string  `json:"timeExecuted"`
}

func CreateTrade(ctx *gin.Context, c pb.TradeServiceClient) {
	b := CreateTradeRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if c == nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := c.CreateTrade(context.Background(), &pb.CreateTradeRequest{
		Comments:        b.Comments,
		CreatedBy:       b.CreatedBy,
		Direction:       b.Direction,
		EntryPrice:      b.EntryPrice,
		ExitPrice:       b.ExitPrice,
		Journal:         b.Journal,
		BaseInstrument:  b.BaseInstrument,
		QuoteInstrument: b.QuoteInstrument,
		Market:          b.Market,
		Outcome:         b.Outcome,
		Quantity:        b.Quantity,
		StopLoss:        b.StopLoss,
		Strategy:        b.Strategy,
		TakeProfit:      b.TakeProfit,
		TimeClosed:      b.TimeClosed,
		TimeExecuted:    b.TimeExecuted,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

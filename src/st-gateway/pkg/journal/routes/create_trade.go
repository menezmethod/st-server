package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
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

func CreateTrade(ctx *gin.Context, c pb.JournalServiceClient) {
	b := CreateTradeRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"

	"st-gateway/pkg/trade/pb"
)

type CreateTradeRequestBody struct {
	Comments     string    `json:"comments"`
	Direction    string    `json:"direction"`
	EntryPrice   uint64    `json:"entryPrice"`
	ExitPrice    uint64    `json:"exitPrice"`
	Instrument   string    `json:"instrument"`
	Market       string    `json:"market"`
	Outcome      string    `json:"outcome"`
	Quantity     uint32    `json:"quantity"`
	StopLoss     uint64    `json:"stopLoss"`
	Strategy     string    `json:"strategy"`
	TakeProfit   uint64    `json:"takeProfit"`
	TimeClosed   time.Time `json:"timeClosed"`
	TimeExecuted time.Time `json:"timeExecuted"`
}

func CreateTrade(ctx *gin.Context, c pb.TradeServiceClient) {
	b := CreateTradeRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateTrade(context.Background(), &pb.CreateTradeRequest{
		Comments:     b.Comments,
		Direction:    b.Direction,
		EntryPrice:   b.EntryPrice,
		ExitPrice:    b.ExitPrice,
		Instrument:   b.Instrument,
		Market:       b.Market,
		Outcome:      b.Outcome,
		Quantity:     b.Quantity,
		StopLoss:     b.StopLoss,
		Strategy:     b.Strategy,
		TakeProfit:   b.TakeProfit,
		TimeClosed:   timestamppb.New(b.TimeClosed),
		TimeExecuted: timestamppb.New(b.TimeExecuted),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

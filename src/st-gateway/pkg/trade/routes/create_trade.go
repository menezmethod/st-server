package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"time"

	"st-gateway/pkg/trade/pb"
)

type CreateTradeRequestBody struct {
	Comments     string    `json:"comments"`
	Direction    string    `json:"direction"`
	EntryPrice   float32   `json:"entryPrice"`
	ExitPrice    float32   `json:"exitPrice"`
	Instrument   string    `json:"instrument"`
	Market       string    `json:"market"`
	Outcome      string    `json:"outcome"`
	Quantity     float32   `json:"quantity"`
	StopLoss     float32   `json:"stopLoss"`
	Strategy     string    `json:"strategy"`
	TakeProfit   float32   `json:"takeProfit"`
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
		Comments:     wrapperspb.String(b.Comments),
		Direction:    wrapperspb.String(b.Direction),
		EntryPrice:   wrapperspb.Float(b.EntryPrice),
		ExitPrice:    wrapperspb.Float(b.ExitPrice),
		Instrument:   wrapperspb.String(b.Instrument),
		Market:       wrapperspb.String(b.Market),
		Outcome:      wrapperspb.String(b.Outcome),
		Quantity:     wrapperspb.Float(b.Quantity),
		StopLoss:     wrapperspb.Float(b.StopLoss),
		Strategy:     wrapperspb.String(b.Strategy),
		TakeProfit:   wrapperspb.Float(b.TakeProfit),
		TimeClosed:   timestamppb.New(b.TimeClosed),
		TimeExecuted: timestamppb.New(b.TimeExecuted),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

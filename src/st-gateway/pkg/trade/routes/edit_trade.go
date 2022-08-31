package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"st-gateway/pkg/trade/pb"
	"strconv"
	"time"
)

type EditTradeRequestBody struct {
	Id           uint64    `json:"id"`
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

func EditTrade(ctx *gin.Context, c pb.TradeServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	b := EditTradeRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EditTrade(context.Background(), &pb.EditTradeRequest{
		Id:           uint64(id),
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

package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"strconv"
	"time"
)

type EditTradeRequestBody struct {
	Id              uint64    `json:"id"`
	Comments        string    `json:"comments"`
	Direction       string    `json:"direction"`
	EntryPrice      float32   `json:"entryPrice"`
	ExitPrice       float32   `json:"exitPrice"`
	Journal         uint64    `json:"journal"`
	BaseInstrument  string    `json:"baseInstrument"`
	QuoteInstrument string    `json:"quoteInstrument"`
	Market          string    `json:"market"`
	Outcome         string    `json:"outcome"`
	Quantity        float32   `json:"quantity"`
	StopLoss        float32   `json:"stopLoss"`
	Strategy        string    `json:"strategy"`
	TakeProfit      float32   `json:"takeProfit"`
	TimeClosed      time.Time `json:"timeClosed"`
	TimeExecuted    time.Time `json:"timeExecuted"`
}

func EditTrade(ctx *gin.Context, c pb.JournalServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	b := EditTradeRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EditTrade(context.Background(), &pb.EditTradeRequest{
		Id:              uint64(id),
		Comments:        wrapperspb.String(b.Comments),
		Direction:       wrapperspb.String(b.Direction),
		EntryPrice:      wrapperspb.Float(b.EntryPrice),
		ExitPrice:       wrapperspb.Float(b.ExitPrice),
		Journal:         b.Journal,
		BaseInstrument:  wrapperspb.String(b.BaseInstrument),
		QuoteInstrument: wrapperspb.String(b.QuoteInstrument),
		Market:          wrapperspb.String(b.Market),
		Outcome:         wrapperspb.String(b.Outcome),
		Quantity:        wrapperspb.Float(b.Quantity),
		StopLoss:        wrapperspb.Float(b.StopLoss),
		Strategy:        wrapperspb.String(b.Strategy),
		TakeProfit:      wrapperspb.Float(b.TakeProfit),
		TimeClosed:      timestamppb.New(b.TimeClosed),
		TimeExecuted:    timestamppb.New(b.TimeExecuted),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

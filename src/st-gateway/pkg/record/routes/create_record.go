package routes

import (
	"context"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

type CreateRecordRequestBody struct {
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

func CreateRecord(ctx *gin.Context, c pb.RecordServiceClient) {
	b := CreateRecordRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if c == nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := c.CreateRecord(context.Background(), &pb.CreateRecordRequest{
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

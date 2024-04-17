package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

type UpdateRecordRequestBody struct {
	BaseInstrument  string  `json:"baseInstrument"`
	Comments        string  `json:"comments"`
	Direction       string  `json:"direction"`
	EntryPrice      float32 `json:"entryPrice"`
	ExitPrice       float32 `json:"exitPrice"`
	Id              uint64  `json:"id"`
	Journal         uint64  `json:"journal"`
	Market          string  `json:"market"`
	Outcome         string  `json:"outcome"`
	Quantity        float32 `json:"quantity"`
	QuoteInstrument string  `json:"quoteInstrument"`
	StopLoss        float32 `json:"stopLoss"`
	Strategy        string  `json:"strategy"`
	TakeProfit      float32 `json:"takeProfit"`
	TimeClosed      string  `json:"timeClosed"`
	TimeExecuted    string  `json:"timeExecuted"`
}

func UpdateRecord(ctx *gin.Context, c pb.RecordServiceClient, user *authPb.User) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	b := UpdateRecordRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.UpdateRecord(mdCtx, &pb.UpdateRecordRequest{
		BaseInstrument:  b.BaseInstrument,
		Comments:        b.Comments,
		Direction:       b.Direction,
		EntryPrice:      b.EntryPrice,
		ExitPrice:       b.ExitPrice,
		Id:              uint64(id),
		Journal:         b.Journal,
		LastUpdatedBy:   user.Id,
		Market:          b.Market,
		Outcome:         b.Outcome,
		Quantity:        b.Quantity,
		QuoteInstrument: b.QuoteInstrument,
		StopLoss:        b.StopLoss,
		Strategy:        b.Strategy,
		TakeProfit:      b.TakeProfit,
		TimeClosed:      b.TimeClosed,
		TimeExecuted:    b.TimeExecuted,
	})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"st-gateway/pkg/util"
	"strconv"
)

func FineOneTrade(ctx *gin.Context, c pb.TradeServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOneTrade(context.Background(), &pb.FindOneTradeRequest{
		Id: uint64(id),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

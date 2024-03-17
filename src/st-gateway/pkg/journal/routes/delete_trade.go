package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"st-gateway/pkg/util"
	"strings"
)

func DeleteTrade(ctx *gin.Context, c pb.TradeServiceClient) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	res, err := c.DeleteTrade(context.Background(), &pb.DeleteTradeRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

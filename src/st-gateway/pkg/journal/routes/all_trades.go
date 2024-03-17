package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"st-gateway/pkg/util"
)

func FindAllTrades(ctx *gin.Context, c pb.TradeServiceClient) {
	res, err := c.FindAllTrades(context.Background(), &pb.FindAllTradesRequest{})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

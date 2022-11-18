package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
)

func FindAllTrades(ctx *gin.Context, c pb.JournalServiceClient) {
	var id []uint64

	id = append(id, 1)

	res, err := c.FindAllTrades(context.Background(), &pb.FindAllTradesRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

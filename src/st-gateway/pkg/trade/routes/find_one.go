package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"st-gateway/pkg/trade/pb"
)

func FineOne(ctx *gin.Context, c pb.TradeServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: uint64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

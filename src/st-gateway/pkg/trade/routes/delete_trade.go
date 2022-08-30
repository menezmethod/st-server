package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/trade/pb"
	"strings"
)

func Delete(ctx *gin.Context, c pb.TradeServiceClient) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: id,
	})

	if res == nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

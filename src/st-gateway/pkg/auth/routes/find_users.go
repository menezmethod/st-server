package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"st-gateway/pkg/auth/pb"
)

func FindOneUser(ctx *gin.Context, c pb.AuthServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOneUser(context.Background(), &pb.FindOneUserRequest{
		Id: uint64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

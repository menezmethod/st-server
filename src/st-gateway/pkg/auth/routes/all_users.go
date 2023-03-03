package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/auth/pb"
)

func FindAllUsers(ctx *gin.Context, c pb.AuthServiceClient) {
	var id []uint64

	id = append(id, 1)

	res, err := c.FindAllUsers(context.Background(), &pb.FindAllUsersRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

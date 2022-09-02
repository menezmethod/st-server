package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/auth/pb"
	"strings"
)

func DeleteUser(ctx *gin.Context, c pb.AuthServiceClient) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Id: id,
	})

	if res == nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

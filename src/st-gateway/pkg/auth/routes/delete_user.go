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
	id := ctx.Param("id")

	if id == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("no id"))
		return
	}

	idSlice := strings.Split(id, ",")
	if len(idSlice) == 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("no id"))
		return
	}

	_, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Id: idSlice,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

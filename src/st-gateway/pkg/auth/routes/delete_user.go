package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/auth/pb"
	"strings"
)

func handleError(ctx *gin.Context, errMsg string, statusCode int) {
	ctx.JSON(statusCode, gin.H{"error": errMsg})
}

func DeleteUser(ctx *gin.Context, c pb.AuthServiceClient) {
	idParam := ctx.Param("id")
	if idParam == "" {
		handleError(ctx, "no id provided", http.StatusBadRequest)
		return
	}

	ids := strings.Split(idParam, ",")
	if len(ids) == 0 {
		handleError(ctx, "no id provided", http.StatusBadRequest)
		return
	}

	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{Id: ids})
	if err != nil {
		handleError(ctx, "error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

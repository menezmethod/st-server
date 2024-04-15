package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
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

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

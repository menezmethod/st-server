package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"st-gateway/pkg/util"
	"strings"
)

func handleError(ctx *gin.Context, errMsg string, statusCode int) {
	ctx.JSON(statusCode, gin.H{"error": errMsg})
}

func DeleteJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	idParam := ctx.Param("id")
	if idParam == "" {
		handleError(ctx, "no id provided", http.StatusBadRequest)
		return
	}

	ids := strings.Split(idParam, ",")
	res, err := c.DeleteJournal(context.Background(), &pb.DeleteJournalRequest{Id: ids})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

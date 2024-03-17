package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"st-gateway/pkg/util"
)

func FindAllJournals(ctx *gin.Context, c pb.JournalServiceClient) {
	var id []uint64

	id = append(id, 1)

	res, err := c.FindAllJournals(context.Background(), &pb.FindAllJournalsRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
)

func ListJournals(ctx *gin.Context, c pb.JournalServiceClient) {
	var id []uint64

	id = append(id, 1)

	res, err := c.ListJournals(context.Background(), &pb.FindAllJournalsRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func ListJournals(ctx *gin.Context, c pb.JournalServiceClient, user *authPb.User) {
	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.ListJournals(mdCtx, &pb.FindAllJournalsRequest{})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func FineOneJournal(ctx *gin.Context, c pb.JournalServiceClient, user *authPb.User) {
	idParam := ctx.Param("id")
	if idParam == "" {
		util.HandleError(ctx, "no id provided", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		util.HandleError(ctx, "invalid id format", http.StatusBadRequest)
		return
	}

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.GetJournal(mdCtx, &pb.FindOneJournalRequest{
		Id: id,
	})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

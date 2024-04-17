package routes

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func RemoveJournal(ctx *gin.Context, c pb.JournalServiceClient, user *authPb.User) {
	idParam := ctx.Param("id")
	if idParam == "" {
		util.HandleError(ctx, "no id provided", http.StatusBadRequest)
		return
	}

	ids := strings.Split(idParam, ",")

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.RemoveJournal(mdCtx, &pb.DeleteJournalRequest{Id: ids})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

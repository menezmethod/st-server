package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
	"strings"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

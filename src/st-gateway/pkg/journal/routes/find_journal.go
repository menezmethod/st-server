package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
	"strconv"
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
		util.HandleError(ctx, "An internal error occurred", http.StatusInternalServerError)
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

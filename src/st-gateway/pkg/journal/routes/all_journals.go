package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
)

func ListJournals(ctx *gin.Context, c pb.JournalServiceClient, user *authPb.User) {
	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.ListJournals(mdCtx, &pb.FindAllJournalsRequest{})

	if err != nil {
		util.HandleError(ctx, "An internal error occurred", http.StatusInternalServerError)
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

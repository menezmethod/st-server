package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func ListRecords(ctx *gin.Context, c pb.RecordServiceClient, user *authPb.User) {
	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.ListRecords(mdCtx, &pb.FindAllRecordsRequest{})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

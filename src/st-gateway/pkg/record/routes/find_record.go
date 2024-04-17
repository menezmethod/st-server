package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func FineOneRecord(ctx *gin.Context, c pb.RecordServiceClient, user *authPb.User) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)

	res, err := c.GetRecord(mdCtx, &pb.FindOneRecordRequest{
		Id: uint64(id),
	})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

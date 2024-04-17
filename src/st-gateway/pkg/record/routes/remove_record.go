package routes

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func RemoveRecord(ctx *gin.Context, c pb.RecordServiceClient, user *authPb.User) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.RemoveRecord(mdCtx, &pb.DeleteRecordRequest{
		Id: id,
	})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

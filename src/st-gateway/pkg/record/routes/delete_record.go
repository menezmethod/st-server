package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
	"strings"
)

func RemoveRecord(ctx *gin.Context, c pb.RecordServiceClient) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.JSON(http.StatusBadGateway, err)
		return
	}

	res, err := c.RemoveRecord(context.Background(), &pb.DeleteRecordRequest{
		Id: id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

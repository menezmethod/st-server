package routes

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"strings"
)

func DeleteJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	id, err := strings.Split(ctx.Param("id"), ","), errors.New("no id")

	if ctx.Param("id") == "" {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.DeleteJournal(context.Background(), &pb.DeleteJournalRequest{
		Id: id,
	})

	if res == nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

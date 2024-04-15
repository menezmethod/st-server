package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"net/http"
	"strconv"
)

type UpdateJournalRequestBody struct {
	Id              uint64   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CreatedBy       string   `json:"createdBy"`
	UsersSubscribed []uint64 `json:"subscribed"`
}

func UpdateJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	b := UpdateJournalRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateJournal(context.Background(), &pb.UpdateJournalRequest{
		Id:              uint64(id),
		Name:            b.Name,
		Description:     b.Description,
		StartDate:       b.StartDate,
		EndDate:         b.EndDate,
		CreatedBy:       b.CreatedBy,
		UsersSubscribed: b.UsersSubscribed,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"strconv"
)

type EditJournalRequestBody struct {
	Id              uint64   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CreatedBy       string   `json:"createdBy"`
	UsersSubscribed []uint64 `json:"subscribed"`
}

func EditJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	b := EditJournalRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.EditJournal(context.Background(), &pb.EditJournalRequest{
		Id:              uint64(id),
		Name:            b.Name,
		Description:     b.Description,
		StartDate:       b.StartDate,
		EndDate:         b.EndDate,
		CreatedBy:       b.CreatedBy,
		UsersSubscribed: b.UsersSubscribed,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

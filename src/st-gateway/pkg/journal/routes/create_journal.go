package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/journal/pb"
)

type CreateJournalRequestBody struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CreatedBy       string   `json:"createdBy"`
	UsersSubscribed []uint64 `json:"usersSubscribed"`
}

func CreateJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	b := CreateJournalRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateJournal(context.Background(), &pb.CreateJournalRequest{
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

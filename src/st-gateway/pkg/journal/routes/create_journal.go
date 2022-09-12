package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"st-gateway/pkg/journal/pb"
	"time"
)

type CreateJournalRequestBody struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	CreatedBy       string    `json:"createdBy"`
	UsersSubscribed []uint64  `json:"usersSubscribed"`
}

func CreateJournal(ctx *gin.Context, c pb.JournalServiceClient) {
	b := CreateJournalRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateJournal(context.Background(), &pb.CreateJournalRequest{
		Name:            wrapperspb.String(b.Name),
		Description:     wrapperspb.String(b.Description),
		StartDate:       timestamppb.New(b.StartDate),
		EndDate:         timestamppb.New(b.EndDate),
		CreatedBy:       wrapperspb.String(b.CreatedBy),
		UsersSubscribed: b.UsersSubscribed,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

type CreateJournalRequestBody struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StartDate       string   `json:"startDate"`
	EndDate         string   `json:"endDate"`
	CreatedBy       uint64   `json:"createdBy"`
	UsersSubscribed []uint64 `json:"usersSubscribed"`
}

func CreateJournal(ctx *gin.Context, c pb.JournalServiceClient, user *authPb.User) {
	b := CreateJournalRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	mdCtx := util.NewContextWithUserID(context.Background(), user.Id)
	res, err := c.CreateJournal(mdCtx, &pb.CreateJournalRequest{
		Name:            b.Name,
		Description:     b.Description,
		StartDate:       b.StartDate,
		EndDate:         b.EndDate,
		CreatedBy:       user.Id,
		UsersSubscribed: b.UsersSubscribed,
	})

	if err != nil {
		util.RespondWithStatus(ctx, http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

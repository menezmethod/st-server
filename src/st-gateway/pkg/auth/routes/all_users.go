package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/util"
)

func FindAllUsers(ctx *gin.Context, c pb.AuthServiceClient) {
	var id []uint64

	id = append(id, 1)

	res, err := c.FindAllUsers(context.Background(), &pb.FindAllUsersRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

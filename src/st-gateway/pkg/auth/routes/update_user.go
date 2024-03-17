package routes

import (
	"context"
	"net/http"
	"st-gateway/pkg/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth/pb"
)

type UpdateRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Bio       string `json:"bio"`
	Role      string `json:"role"`
}

func UpdateUser(ctx *gin.Context, c pb.AuthServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	b := UpdateRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		Id:        uint64(id),
		Email:     b.Email,
		Password:  b.Password,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		Bio:       b.Bio,
		Role:      b.Role,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

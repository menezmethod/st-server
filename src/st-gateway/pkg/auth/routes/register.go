package routes

import (
	"context"
	"net/http"
	"st-gateway/pkg/util"

	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:     b.Email,
		Password:  b.Password,
		FirstName: b.FirstName,
		LastName:  b.LastName,
		Role:      b.Role,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

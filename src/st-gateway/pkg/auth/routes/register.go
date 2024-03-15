package routes

import (
	"context"
	"net/http"

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
		ctx.AbortWithError(http.StatusBadRequest, err)
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
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	} else if res.Status == 409 {
		ctx.AbortWithError(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusCreated, &res)
}

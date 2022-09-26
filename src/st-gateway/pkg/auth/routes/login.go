package routes

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"

	"github.com/gin-gonic/gin"
	"st-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    wrapperspb.String(b.Email),
		Password: wrapperspb.String(b.Password),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	} else if res.Status == 404 {
		ctx.AbortWithError(http.StatusNotFound, err)
	}

	ctx.JSON(http.StatusCreated, &res)
}

package routes

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		Id:        uint64(id),
		Email:     wrapperspb.String(b.Email),
		Password:  wrapperspb.String(b.Password),
		FirstName: wrapperspb.String(b.FirstName),
		LastName:  wrapperspb.String(b.LastName),
		Bio:       wrapperspb.String(b.Bio),
		Role:      wrapperspb.String(b.Role),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	} else if res.Status == 409 {
		ctx.AbortWithError(http.StatusConflict, err)
	}

	ctx.JSON(http.StatusCreated, &res)
}

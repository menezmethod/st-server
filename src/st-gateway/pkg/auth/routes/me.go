package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"st-gateway/pkg/auth/pb"
	"st-gateway/pkg/config"
	"strings"
)

func Me(ctx *gin.Context, client pb.AuthServiceClient) {
	c, err := config.LoadConfig()

	tokenString := ctx.GetHeader("Authorization")
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.JWTSecretKey), nil
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	value, exists := claims["Id"]

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", client)
		return
	}

	userId := value.(float64)

	res, err := client.FindMe(context.Background(), &pb.FindOneUserRequest{
		Id: uint64(userId),
	})

	log.Println(token.Claims)
	ctx.JSON(http.StatusCreated, &res)
}

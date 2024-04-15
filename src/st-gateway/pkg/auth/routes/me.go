package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
	"log"
	"net/http"
	"strings"
)

func Me(ctx *gin.Context, client pb.AuthServiceClient) {
	c, err := configs.LoadConfig()

	tokenString := ctx.GetHeader("Authorization")
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.JWTSecretKey), nil
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err)
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

	if err != nil {
		ctx.JSON(http.StatusBadGateway, res)
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred"})
		return
	}

	util.RespondWithStatus(ctx, int(res.Status), res)
}

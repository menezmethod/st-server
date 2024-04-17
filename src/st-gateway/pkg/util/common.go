package util

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
)

func HandleError(ctx *gin.Context, errMsg string, statusCode int) {
	ctx.JSON(statusCode, gin.H{"error": errMsg})
}

func NewContextWithUserID(ctx context.Context, userID uint64) context.Context {
	md := metadata.New(map[string]string{
		"user-id": strconv.FormatUint(userID, 10),
	})
	return metadata.NewOutgoingContext(ctx, md)
}

func RespondWithStatus(ctx *gin.Context, status int, res interface{}) {
	switch status {
	case http.StatusOK, http.StatusCreated:
		ctx.JSON(status, res)
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusConflict:
		if err, ok := res.(error); ok {
			ctx.JSON(status, gin.H{"error": err.Error()})
		} else if errStr, ok := res.(string); ok {
			ctx.JSON(status, gin.H{"error": errStr})
		} else {
			ctx.JSON(status, res)
		}
	default:
		if err, ok := res.(error); ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else if errStr, ok := res.(string); ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errStr})
		} else {
			ctx.JSON(http.StatusInternalServerError, res)
		}
	}
	ctx.Status(status)
}

func GetUserFromContext(ctx *gin.Context) (*authPb.User, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return nil, fmt.Errorf("user ID not found in context")
	}

	userIDUint, ok := userID.(uint64)
	if !ok {
		return nil, fmt.Errorf("invalid user ID type: expected uint64, got %T", userID)
	}

	return &authPb.User{
		Id: userIDUint,
	}, nil
}

package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

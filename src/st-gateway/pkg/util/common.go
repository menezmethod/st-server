package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondWithStatus(ctx *gin.Context, status int, res interface{}) {
	switch status {
	case http.StatusOK, http.StatusCreated, http.StatusBadRequest, http.StatusUnauthorized, http.StatusNotFound, http.StatusConflict:
		ctx.JSON(status, res)
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
	}
	ctx.Status(status)
}

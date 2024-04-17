package routes

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/helper/pb"
)

type AnalyzeFinancialDataRequestBody struct {
	UserID        string `json:"userId"`
	FinancialData string `json:"financialData"`
}

func AnalyzeFinancialData(ctx *gin.Context, c pb.STHelperClient) {
	b := AnalyzeFinancialDataRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	res, err := c.AnalyzeFinancialData(timeoutCtx, &pb.FinancialRequest{
		UserId:        b.UserID,
		FinancialData: b.FinancialData,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error occurred", "details": err.Error()})
		return
	}

	ctx.Data(http.StatusOK, "application/json", []byte(res.Suggestions))
}

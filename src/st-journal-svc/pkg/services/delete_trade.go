package services

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
)

func (s *Server) DeleteTrade(ctx context.Context, req *pb.DeleteTradeRequest) (*pb.DeleteTradeResponse, error) {
	tradeWord := "trade"
	idWord := "ID"
	if len(req.Id) > 1 {
		tradeWord = "trades"
		idWord = "IDs"
	}

	s.Logger.Info(fmt.Sprintf("Received DeleteTrade request for %s: %v", idWord, req.Id))

	result, err := s.H.DB.NewDelete().Model((*models.Trade)(nil)).Where("id IN (?)", bun.In(req.Id)).Exec(ctx)
	if err != nil {
		response := createDeleteTradeResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to delete %s with %s %v: %v", tradeWord, idWord, req.Id, err), 0, err.Error(), "ERROR")
		utils.LogResponse(s.Logger, "DeleteTrade", response, http.StatusInternalServerError)
		return response, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response := createDeleteTradeResponse(http.StatusInternalServerError, fmt.Sprintf("Error checking affected rows after attempting to delete %s with %s %v: %v", tradeWord, idWord, req.Id, err), 0, err.Error(), "ERROR")
		utils.LogResponse(s.Logger, "DeleteTrade", response, http.StatusInternalServerError)
		return response, err
	}

	if rowsAffected == 0 {
		response := createDeleteTradeResponse(http.StatusNotFound, fmt.Sprintf("No %s found for the provided %s", tradeWord, idWord), rowsAffected, "", "WARNING")
		utils.LogResponse(s.Logger, "DeleteTrade", response, http.StatusNotFound)
		return response, nil
	}

	response := createDeleteTradeResponse(http.StatusOK, fmt.Sprintf("Successfully deleted %d %s with %s: %v", rowsAffected, tradeWord, idWord, req.Id), rowsAffected, "", "INFO")
	utils.LogResponse(s.Logger, "DeleteTrade", response, http.StatusOK)
	return response, nil
}

func createDeleteTradeResponse(status int, message string, rowsAffected int64, errorDetail, level string) *pb.DeleteTradeResponse {
	return &pb.DeleteTradeResponse{
		Status:       uint64(status),
		Message:      message,
		Level:        level,
		RowsAffected: uint64(rowsAffected),
		Error:        errorDetail,
	}
}

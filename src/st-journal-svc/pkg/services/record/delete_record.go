package record

import (
	"context"
	"fmt"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
	"github.com/uptrace/bun"
	"net/http"
)

func (s *Server) RemoveRecord(ctx context.Context, req *pb.DeleteRecordRequest) (*pb.DeleteRecordResponse, error) {
	tradeWord := "trade"
	idWord := "ID"
	if len(req.Id) > 1 {
		tradeWord = "trades"
		idWord = "IDs"
	}

	s.Logger.Info(fmt.Sprintf("Received RemoveRecord request for %s: %v", idWord, req.Id))

	result, err := s.H.DB.NewDelete().Model((*models.Record)(nil)).Where("id IN (?)", bun.In(req.Id)).Exec(ctx)
	if err != nil {
		response := createDeleteTradeResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to delete %s with %s %v: %v", tradeWord, idWord, req.Id, err), 0, err.Error(), "ERROR")
		utils.LogResponse(s.Logger, "RemoveRecord", response, http.StatusInternalServerError)
		return response, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response := createDeleteTradeResponse(http.StatusInternalServerError, fmt.Sprintf("Error checking affected rows after attempting to delete %s with %s %v: %v", tradeWord, idWord, req.Id, err), 0, err.Error(), "ERROR")
		utils.LogResponse(s.Logger, "RemoveRecord", response, http.StatusInternalServerError)
		return response, err
	}

	if rowsAffected == 0 {
		response := createDeleteTradeResponse(http.StatusNotFound, fmt.Sprintf("No %s found for the provided %s", tradeWord, idWord), rowsAffected, "", "WARNING")
		utils.LogResponse(s.Logger, "RemoveRecord", response, http.StatusNotFound)
		return response, nil
	}

	response := createDeleteTradeResponse(http.StatusOK, fmt.Sprintf("Successfully deleted %d %s with %s: %v", rowsAffected, tradeWord, idWord, req.Id), rowsAffected, "", "INFO")
	utils.LogResponse(s.Logger, "RemoveRecord", response, http.StatusOK)
	return response, nil
}

func createDeleteTradeResponse(status int, message string, rowsAffected int64, errorDetail, level string) *pb.DeleteRecordResponse {
	return &pb.DeleteRecordResponse{
		Status:       uint64(status),
		Message:      message,
		Level:        level,
		RowsAffected: uint64(rowsAffected),
		Error:        errorDetail,
	}
}

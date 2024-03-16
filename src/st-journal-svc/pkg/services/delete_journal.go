package services

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
)

func (s *Server) DeleteJournal(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	s.Logger.Debug("Received DeleteJournal request", zap.Strings("IDs", req.Id))

	if len(req.Id) == 0 {
		response := createDeleteJournalResponse(http.StatusBadRequest, "No ID provided for deletion", "Request must include at least one ID to delete a journal", 0)
		utils.LogResponse(s.Logger, "DeleteJournal", response, http.StatusBadRequest)
		return response, nil
	}

	result, err := s.H.DB.NewDelete().Model((*models.Journal)(nil)).Where("id IN (?)", bun.In(req.Id)).Exec(ctx)
	var status int
	var message string
	var errorDetail string
	var rowsAffected int64

	if err != nil {
		status = http.StatusInternalServerError
		message = "Failed to delete journal"
		errorDetail = err.Error()
	} else {
		rowsAffected, err = result.RowsAffected()
		if err != nil {
			status = http.StatusInternalServerError
			message = "Error checking affected rows after deletion"
			errorDetail = err.Error()
		} else if rowsAffected == 0 {
			status = http.StatusNotFound
			message = "No journal found for the provided ID"
			errorDetail = "Journal not found"
		} else {
			status = http.StatusOK
			message = fmt.Sprintf("Successfully deleted journal with ID(s): %v", req.Id)
		}
	}

	response := createDeleteJournalResponse(status, message, errorDetail, rowsAffected)
	utils.LogResponse(s.Logger, "DeleteJournal", response, status)
	return response, nil
}

func createDeleteJournalResponse(status int, message, errorDetail string, rowsAffected int64) *pb.DeleteJournalResponse {
	level := utils.GetStatusLevel(status)
	return &pb.DeleteJournalResponse{
		Status:       uint64(status),
		Message:      message,
		Level:        level,
		Error:        errorDetail,
		RowsAffected: uint64(rowsAffected),
	}
}

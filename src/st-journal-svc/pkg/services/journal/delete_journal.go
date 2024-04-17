package journal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/uptrace/bun"
	"go.uber.org/zap"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

func (s *Server) RemoveJournal(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	s.Logger.Debug("Received RemoveJournal request", zap.Strings("IDs", req.Id))

	if len(req.Id) == 0 {
		response := createDeleteJournalResponse(http.StatusBadRequest, "No ID provided for deletion", "Request must include at least one ID to delete a journal", 0)
		utils.LogResponse(s.Logger, "RemoveJournal", response, http.StatusBadRequest)
		return response, nil
	}

	var journal models.Journal
	if err := s.H.DB.NewSelect().Model(&journal).Where("ID IN (?)", req.Id).Scan(ctx); err != nil {
		s.Logger.Error("Failed to get journal", zap.Strings("IDs", req.Id), zap.Error(err))
		response := createDeleteJournalResponse(http.StatusInternalServerError, "Failed to get journal", "An internal error occurred", 0)
		utils.LogResponse(s.Logger, "RemoveJournal", response, http.StatusInternalServerError)
		return response, nil
	}

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: journal.CreatedBy})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.String("UserID", strconv.FormatUint(journal.CreatedBy, 10)), zap.Error(err))
		response := createDeleteJournalResponse(http.StatusInternalServerError, "Failed to get user from auth service", "An internal error occurred", 0)
		utils.LogResponse(s.Logger, "RemoveJournal", response, http.StatusInternalServerError)
		return response, nil
	}

	if authRes.Data.Role != "ADMIN" {
		count, err := s.H.DB.NewSelect().Model(&models.Journal{}).Where("ID IN (?)", req.Id).Where("CreatedBy = ?", authRes.Data.Id).Count(ctx)
		if err != nil {
			s.Logger.Error("Failed to check journal ownership", zap.Strings("IDs", req.Id), zap.Error(err))
			response := createDeleteJournalResponse(http.StatusInternalServerError, "Failed to check journal ownership", "An internal error occurred", 0)
			utils.LogResponse(s.Logger, "RemoveJournal", response, http.StatusInternalServerError)
			return response, nil
		}
		if count != len(req.Id) {
			errorMsg := "Unauthorized to delete one or more journals"
			s.Logger.Error(errorMsg, zap.Strings("IDs", req.Id), zap.String("UserID", strconv.FormatUint(authRes.Data.Id, 10)))
			response := createDeleteJournalResponse(http.StatusForbidden, "Unauthorized to delete one or more journals", errorMsg, 0)
			utils.LogResponse(s.Logger, "RemoveJournal", response, http.StatusForbidden)
			return response, nil
		}
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
			message = "Journal not found for the provided ID"
			errorDetail = "Journal not found"
		} else {
			status = http.StatusOK
			message = fmt.Sprintf("Successfully deleted journal with ID(s): %v", req.Id)
		}
	}

	response := createDeleteJournalResponse(status, message, errorDetail, rowsAffected)
	utils.LogResponse(s.Logger, "RemoveJournal", response, status)
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

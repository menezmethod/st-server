package services

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
	"strconv"
	"time"
)

func populateJournalFromEditRequest(req *pb.UpdateJournalRequest) models.Journal {
	return models.Journal{
		ID:              req.GetId(),
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		StartDate:       req.GetStartDate(),
		EndDate:         req.GetEndDate(),
		CreatedBy:       req.GetCreatedBy(),
		UsersSubscribed: req.GetUsersSubscribed(),
	}
}

func createUpdateJournalResponse(logger *zap.Logger, journal models.Journal, status uint64, errorMessage string) *pb.UpdateJournalResponse {
	timestamp := time.Now().Format(time.RFC3339)
	level := "INFO"
	message := "Journal updated successfully"

	if errorMessage != "" {
		level = "ERROR"
		message = "Failed to update journal"
		logger.Error("Failed to update journal", zap.String("error", errorMessage), zap.Any("journal", journal))
	}

	response := &pb.UpdateJournalResponse{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Status:    status,
		Data: &pb.UpdateJournalData{
			Id:              journal.ID,
			Name:            journal.Name,
			Description:     journal.Description,
			StartDate:       journal.StartDate,
			EndDate:         journal.EndDate,
			CreatedBy:       journal.CreatedBy,
			UsersSubscribed: journal.UsersSubscribed,
		},
		Error: errorMessage,
	}

	logger.Info("Journal updated successfully", zap.Any("journal", journal))
	return response
}

func (s *Server) UpdateJournal(ctx context.Context, req *pb.UpdateJournalRequest) (*pb.UpdateJournalResponse, error) {
	s.Logger.Debug("Received UpdateJournal request", zap.String("ID", strconv.FormatUint(req.GetId(), 10)))

	journal := populateJournalFromEditRequest(req)
	errorMsg := utils.ValidateJournal(&journal)

	var resp *pb.UpdateJournalResponse
	if errorMsg != "" {
		s.Logger.Error("Validation failed for UpdateJournal", zap.String("ID", fmt.Sprintf("%v", journal.ID)), zap.String("error", errorMsg))
		resp = createUpdateJournalResponse(s.Logger, journal, http.StatusBadRequest, errorMsg)
	} else if _, err := s.H.DB.NewUpdate().Model(&journal).Where("ID = ?", journal.ID).Exec(ctx); err != nil {
		s.Logger.Error("Failed to update journal", zap.String("ID", fmt.Sprintf("%v", journal.ID)), zap.Error(err))
		resp = createUpdateJournalResponse(s.Logger, journal, http.StatusConflict, err.Error())
	} else {
		resp = createUpdateJournalResponse(s.Logger, journal, http.StatusOK, "")
	}

	utils.LogResponse(s.Logger, "UpdateJournal", resp, int(resp.Status))

	return resp, nil
}

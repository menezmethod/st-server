package journal

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc/metadata"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

func (s *Server) UpdateJournal(ctx context.Context, req *pb.UpdateJournalRequest) (*pb.UpdateJournalResponse, error) {
	s.Logger.Debug("Received UpdateJournal request", zap.String("JournalID", strconv.FormatUint(req.GetId(), 10)))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Warn("No metadata received with request")
		return createUpdateJournalResponse(&models.Journal{}, http.StatusBadRequest, "no metadata received with request"), nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		s.Logger.Warn("User ID not provided in metadata")
		return createUpdateJournalResponse(&models.Journal{}, http.StatusUnauthorized, "user-id not provided in metadata"), nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		s.Logger.Error("Invalid user ID format", zap.Error(err))
		return createUpdateJournalResponse(&models.Journal{}, http.StatusUnauthorized, err.Error()), nil
	}

	s.Logger.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.Error(err))
		return createUpdateJournalResponse(&models.Journal{}, http.StatusInternalServerError, err.Error()), nil
	}

	if authRes == nil || authRes.Data == nil {
		s.Logger.Error("Invalid response from auth service")
		return createUpdateJournalResponse(&models.Journal{}, http.StatusInternalServerError, "invalid response from auth service"), nil
	}

	s.Logger.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var existingJournal models.Journal
	if err := s.H.DB.NewSelect().Model(&existingJournal).Where("ID = ?", req.GetId()).Scan(ctx); err != nil {
		s.Logger.Error("Failed to retrieve journal from database", zap.Error(err))
		return createUpdateJournalResponse(&models.Journal{}, http.StatusInternalServerError, "failed to retrieve journal"), nil
	}

	if existingJournal.CreatedBy != loggedInUserID && authRes.Data.Role != "ADMIN" {
		s.Logger.Error("Unauthorized attempt to update journal", zap.Uint64("JournalID", req.GetId()), zap.Uint64("AttemptedByUserID", loggedInUserID))
		return createUpdateJournalResponse(&models.Journal{}, http.StatusForbidden, "unauthorized to update this journal"), nil
	}

	journal := populateJournalFromEditRequest(req)
	if _, err := s.H.DB.NewUpdate().Model(&journal).Where("ID = ?", journal.ID).Exec(ctx); err != nil {
		s.Logger.Error("Failed to update journal in database", zap.Error(err))
		return createUpdateJournalResponse(&models.Journal{}, http.StatusInternalServerError, "failed to update journal"), nil
	}

	s.Logger.Info("Journal updated successfully", zap.Uint64("JournalID", journal.ID))
	return createUpdateJournalResponse(&journal, http.StatusOK, ""), nil
}

func populateJournalFromEditRequest(req *pb.UpdateJournalRequest) models.Journal {
	return models.Journal{
		ID:              req.GetId(),
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		StartDate:       req.GetStartDate(),
		EndDate:         req.GetEndDate(),
		LastUpdatedBy:   req.GetLastUpdatedBy(),
		UsersSubscribed: req.GetUsersSubscribed(),
	}
}

func createUpdateJournalResponse(journal *models.Journal, status uint64, errorMessage string) *pb.UpdateJournalResponse {
	timestamp := time.Now().Format(time.RFC3339)
	response := &pb.UpdateJournalResponse{
		Timestamp: timestamp,
		Level:     utils.GetStatusLevel(int(status)),
		Status:    status,
	}

	if errorMessage != "" {
		response.Error = errorMessage
	} else {
		response.Message = "Journal updated successfully"
		if journal.ID != 0 {
			response.Data = &pb.Journal{
				Id:              journal.ID,
				Name:            journal.Name,
				Description:     journal.Description,
				StartDate:       journal.StartDate,
				EndDate:         journal.EndDate,
				LastUpdatedBy:   journal.LastUpdatedBy,
				UsersSubscribed: journal.UsersSubscribed,
			}
		}
	}

	return response
}

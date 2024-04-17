package journal

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

func (s *Server) RemoveJournal(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	s.Logger.Debug("Received RemoveJournal request", zap.Strings("IDs", req.Id))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Warn("No metadata received with request")
		return createDeleteJournalResponse(http.StatusBadRequest, "No metadata received with request", "Request must include metadata", 0), nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		s.Logger.Warn("User ID not provided in metadata")
		return createDeleteJournalResponse(http.StatusUnauthorized, "User ID not provided in metadata", "user-id not provided in metadata", 0), nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		s.Logger.Error("Invalid user ID format", zap.Error(err))
		return createDeleteJournalResponse(http.StatusUnauthorized, "Invalid user ID format", err.Error(), 0), nil
	}

	s.Logger.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.Error(err))
		return createDeleteJournalResponse(http.StatusInternalServerError, "Failed to get user from auth service", err.Error(), 0), nil
	}

	if authRes == nil || authRes.Data == nil {
		s.Logger.Error("Invalid response from auth service")
		return createDeleteJournalResponse(http.StatusInternalServerError, "Invalid response from auth service", "invalid response from auth service", 0), nil
	}

	s.Logger.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var journalsToDelete []models.Journal
	if err := s.H.DB.NewSelect().Model(&journalsToDelete).Where("ID IN (?)", bun.In(req.Id)).Scan(ctx); err != nil {
		s.Logger.Error("Failed to retrieve journals from database", zap.Error(err))
		return createDeleteJournalResponse(http.StatusInternalServerError, "Failed to retrieve journals", "failed to retrieve journals", 0), nil
	}

	if authRes.Data.Role != "ADMIN" {
		for _, journal := range journalsToDelete {
			if journal.CreatedBy != loggedInUserID {
				s.Logger.Error("Unauthorized attempt to delete journal", zap.Uint64("JournalID", journal.ID), zap.Uint64("AttemptedByUserID", loggedInUserID))
				return createDeleteJournalResponse(http.StatusForbidden, "Unauthorized to delete one or more journals", "unauthorized to delete one or more journals", 0), nil
			}
		}
	}

	result, err := s.H.DB.NewDelete().Model((*models.Journal)(nil)).Where("id IN (?)", bun.In(req.Id)).Exec(ctx)
	var status int
	var message string
	var errorDetail string
	var rowsAffected int64

	if err != nil {
		status = http.StatusInternalServerError
		message = "Failed to delete journals"
		errorDetail = err.Error()
	} else {
		rowsAffected, err = result.RowsAffected()
		if err != nil {
			status = http.StatusInternalServerError
			message = "Error checking affected rows after deletion"
			errorDetail = err.Error()
		} else if rowsAffected == 0 {
			status = http.StatusNotFound
			message = "Journals not found for the provided IDs"
			errorDetail = "Journals not found"
		} else {
			status = http.StatusOK
			message = fmt.Sprintf("Successfully deleted %d journal(s) with ID(s): %v", rowsAffected, req.Id)
		}
	}

	response := createDeleteJournalResponse(uint64(status), message, errorDetail, uint64(rowsAffected))
	utils.LogResponse(s.Logger, "RemoveJournal", response, status)
	return response, nil
}

func createDeleteJournalResponse(status uint64, message, errorDetail string, rowsAffected uint64) *pb.DeleteJournalResponse {
	response := &pb.DeleteJournalResponse{
		Level:        utils.GetStatusLevel(int(status)),
		Status:       status,
		RowsAffected: rowsAffected,
	}

	if errorDetail != "" {
		response.Error = errorDetail
	} else {
		response.Message = message
	}

	return response
}

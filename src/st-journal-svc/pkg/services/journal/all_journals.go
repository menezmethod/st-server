package journal

import (
	"context"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
)

func (s *Server) ListJournals(ctx context.Context, _ *pb.FindAllJournalsRequest) (*pb.FindAllJournalsResponse, error) {
	s.Logger.Debug("Received ListJournals request")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Warn("No metadata received with request")
		return &pb.FindAllJournalsResponse{
			Status: http.StatusBadRequest,
			Error:  "No metadata received with request",
		}, nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		s.Logger.Warn("User ID not provided in metadata")
		return &pb.FindAllJournalsResponse{
			Status: http.StatusUnauthorized,
			Error:  "User ID not provided in metadata",
		}, nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		s.Logger.Error("Invalid user ID format", zap.Error(err))
		return &pb.FindAllJournalsResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid user ID format",
		}, nil
	}

	s.Logger.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.Error(err))
		return &pb.FindAllJournalsResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to get user from auth service",
		}, nil
	}

	if authRes == nil || authRes.Data == nil {
		s.Logger.Error("Invalid response from auth service")
		return &pb.FindAllJournalsResponse{
			Status: http.StatusInternalServerError,
			Error:  "Invalid response from auth service",
		}, nil
	}

	s.Logger.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var modelJournals []models.Journal
	if authRes.Data.Role == "ADMIN" {
		if err := s.H.DB.NewSelect().Model(&modelJournals).Scan(ctx); err != nil {
			s.Logger.Error("Error retrieving all journals", zap.Error(err))
			return &pb.FindAllJournalsResponse{
				Status: http.StatusNotFound,
				Error:  err.Error(),
			}, nil
		}
	} else {
		if err := s.H.DB.NewSelect().Model(&modelJournals).Where("created_by = ?", loggedInUserID).Scan(ctx); err != nil {
			s.Logger.Error("Error retrieving user's journals", zap.Error(err))
			return &pb.FindAllJournalsResponse{
				Status: http.StatusNotFound,
				Error:  err.Error(),
			}, nil
		}
	}

	journals := make([]*pb.Journal, len(modelJournals))
	for i := 0; i < len(modelJournals); i++ {
		journals[i] = mapJournalToPBJournal(&modelJournals[i])
	}

	s.Logger.Info("Successfully retrieved journals", zap.Int("count", len(journals)))
	return &pb.FindAllJournalsResponse{Data: journals, Status: http.StatusOK}, nil
}

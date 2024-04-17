package journal

import (
	"context"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
)

func mapJournalToPBJournal(journal models.Journal) *pb.Journal {
	return &pb.Journal{
		Id:              journal.ID,
		CreatedAt:       timestamppb.New(journal.CreatedAt),
		CreatedBy:       journal.CreatedBy,
		Description:     journal.Description,
		EndDate:         journal.EndDate,
		LastUpdatedBy:   journal.LastUpdatedBy,
		Name:            journal.Name,
		StartDate:       journal.StartDate,
		UsersSubscribed: journal.UsersSubscribed,
	}
}

func (s *Server) GetJournal(ctx context.Context, req *pb.FindOneJournalRequest) (*pb.FindOneJournalResponse, error) {
	log := s.Logger.With(zap.Int64("requestId", int64(req.Id)))
	log.Debug("Received GetJournal request")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Warn("No metadata received with request")
		return &pb.FindOneJournalResponse{
			Status: http.StatusBadRequest,
			Error:  "No metadata received with request",
		}, nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		log.Warn("User ID not provided in metadata")
		return &pb.FindOneJournalResponse{
			Status: http.StatusUnauthorized,
			Error:  "User ID not provided in metadata",
		}, nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		log.Error("Invalid user ID format", zap.Error(err))
		return &pb.FindOneJournalResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid user ID format",
		}, nil
	}

	log.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		log.Error("Failed to get user from auth service", zap.Error(err))
		return &pb.FindOneJournalResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to get user from auth service",
		}, nil
	}

	if authRes == nil || authRes.Data == nil {
		log.Error("Invalid response from auth service")
		return &pb.FindOneJournalResponse{
			Status: http.StatusInternalServerError,
			Error:  "Invalid response from auth service",
		}, nil
	}

	log.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var journal models.Journal
	if err := s.H.DB.NewSelect().Model(&journal).Where("id = ?", req.Id).Scan(ctx); err != nil {
		log.Error("Error retrieving journal", zap.Error(err))
		return &pb.FindOneJournalResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	if journal.CreatedBy != loggedInUserID && authRes.Data.Role != "ADMIN" {
		log.Error("Unauthorized attempt to access journal", zap.Uint64("JournalID", journal.ID), zap.Uint64("AttemptedByUserID", loggedInUserID))
		return &pb.FindOneJournalResponse{
			Status: http.StatusForbidden,
			Error:  "Unauthorized to access this journal",
		}, nil
	}

	log.Info("Successfully retrieved journal")
	return &pb.FindOneJournalResponse{
		Status: http.StatusOK,
		Data:   mapJournalToPBJournal(journal),
	}, nil
}

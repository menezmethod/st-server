package record

import (
	"context"
	"net/http"
	"strconv"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
)

func (s *Server) ListRecords(ctx context.Context, _ *pb.FindAllRecordsRequest) (*pb.FindAllRecordsResponse, error) {
	s.Logger.Debug("Received ListRecords request")

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Warn("No metadata received with request")
		return &pb.FindAllRecordsResponse{
			Status: http.StatusBadRequest,
			Error:  "No metadata received with request",
		}, nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		s.Logger.Warn("User ID not provided in metadata")
		return &pb.FindAllRecordsResponse{
			Status: http.StatusUnauthorized,
			Error:  "User ID not provided in metadata",
		}, nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		s.Logger.Error("Invalid user ID format", zap.Error(err))
		return &pb.FindAllRecordsResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid user ID format",
		}, nil
	}

	s.Logger.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.Error(err))
		return &pb.FindAllRecordsResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to get user from auth service",
		}, nil
	}

	if authRes == nil || authRes.Data == nil {
		s.Logger.Error("Invalid response from auth service")
		return &pb.FindAllRecordsResponse{
			Status: http.StatusInternalServerError,
			Error:  "Invalid response from auth service",
		}, nil
	}

	s.Logger.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var modelRecords []models.Record
	if authRes.Data.Role == "ADMIN" {
		if err := s.H.DB.NewSelect().Model(&modelRecords).Scan(ctx); err != nil {
			s.Logger.Error("Error retrieving all records", zap.Error(err))
			return &pb.FindAllRecordsResponse{
				Status: http.StatusNotFound,
				Error:  err.Error(),
			}, nil
		}
	} else {
		if err := s.H.DB.NewSelect().Model(&modelRecords).Where("created_by = ?", loggedInUserID).Scan(ctx); err != nil {
			s.Logger.Error("Error retrieving user's records", zap.Error(err))
			return &pb.FindAllRecordsResponse{
				Status: http.StatusNotFound,
				Error:  err.Error(),
			}, nil
		}
	}

	records := make([]*pb.Record, len(modelRecords))
	for i, record := range modelRecords {
		records[i] = mapModelRecordToPBRecord(record)
	}

	s.Logger.Info("Successfully retrieved records", zap.Int("count", len(records)))
	return &pb.FindAllRecordsResponse{Data: records, Status: http.StatusOK}, nil
}

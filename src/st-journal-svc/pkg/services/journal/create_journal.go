package journal

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/db"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

type Server struct {
	AuthServiceClient pb.AuthServiceClient
	H                 db.DB
	Logger            *zap.Logger
	Validator         *validator.Validate
	pb.JournalServiceServer
}

func (s *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	if s.Logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	log := s.Logger.With(
		zap.String("action", "CreateJournal"),
		zap.String("name", req.GetName()),
		zap.Time("requestTime", time.Now()),
	)

	log.Debug("Received CreateJournal request")

	journal := populateJournalFromRequest(req)
	errorMsg := utils.ValidateJournal(&journal)

	var resp *pb.CreateJournalResponse
	if errorMsg != "" {
		log.Error("Validation failed for CreateJournal", zap.String("error", errorMsg))
		resp = createJournalResponse(journal, http.StatusBadRequest, errorMsg)
	} else {
		if _, err := s.H.DB.NewInsert().Model(&journal).Exec(ctx); err != nil {
			log.Error("Failed to insert journal", zap.Error(err))
			resp = createJournalResponse(journal, http.StatusInternalServerError, "Failed to insert journal")
		} else {
			log.Info("Journal created successfully", zap.Any("journal", journal))
			resp = createJournalResponse(journal, http.StatusCreated, "")
		}
	}

	utils.LogResponse(s.Logger, "CreateJournal", resp, int(resp.Status))

	return resp, nil
}
func populateJournalFromRequest(req *pb.CreateJournalRequest) models.Journal {
	return models.Journal{
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		CreatedAt:       time.Now(),
		StartDate:       req.GetStartDate(),
		EndDate:         req.GetEndDate(),
		CreatedBy:       req.GetCreatedBy(),
		UsersSubscribed: req.GetUsersSubscribed(),
	}
}

func createJournalResponse(journal models.Journal, status int, message string) *pb.CreateJournalResponse {
	return &pb.CreateJournalResponse{
		Timestamp: time.Now().Format(time.RFC1123),
		Level:     utils.GetStatusLevel(status),
		Message:   message,
		Status:    uint64(status),
		Journal: &pb.Journal{
			CreatedAt:       timestamppb.New(journal.CreatedAt),
			CreatedBy:       journal.CreatedBy,
			Description:     journal.Description,
			EndDate:         journal.EndDate,
			Id:              journal.ID,
			Name:            journal.Name,
			StartDate:       journal.StartDate,
			UsersSubscribed: journal.UsersSubscribed,
		},
	}
}

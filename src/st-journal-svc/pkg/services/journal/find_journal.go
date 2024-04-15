package journal

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
)

func mapJournalToPBJournal(journal models.Journal) *pb.Journal {
	return &pb.Journal{
		Id:              journal.ID,
		Name:            journal.Name,
		Description:     journal.Description,
		StartDate:       journal.StartDate,
		EndDate:         journal.EndDate,
		CreatedBy:       journal.CreatedBy,
		CreatedAt:       timestamppb.New(journal.CreatedAt),
		UsersSubscribed: journal.UsersSubscribed,
	}
}

func (s *Server) GetJournal(ctx context.Context, req *pb.FindOneJournalRequest) (*pb.FindOneJournalResponse, error) {
	log := s.Logger.With(zap.Int64("requestId", int64(req.Id)))
	log.Debug("Received GetJournal request")

	var journal models.Journal
	if err := s.H.DB.NewSelect().Model(&journal).Where("id = ?", req.Id).Scan(ctx); err != nil {
		log.Error("Error retrieving journal", zap.Error(err))
		return &pb.FindOneJournalResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	log.Info("Successfully retrieved journal")
	return &pb.FindOneJournalResponse{
		Status: http.StatusOK,
		Data:   mapJournalToPBJournal(journal),
	}, nil
}

func (s *Server) ListJournals(ctx context.Context, _ *pb.FindAllJournalsRequest) (*pb.FindAllJournalsResponse, error) {
	s.Logger.Debug("Received ListJournals request")

	var modelJournals []models.Journal
	if err := s.H.DB.NewSelect().Model(&modelJournals).Scan(ctx); err != nil {
		s.Logger.Error("Error retrieving all journals", zap.Error(err))
		return &pb.FindAllJournalsResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	journals := make([]*pb.Journal, len(modelJournals))
	for i, journal := range modelJournals {
		journals[i] = mapJournalToPBJournal(journal)
	}

	s.Logger.Info("Successfully retrieved journals", zap.Int("count", len(journals)))
	return &pb.FindAllJournalsResponse{Data: journals, Status: http.StatusOK}, nil
}

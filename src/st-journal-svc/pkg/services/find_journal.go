package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

type Server struct {
	H db.DB
	pb.JournalServiceServer
}

func mapJournalToPBJournal(journal models.Journal) *pb.Journal {
	return &pb.Journal{
		Id:              journal.ID,
		Name:            journal.Name,
		Description:     journal.Description,
		StartDate:       journal.StartDate,
		EndDate:         journal.EndDate,
		CreatedBy:       journal.CreatedBy,
		CreatedAt:       journal.CreatedAt.String(),
		UsersSubscribed: journal.UsersSubscribed,
	}
}

func (s *Server) FindAllJournals(ctx context.Context, _ *pb.FindAllJournalsRequest) (*pb.FindAllJournalsResponse, error) {
	var modelJournals []models.Journal
	if err := s.H.DB.NewSelect().Model(&modelJournals).Scan(ctx); err != nil {
		return &pb.FindAllJournalsResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	journals := make([]*pb.Journal, len(modelJournals))
	for i, journal := range modelJournals {
		journals[i] = mapJournalToPBJournal(journal)
	}

	return &pb.FindAllJournalsResponse{Data: journals}, nil
}

func (s *Server) FindOneJournal(ctx context.Context, req *pb.FindOneJournalRequest) (*pb.FindOneJournalResponse, error) {
	var journal models.Journal
	if err := s.H.DB.NewSelect().Model(&journal).Where("id = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneJournalResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.FindOneJournalResponse{
		Status: http.StatusOK,
		Data:   mapJournalToPBJournal(journal),
	}, nil
}

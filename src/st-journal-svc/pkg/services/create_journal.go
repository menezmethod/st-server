package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"time"
)

func validateJournal(journal *models.Journal) string {
	if journal.Name == "" {
		return "Name cannot be empty"
	}
	if journal.Description == "" {
		return "Description cannot be empty"
	}
	return ""
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

func createJournalResponse(journal models.Journal) *pb.CreateJournalResponse {
	return &pb.CreateJournalResponse{
		Status: http.StatusCreated,
		Data: &pb.Journal{
			Id:              journal.ID,
			Name:            journal.Name,
			Description:     journal.Description,
			StartDate:       journal.StartDate,
			EndDate:         journal.EndDate,
			CreatedAt:       journal.CreatedAt.String(),
			CreatedBy:       journal.CreatedBy,
			UsersSubscribed: journal.UsersSubscribed,
		},
	}
}

func (s *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	journal := populateJournalFromRequest(req)
	errorMsg := validateJournal(&journal)
	if errorMsg != "" {
		return &pb.CreateJournalResponse{
			Status: http.StatusBadRequest,
			Error:  errorMsg,
		}, nil
	}

	if _, err := s.H.DB.NewInsert().Model(&journal).Exec(ctx); err != nil {
		return &pb.CreateJournalResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return createJournalResponse(journal), nil
}

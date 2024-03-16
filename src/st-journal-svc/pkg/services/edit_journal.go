package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

func populateJournalFromEditRequest(req *pb.EditJournalRequest) models.Journal {
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

func createEditJournalResponse(journal models.Journal) *pb.EditJournalResponse {
	return &pb.EditJournalResponse{
		Status: http.StatusOK,
		Data: &pb.EditJournalData{
			Id:              journal.ID,
			Name:            journal.Name,
			Description:     journal.Description,
			StartDate:       journal.StartDate,
			EndDate:         journal.EndDate,
			CreatedBy:       journal.CreatedBy,
			UsersSubscribed: journal.UsersSubscribed,
		},
	}
}

func (s *Server) EditJournal(ctx context.Context, req *pb.EditJournalRequest) (*pb.EditJournalResponse, error) {
	journal := populateJournalFromEditRequest(req)
	errorMsg := validateJournal(&journal)
	if errorMsg != "" {
		return &pb.EditJournalResponse{
			Status: http.StatusBadRequest,
			Error:  errorMsg,
		}, nil
	}

	if _, err := s.H.DB.NewUpdate().Model(&journal).Where("ID = ?", journal.ID).Exec(ctx); err != nil {
		return &pb.EditJournalResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return createEditJournalResponse(journal), nil
}

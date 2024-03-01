package services

import (
	"context"
	"github.com/uptrace/bun"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

func (s *Server) DeleteJournals(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	ids := req.Id

	result, err := s.H.DB.NewDelete().Model((*models.Journal)(nil)).Where("id IN (?)", bun.In(ids)).Exec(ctx)
	if err != nil {
		return &pb.DeleteJournalResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &pb.DeleteJournalResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error checking affected rows: " + err.Error(),
		}, nil
	}

	if rowsAffected == 0 {
		return &pb.DeleteJournalResponse{
			Status: http.StatusNotFound,
			Error:  "Journal(s) not found",
		}, nil
	}

	return &pb.DeleteJournalResponse{
		Status: http.StatusOK,
	}, nil
}

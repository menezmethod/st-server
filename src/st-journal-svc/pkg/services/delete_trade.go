package services

import (
	"context"
	"github.com/uptrace/bun"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

func (s *Server) DeleteTrade(ctx context.Context, req *pb.DeleteTradeRequest) (*pb.DeleteTradeResponse, error) {
	ids := req.Id

	result, err := s.H.DB.NewDelete().Model((*models.Trade)(nil)).Where("id IN (?)", bun.In(ids)).Exec(ctx)
	if err != nil {
		return &pb.DeleteTradeResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &pb.DeleteTradeResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to determine affected rows: " + err.Error(),
		}, nil
	}

	if rowsAffected == 0 {
		return &pb.DeleteTradeResponse{
			Status: http.StatusNotFound,
			Error:  "Trades(s) not found",
		}, nil
	}

	return &pb.DeleteTradeResponse{
		Status: http.StatusOK,
	}, nil
}

package services

import (
	"context"
	"net/http"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	result, err := s.H.DB.NewDelete().Model(&models.User{}).Where("ID IN (?)", req.Id).Exec(ctx)
	if err != nil {
		return &pb.DeleteUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to delete user: " + err.Error(),
		}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &pb.DeleteUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to determine rows affected: " + err.Error(),
		}, err
	}

	if rowsAffected == 0 {
		return &pb.DeleteUserResponse{
			Status: http.StatusNotFound,
			Error:  "No users found with the provided ID(s)",
		}, nil
	}

	return &pb.DeleteUserResponse{
		Status: http.StatusOK,
	}, nil
}

package services

import (
	"context"
	"fmt"
	"net/http"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
	"time"
)

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	s.Logger.Debug("Received UpdateUser request")

	if req == nil {
		return s.generateResponse(utils.GetStatusLevel(http.StatusBadRequest), "Failed to update user", "Request is nil", http.StatusBadRequest, nil)
	}

	user, err := s.getUserByID(ctx, req.Id)
	if err != nil {
		return s.generateResponse(utils.GetStatusLevel(http.StatusNotFound), "Failed to find user", fmt.Sprintf("User with ID %d not found: %v", req.Id, err), http.StatusNotFound, nil)
	}

	s.updateFieldsFromRequest(&user, req)

	result, err := s.H.DB.NewUpdate().Model(&user).ExcludeColumn("created_at").Where("ID = ?", user.Id).Exec(ctx)
	if err != nil || result == nil {
		return s.generateResponse(utils.GetStatusLevel(http.StatusInternalServerError), "Failed to update user", fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError, nil)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return s.generateResponse(utils.GetStatusLevel(http.StatusNoContent), "No user updated", fmt.Sprintf("No changes made for user with ID %d", req.Id), http.StatusNoContent, nil)
	}

	return s.generateResponse(utils.GetStatusLevel(http.StatusOK), "User updated successfully", "", http.StatusOK, &user)
}

func (s *Server) updateFieldsFromRequest(user *models.User, req *pb.UpdateUserRequest) {
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(s.Logger, req.Password, 10)
		if err != nil {
			return
		}
		user.Password = hashedPassword
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Role != "" {
		user.Role = req.Role
	}
}

func (s *Server) generateResponse(level, message, errMsg string, status int64, user *models.User) (*pb.UpdateUserResponse, error) {
	resp := &pb.UpdateUserResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Status:    uint64(status),
	}

	if errMsg != "" {
		resp.Error = errMsg
	}

	if user != nil {
		resp.Data = &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
		}
	}

	utils.LogResponse(s.Logger, "UpdateUser", resp, int(resp.Status))
	return resp, nil
}

func (s *Server) getUserByID(ctx context.Context, id uint64) (models.User, error) {
	var user models.User
	err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", id).Scan(ctx)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

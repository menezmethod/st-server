package services

import (
	"context"
	"net/http"

	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/pb"
	"go.uber.org/zap"
)

func (s *Server) FindAllUsers(ctx context.Context, _ *pb.FindAllUsersRequest) (*pb.FindAllUsersResponse, error) {
	s.Logger.Info("FindAllUsers request received")

	var userModel []models.User
	users := make([]*pb.User, 0)
	if err := s.H.DB.NewSelect().Model(&userModel).Column("id", "email", "first_name", "last_name", "role", "created_at").Scan(ctx); err != nil {
		s.Logger.Error("Error retrieving users", zap.Error(err))
		return &pb.FindAllUsersResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	for i := range userModel {
		user := &userModel[i]
		users = append(users, &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
		})
	}

	s.Logger.Info("Successfully retrieved users", zap.Int("count", len(users)))

	return &pb.FindAllUsersResponse{
		Status: http.StatusOK,
		Data:   users,
	}, nil
}

func (s *Server) FindOneUser(ctx context.Context, req *pb.FindOneUserRequest) (*pb.FindOneUserResponse, error) {
	s.Logger.Info("FindOneUser request received for", zap.Uint64("ID", req.Id))

	var user models.User
	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		s.Logger.Error("Error finding user", zap.Uint64("id", req.Id), zap.Error(err))
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	s.Logger.Info("Successfully found user", zap.Uint64("id", user.Id))

	return &pb.FindOneUserResponse{
		Status: http.StatusOK,
		Data: &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
		},
	}, nil
}

func (s *Server) FindMe(ctx context.Context, req *pb.FindOneUserRequest) (*pb.FindOneUserResponse, error) {
	s.Logger.Info("FindMe request received", zap.Uint64("id", req.Id))

	var user models.User
	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		s.Logger.Error("Error finding user", zap.Uint64("id", req.Id), zap.Error(err))
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	s.Logger.Info("Successfully found user", zap.Uint64("id", user.Id))

	return &pb.FindOneUserResponse{
		Status: http.StatusOK,
		Data: &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
		},
	}, nil
}

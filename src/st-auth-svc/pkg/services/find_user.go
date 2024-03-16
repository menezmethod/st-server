package services

import (
	"context"
	"net/http"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
)

func (s *Server) FindAllUsers(ctx context.Context, _ *pb.FindAllUsersRequest) (*pb.FindAllUsersResponse, error) {
	users := make([]*pb.User, 0)

	if err := s.H.DB.NewSelect().Model(&users).Column("id", "email", "first_name", "last_name", "role", "created_at").Scan(ctx); err != nil {
		return &pb.FindAllUsersResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	res := new(pb.FindAllUsersResponse)

	for _, r := range users {
		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (s *Server) FindOneUser(ctx context.Context, req *pb.FindOneUserRequest) (*pb.FindOneUserResponse, error) {
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

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
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

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

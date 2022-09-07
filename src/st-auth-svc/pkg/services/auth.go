package services

import (
	"context"
	"github.com/uptrace/bun"
	"log"
	"net/http"
	"time"

	"st-auth-svc/pkg/db"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

type Server struct {
	H   db.DB
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	exists, err := s.H.DB.NewSelect().Model(&user).Where("email = ?", req.Email.Value).Exists(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	if v := req.GetEmail(); v != nil {
		user.Email = v.Value
	}
	if v := req.GetPassword(); v != nil {
		user.Password = utils.HashPassword(v.Value)
	}
	if v := req.GetFullName(); v != nil {
		user.FullName = v.Value
	}
	if v := req.GetRole(); v != nil {
		user.Role = v.Value
	}
	user.TimeRegistered = time.Now()

	if _, err := s.H.DB.NewInsert().Model(&user).Exec(ctx); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Data: &pb.User{
			Id:       user.Id,
			Email:    user.Email,
			FullName: user.FullName,
			Role:     user.Role,
		},
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("email = ?", req.Email.Value).Scan(ctx); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	if v := req.GetEmail(); v != nil {
		user.Email = v.Value
	}
	if v := req.GetPassword(); v != nil {
		user.Password = utils.HashPassword(v.Value)
	}

	if !utils.CheckPasswordHash(req.Password.Value, user.Password) {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	exists, err := s.H.DB.NewSelect().Model(&user).Where("email LIKE ?", req.Email.Value).Exists(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
		Role:   user.Role,
		Id:     user.Id,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var user models.User
	claims, err := s.Jwt.ValidateToken(req.Token.Value)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	exists, err := s.H.DB.NewSelect().Model(&user).Where("email LIKE ?", claims.Email).Exists(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if !exists {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user models.User

	exists, err := s.H.DB.NewSelect().Model(&user).Where("email LIKE ?", req.Email.Value).Exists(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	if req.Id != user.Id && exists {
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	if req.GetEmail().Value != "" {
		user.Email = req.GetEmail().Value
	}

	if req.GetPassword().Value != "" {
		user.Password = utils.HashPassword(req.GetPassword().Value)
	}

	if req.GetFullName().Value != "" {
		user.FullName = req.GetFullName().Value
	}

	if req.GetRole().Value != "" {
		user.Role = req.GetRole().Value
	}

	user.Id = req.GetId()

	if _, err := s.H.DB.NewUpdate().Model(&user).Where("ID = ?", user.Id).Exec(ctx); err != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusCreated,
		Data: &pb.User{
			Id:       user.Id,
			Email:    user.Email,
			FullName: user.FullName,
			Role:     user.Role,
		},
	}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	if _, err := s.H.DB.NewDelete().Model(&models.User{}).Where("ID IN (?)", bun.In(req.Id)).Exec(ctx); err != nil {
		return &pb.DeleteUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteUserResponse{
		Status: http.StatusOK,
	}, nil
}

package services

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"

	"st-auth-svc/pkg/db"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email.Value}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "email already exists",
		}, nil
	}

	if v := req.GetEmail(); v != nil {
		user.Email = v.Value
	}
	if v := req.GetPassword(); v != nil {
		user.Password = v.Value
	}
	if v := req.GetFullName(); v != nil {
		user.FullName = v.Value
	}
	if v := req.GetRole(); v != nil {
		user.Role = v.Value
	}
	if v := req.GetTimeRegistered(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeRegistered: %s%w", err.Error(), errors.New(""))
			return nil, err
		}
		if t := v.AsTime(); !t.IsZero() {
			user.TimeRegistered = v.AsTime()
		}
	}

	if result := s.H.DB.Create(&user); result.Error != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Data: &pb.User{
			Id:             user.Id,
			Email:          user.Email,
			Password:       user.Password,
			FullName:       user.FullName,
			Role:           user.Role,
			TimeRegistered: timestamppb.New(user.TimeRegistered),
		},
	}, nil
}

func (s *Server) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where("email = ?", req.Email.Value).First(&user); result.Error != nil {
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

	if loginDb := s.H.DB.First(&user); loginDb.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  loginDb.Error.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
		Role:   user.Role,
		Id:     user.Id,
	}, nil
}

func (s *Server) Validate(_ context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	var user models.User
	claims, err := s.Jwt.ValidateToken(req.Token.Value)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *Server) UpdateUser(_ context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Email.Value}).First(&user); result.Error == nil {
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

	if v := req.GetTimeRegistered(); v != nil {
		if err := v.CheckValid(); err != nil {
			err = fmt.Errorf("invalid TimeRegistered: %s%w", err.Error(), errors.New(""))
			return nil, err
		}
		if t := v.AsTime(); !t.IsZero() {
			user.TimeRegistered = v.AsTime()
		}
	}
	user.Id = req.GetId()

	if result := s.H.DB.Updates(&user); result.Error != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusCreated,
		Data: &pb.User{
			Id:             user.Id,
			Email:          user.Email,
			FullName:       user.FullName,
			Role:           user.Role,
			TimeRegistered: timestamppb.New(user.TimeRegistered),
		},
	}, nil
}

func (s *Server) DeleteUser(_ context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	if result := s.H.DB.First(&models.User{}, req.Id); result.Error != nil {
		return &pb.DeleteUserResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if result := s.H.DB.Where("ID IN (?)", req.Id).Delete(&models.User{}); result.Error != nil {
		return &pb.DeleteUserResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.DeleteUserResponse{
		Status: http.StatusOK,
	}, nil
}

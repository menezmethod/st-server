package services

import (
	"context"
	"net/http"
	"time"

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

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req == nil {
		return ErrorResponse(http.StatusBadRequest, "Request is empty"), nil
	}

	email := req.GetEmail()
	password := req.GetPassword()

	if email == nil || email.Value == "" {
		return ErrorResponse(http.StatusBadRequest, "Email is required"), nil
	}

	if password == nil || password.Value == "" {
		return ErrorResponse(http.StatusBadRequest, "Password is required"), nil
	}

	firstName := req.GetFirstName()
	lastName := req.GetLastName()

	user := models.User{
		Email:     email.Value,
		Password:  utils.HashPassword(password.Value),
		FirstName: firstName.Value,
		LastName:  lastName.Value,
		Role:      "USER",
		CreatedAt: time.Now(),
	}

	if exists, err := s.userExists(ctx, user.Email); err != nil {
		return ErrorResponse(http.StatusInternalServerError, "Failed to check if email already exists"), err
	} else if exists {
		return ErrorResponse(http.StatusConflict, "Email already exists"), nil
	}

	if err := s.createUser(ctx, &user); err != nil {
		return ErrorResponse(http.StatusInternalServerError, "Failed to create user"), err
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return ErrorResponse(http.StatusInternalServerError, "Failed to generate token"), err
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		Data: &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
			Token:     token,
		},
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req == nil {
		return ErrorResponseLogin(http.StatusBadRequest, "Empty request"), nil
	}

	email := req.GetEmail()
	if email == nil || email.Value == "" {
		return ErrorResponseLogin(http.StatusBadRequest, "Email is required"), nil
	}

	var user models.User
	if err := s.H.DB.NewSelect().Model(&user).Where("email = ?", email.Value).Scan(ctx); err != nil {
		return ErrorResponseLogin(http.StatusNotFound, "User not found"), nil
	}

	match := utils.CheckPasswordHash(req.GetPassword().Value, user.Password)

	if !match {
		return ErrorResponseLogin(http.StatusNotFound, "User not found"), nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Data: &pb.LoginData{
			Token: token,
		},
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	if req == nil || req.Token == nil || req.Token.Value == "" {
		return ErrorResponseValidate(http.StatusBadRequest, "Invalid token"), nil
	}

	var user models.User
	claims, err := s.Jwt.ValidateToken(req.Token.Value)

	if err != nil {
		return ErrorResponseValidate(http.StatusBadRequest, err.Error()), nil
	}

	if exists, err := s.userExists(ctx, claims.Email); err != nil {
		return ErrorResponseValidate(http.StatusInternalServerError, "Failed to validate user"), err
	} else if !exists {
		return ErrorResponseValidate(http.StatusNotFound, "User not found"), nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}

func (s *Server) userExists(ctx context.Context, email string) (bool, error) {
	var user models.User
	return s.H.DB.NewSelect().Model(&user).Where("email = ?", email).Exists(ctx)
}

func (s *Server) createUser(ctx context.Context, user *models.User) error {
	_, err := s.H.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func ErrorResponse(status int, message string) *pb.RegisterResponse {
	return &pb.RegisterResponse{
		Status: uint64(status),
		Error:  message,
	}
}

func ErrorResponseLogin(status int, message string) *pb.LoginResponse {
	return &pb.LoginResponse{
		Status: uint64(status),
		Error:  message,
	}
}

func ErrorResponseValidate(status int, message string) *pb.ValidateResponse {
	return &pb.ValidateResponse{
		Status: uint64(status),
		Error:  message,
	}
}

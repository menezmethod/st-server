package services

import (
	"context"
	"net/http"
	"time"

	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.Logger.Debug("Received Register request")

	if req == nil || req.Email == "" || req.Password == "" {
		return s.generateRegisterResponse(http.StatusBadRequest, "Invalid request parameters", "Request is nil or missing fields", nil, ""), nil
	}
	hashedPassword, err := utils.HashPassword(s.Logger, req.Password, 10)

	user := models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      "USER",
		CreatedAt: time.Now(),
	}

	if exists, err := s.userExists(ctx, user.Email); err != nil {
		return s.generateRegisterResponse(http.StatusInternalServerError, "User existence check failed", err.Error(), nil, ""), err
	} else if exists {
		return s.generateRegisterResponse(http.StatusConflict, "User already exists", "", nil, ""), nil
	}

	if err := s.createUser(ctx, &user); err != nil {
		return s.generateRegisterResponse(http.StatusInternalServerError, "Failed to create user", err.Error(), nil, ""), err
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		return s.generateRegisterResponse(http.StatusInternalServerError, "Token generation failed", err.Error(), nil, ""), err
	}

	return s.generateRegisterResponse(http.StatusCreated, "User registered successfully", "", &user, token), nil
}

func (s *Server) generateRegisterResponse(statusCode int, message, errorDetail string, user *models.User, token string) *pb.RegisterResponse {
	level := "INFO"
	if statusCode >= 400 {
		level = "ERROR"
	}

	response := &pb.RegisterResponse{
		Status:  uint64(statusCode),
		Message: message,
		Error:   errorDetail,
		Level:   level,
	}

	if user != nil && token != "" {
		response.Data = &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			Token:     token,
		}
	}

	utils.LogResponse(s.Logger, "Register", response, int(response.Status))
	return response
}

func (s *Server) userExists(ctx context.Context, email string) (bool, error) {
	var user models.User
	return s.H.DB.NewSelect().Model(&user).Where("email = ?", email).Exists(ctx)
}

func (s *Server) createUser(ctx context.Context, user *models.User) error {
	_, err := s.H.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		utils.LogResponse(s.Logger, "Failed to create user", map[string]interface{}{
			"error": err.Error(),
		}, http.StatusBadRequest)
		return err
	}
	return nil
}

package services

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"st-auth-svc/pkg/db"
	"time"

	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
	Logger *zap.Logger
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.Logger.Debug("Received Login request")

	if req == nil || req.Email == "" {
		response := generateLoginResponse(http.StatusBadRequest, "Invalid request parameters", "", "Empty request or email is required")
		utils.LogResponse(s.Logger, "Login", response, http.StatusBadRequest)
		return response, nil
	}

	var user models.User
	if err := s.H.DB.NewSelect().Model(&user).Where("email = ?", req.Email).Scan(ctx); err != nil {
		response := generateLoginResponse(http.StatusNotFound, "User not found", "", "User not found")
		utils.LogResponse(s.Logger, "Login", response, http.StatusNotFound)
		return response, nil
	}

	if valid, err := utils.CheckPasswordHash(s.Logger, req.Password, user.Password); err != nil {
		response := generateLoginResponse(http.StatusInternalServerError, "Authentication failed", "", "Error checking password hash")
		utils.LogResponse(s.Logger, "Login", response, http.StatusInternalServerError)
		return response, err
	} else if !valid {
		response := generateLoginResponse(http.StatusUnauthorized, "Authentication failed", "", "Incorrect password")
		utils.LogResponse(s.Logger, "Login", response, http.StatusUnauthorized)
		return response, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		response := generateLoginResponse(http.StatusInternalServerError, "Token generation failed", "", "Failed to generate token")
		utils.LogResponse(s.Logger, "Login", response, http.StatusInternalServerError)
		return response, err
	}

	response := generateLoginResponse(http.StatusOK, "Login successful", token, "")
	utils.LogResponse(s.Logger, "Login", response, http.StatusOK)
	return response, nil
}

func generateLoginResponse(statusCode int, message, token, errMessage string) *pb.LoginResponse {
	level := "INFO"
	if errMessage != "" {
		level = "ERROR"
	}
	response := &pb.LoginResponse{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Status:    uint64(statusCode),
		Error:     errMessage,
	}
	if token != "" {
		response.Data = &pb.LoginData{Token: token}
	}
	return response
}

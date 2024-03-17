package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"st-auth-svc/pkg/db"
	"strings"
	"time"

	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
	Logger    *zap.Logger
	Validator *validator.Validate
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.Logger.Debug("Received Login request")

	login := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.Validator.StructPartial(login, "Email", "Password"); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errMsgs := make([]string, 0)
			for _, e := range validationErrors {
				errMsg := fmt.Sprintf("%s is invalid or missing for field %s", e.ActualTag(), e.Field())
				errMsgs = append(errMsgs, errMsg)
			}
			detailedErrMsg := strings.Join(errMsgs, ", ")
			response := s.generateLoginResponse(http.StatusBadRequest, "Invalid request parameters", detailedErrMsg, "")
			s.Logger.Error("Validation failed",
				zap.String("method", "Login"),
				zap.String("error", detailedErrMsg),
			)
			return response, nil
		}
	}

	var user models.User
	if err := s.H.DB.NewSelect().Model(&user).Where("email = ?", req.Email).Scan(ctx); err != nil {
		response := s.generateLoginResponse(http.StatusUnauthorized, "User not found", "", "")
		return response, nil
	}

	if valid, err := utils.CheckPasswordHash(s.Logger, req.Password, user.Password); !valid || err != nil {
		if err != nil {
			response := s.generateLoginResponse(http.StatusInternalServerError, "Error during password check", "", "")
			return response, err
		}
		response := s.generateLoginResponse(http.StatusUnauthorized, "Incorrect password", "", "")
		return response, nil
	}

	token, err := s.Jwt.GenerateToken(user)
	if err != nil {
		response := s.generateLoginResponse(http.StatusInternalServerError, "Failed to generate token", "", "")
		return response, err
	}

	response := s.generateLoginResponse(http.StatusOK, "Login successful", "", token)
	return response, nil
}

func (s *Server) generateLoginResponse(statusCode int, message, errorDetail, token string) *pb.LoginResponse {
	level := utils.GetStatusLevel(statusCode)

	response := &pb.LoginResponse{
		Status:    uint64(statusCode),
		Message:   message,
		Error:     errorDetail,
		Level:     level,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	if token != "" {
		response.Data = &pb.LoginData{Token: token}
	}
	utils.LogResponse(s.Logger, "Login", response, statusCode)
	return response
}

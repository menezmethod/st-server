package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/utils"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	s.Logger.Debug("Received Register request")
	hashedPassword, err := utils.HashPassword(s.Logger, req.Password, 10)

	user := models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
		CreatedAt: time.Now(),
	}

	if err := s.Validator.Struct(user); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errMsgs := make([]string, 0)
			for _, e := range validationErrors {
				errMsg := fmt.Sprintf("%s is invalid or missing for field %s", e.ActualTag(), e.Field())
				errMsgs = append(errMsgs, errMsg)
			}
			detailedErrMsg := strings.Join(errMsgs, ", ")
			response := s.generateRegisterResponse(http.StatusBadRequest, "Invalid request parameters", detailedErrMsg, nil, "")
			s.Logger.Error("Validation failed",
				zap.String("method", "Register"),
				zap.String("error", detailedErrMsg),
			)
			return response, nil
		}
	}

	user.Password = hashedPassword

	if exists, err := s.userExists(ctx, user.Email); err != nil {
		return s.generateRegisterResponse(http.StatusInternalServerError, "User existence check failed", err.Error(), nil, ""), err
	} else if exists {
		return s.generateRegisterResponse(http.StatusConflict, "User already exists", "UserRegistrationConflict", nil, ""), nil
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
	level := utils.GetStatusLevel(statusCode)

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

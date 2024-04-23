package services

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/utils"
)

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	s.Logger.Debug("Received Validate request")

	if req == nil || req.Token == "" {
		return s.createValidateResponse(http.StatusBadRequest, "Invalid token", "Token is required", 0, zap.String("token", req.Token)), nil
	}

	s.Logger.Debug("Validating token", zap.String("token", req.Token))

	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return s.createValidateResponse(http.StatusBadRequest, "Token validation failed", err.Error(), 0, zap.String("token", req.Token), zap.Error(err)), nil
	}

	s.Logger.Debug("Token validated successfully", zap.String("token", req.Token))

	exists, err := s.userExists(ctx, claims.Email)
	if err != nil {
		return s.createValidateResponse(http.StatusInternalServerError, "Failed to validate user", "Database error", 0, zap.String("email", claims.Email), zap.Error(err)), err
	} else if !exists {
		return s.createValidateResponse(http.StatusNotFound, "User not found", "", 0, zap.String("email", claims.Email)), nil
	}

	s.Logger.Debug("Validation successful", zap.String("email", claims.Email), zap.Uint64("userId", claims.Id))

	return s.createValidateResponse(http.StatusOK, "Token validated successfully", "", claims.Id), nil
}

func (s *Server) createValidateResponse(statusCode int64, message, errorDetail string, userID uint64, fields ...zap.Field) *pb.ValidateResponse {
	timestamp := time.Now().Format(time.RFC3339)
	level := utils.GetStatusLevel(int(statusCode))

	logger := s.Logger.With(fields...)
	utils.LogResponse(logger, "Validate", &pb.ValidateResponse{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Status:    uint64(statusCode),
		Error:     errorDetail,
		UserId:    userID,
	}, int(statusCode))

	return &pb.ValidateResponse{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Status:    uint64(statusCode),
		Error:     errorDetail,
		UserId:    userID,
	}
}

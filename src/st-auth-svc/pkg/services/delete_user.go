package services

import (
	"context"
	"fmt"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-auth-svc/pkg/utils"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	userTerm := "user"

	s.Logger.Debug("Received DeleteUser request for ", zap.Strings("IDs", req.Id))

	if len(req.Id) == 0 {
		errMessage := "No ID provided for deletion"
		errLevel := "WARNING"
		response := generateDeleteUserResponse(http.StatusBadRequest, errLevel, errMessage, "At least one ID is required", 0)
		utils.LogResponse(s.Logger, "DeleteUser", response, http.StatusBadRequest)
		return response, nil
	}

	placeholders := strings.Repeat("?,", len(req.Id)-1) + "?"
	query := fmt.Sprintf("ID IN (%s)", placeholders)

	result, err := s.H.DB.NewDelete().Model((*models.User)(nil)).Where(query, utils.ToInterfaceSlice(req.Id)...).Exec(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("Failed to delete %s.", userTerm)
		response := generateDeleteUserResponse(http.StatusInternalServerError, "ERROR", errMessage, err.Error(), 0)
		utils.LogResponse(s.Logger, "DeleteUser", response, http.StatusInternalServerError)
		return response, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errMessage := "Error checking affected rows after deletion."
		response := generateDeleteUserResponse(http.StatusInternalServerError, "ERROR", errMessage, err.Error(), 0)
		utils.LogResponse(s.Logger, "DeleteUser", response, http.StatusInternalServerError)
		return response, err
	}

	if rowsAffected > 1 {
		userTerm = "users"
	}

	if rowsAffected == 0 {
		errMessage := fmt.Sprintf("No matching users found for provided IDs.")
		errLevel := "WARNING"
		response := generateDeleteUserResponse(http.StatusBadRequest, errLevel, errMessage, "Ensure the provided IDs are correct.", 0)
		utils.LogResponse(s.Logger, "DeleteUser", response, http.StatusBadRequest)
		return response, nil
	}

	response := generateDeleteUserResponse(http.StatusOK, "INFO", fmt.Sprintf("Successfully deleted %d %s.", rowsAffected, userTerm), "", rowsAffected)
	utils.LogResponse(s.Logger, "DeleteUser", response, http.StatusOK)
	return response, nil
}

func generateDeleteUserResponse(status int, level, message, error string, rowsAffected int64) *pb.DeleteUserResponse {
	return &pb.DeleteUserResponse{
		Status:       uint64(status),
		Level:        level,
		Message:      message,
		Error:        error,
		RowsAffected: uint64(rowsAffected),
	}
}

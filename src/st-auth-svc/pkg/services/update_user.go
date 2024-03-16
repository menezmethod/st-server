package services

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"st-auth-svc/pkg/models"
	"st-auth-svc/pkg/pb"
	"st-auth-svc/pkg/utils"
)

func (s *Server) getUserByID(ctx context.Context, id uint64) (models.User, error) {
	user := models.User{}
	err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", id).Scan(ctx)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user models.User
	var dbRes models.User

	err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx)
	if err != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	user.Email = getValueOrDefault(req.GetEmail())
	user.Password = getPasswordOrDefault(req.GetPassword())
	user.FirstName = getValueOrDefault(req.GetFirstName())
	user.LastName = getValueOrDefault(req.GetLastName())
	user.Bio = getValueOrDefault(req.GetBio())
	user.Role = getValueOrDefault(req.GetRole())
	user.Id = req.GetId()

	result, err := s.H.DB.NewUpdate().Model(&user).ExcludeColumn("created_at").Where("ID = ?", user.Id).Exec(ctx)
	if err != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to update user: " + err.Error(),
		}, err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return &pb.UpdateUserResponse{
			Status: http.StatusNoContent,
			Error:  "No user updated",
		}, nil
	}

	dbRes, err = s.getUserByID(ctx, req.Id)
	if err != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to retrieve updated user data: " + err.Error(),
		}, err
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusOK,
		Data: &pb.User{
			Id:        dbRes.Id,
			Email:     dbRes.Email,
			FirstName: dbRes.FirstName,
			LastName:  dbRes.LastName,
			Bio:       dbRes.Bio,
			Role:      dbRes.Role,
			CreatedAt: dbRes.CreatedAt.String(),
		},
	}, nil
}

func getValueOrDefault(value *wrapperspb.StringValue) string {
	if value != nil {
		return value.Value
	}
	return ""
}

func getPasswordOrDefault(value *wrapperspb.StringValue) string {
	if value != nil {
		return utils.HashPassword(value.Value)
	}
	return ""
}

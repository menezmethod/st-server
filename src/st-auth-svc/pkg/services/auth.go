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
	H   db.Handler
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
	if v := req.GetFirstName(); v != nil {
		user.FirstName = v.Value
	}
	if v := req.GetLastName(); v != nil {
		user.LastName = v.Value
	}
	if v := req.GetRole(); v != nil {
		user.Role = "USER"
	}

	user.CreatedAt = time.Now()

	if _, err := s.H.DB.NewInsert().Model(&user).Exec(ctx); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

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
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("email = ?", req.Email.Value).Scan(ctx); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.GetPassword().Value, user.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
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

func (s *Server) FindAllUsers(ctx context.Context, _ *pb.FindAllUsersRequest) (*pb.FindAllUsersResponse, error) {
	users := make([]*pb.User, 0)

	if err := s.H.DB.NewSelect().Model(&users).Column("id", "email", "first_name", "last_name", "role", "created_at").Scan(ctx); err != nil {
		return &pb.FindAllUsersResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	res := new(pb.FindAllUsersResponse)

	for _, r := range users {
		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (s *Server) FindOneUser(ctx context.Context, req *pb.FindOneUserRequest) (*pb.FindOneUserResponse, error) {
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.FindOneUserResponse{
		Status: http.StatusOK,
		Data: &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
		},
	}, nil
}

func (s *Server) FindMe(ctx context.Context, req *pb.FindOneUserRequest) (*pb.FindOneUserResponse, error) {
	var user models.User

	if err := s.H.DB.NewSelect().Model(&user).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneUserResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.FindOneUserResponse{
		Status: http.StatusOK,
		Data: &pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Bio:       user.Bio,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
		},
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user models.User
	var dbRes models.User

	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		log.Fatalln(err)
	}

	if req.GetEmail().Value != dbRes.Email {
		exists, err := s.H.DB.NewSelect().Model(&user).Where("email LIKE ?", req.Email.Value).Exists(ctx)
		if err != nil {
			log.Fatalln(err)
		} else if exists {
			return &pb.UpdateUserResponse{
				Status: http.StatusConflict,
				Error:  "email already exists",
			}, nil
		}
	}

	if req.GetEmail() == nil || req.GetEmail().String() == "" {
		user.Email = dbRes.Email
	} else {
		user.Email = req.GetEmail().Value
	}
	if req.GetPassword() == nil || req.GetPassword().String() == "" {
		user.Password = dbRes.Password
	} else {
		user.Password = utils.HashPassword(req.GetPassword().Value)
	}
	if req.GetFirstName() == nil || req.GetFirstName().String() == "" {
		user.FirstName = dbRes.FirstName
	} else {
		user.FirstName = req.GetFirstName().Value
	}
	if req.GetLastName() == nil || req.GetLastName().String() == "" {
		user.LastName = dbRes.LastName
	} else {
		user.LastName = req.GetLastName().Value
	}
	if req.GetBio() == nil || req.GetBio().String() == "" && req.Email.String() == "" {
		user.Bio = dbRes.Bio
	} else {
		user.Bio = req.GetBio().Value
	}
	if req.GetRole() == nil || req.GetRole().String() == "" {
		user.Role = dbRes.Role
	} else {
		user.Role = req.GetRole().Value
	}
	user.Id = req.GetId()

	if _, err := s.H.DB.NewUpdate().Model(&user).ExcludeColumn("created_at").Where("ID = ?", user.Id).Exec(ctx); err != nil {
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		log.Fatalln(err)
	}

	return &pb.UpdateUserResponse{
		Status: http.StatusCreated,
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

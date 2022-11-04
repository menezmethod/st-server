package services

import (
	"context"
	"github.com/uptrace/bun"
	"net/http"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"time"
)

type Server struct {
	H db.DB
	pb.JournalServiceServer
}

func (s *Server) CreateJournal(ctx context.Context, req *pb.CreateJournalRequest) (*pb.CreateJournalResponse, error) {
	var journal models.Journal
	//TODO This needs validation
	journal.Name = req.GetName()
	journal.Description = req.GetDescription()
	journal.CreatedAt = time.Now()
	journal.StartDate = req.GetStartDate()
	journal.EndDate = req.GetEndDate()
	journal.CreatedBy = req.GetCreatedBy()
	journal.UsersSubscribed = req.GetUsersSubscribed()

	if _, err := s.H.DB.NewInsert().Model(&journal).Exec(ctx); err != nil {
		return &pb.CreateJournalResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateJournalResponse{
		Status: http.StatusCreated,
		Data: &pb.Journal{
			Id:              journal.ID,
			Name:            journal.Name,
			Description:     journal.Description,
			StartDate:       journal.StartDate,
			EndDate:         journal.EndDate,
			CreatedAt:       journal.CreatedAt.String(),
			CreatedBy:       journal.CreatedBy,
			UsersSubscribed: journal.UsersSubscribed,
		},
	}, nil
}

func (s *Server) EditJournal(ctx context.Context, req *pb.EditJournalRequest) (*pb.EditJournalResponse, error) {
	var dbRes models.Journal
	var journal models.Journal

	if req.GetName() != "" {
		journal.Name = req.GetName()
	}
	if req.GetDescription() != "" {
		journal.Description = req.GetDescription()
	}
	if req.GetStartDate() != "" {
		journal.StartDate = req.GetStartDate()
	}
	if req.GetEndDate() != "" {
		journal.EndDate = req.GetEndDate()
	}
	if req.GetCreatedBy() != "" {
		journal.CreatedBy = req.GetCreatedBy()
	}
	if req.GetUsersSubscribed() != nil {
		journal.UsersSubscribed = req.GetUsersSubscribed()
	}

	journal.ID = req.GetId()

	if _, err := s.H.DB.NewUpdate().Model(&journal).Where("ID = ?", journal.ID).Exec(ctx); err != nil {
		return &pb.EditJournalResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.EditJournalResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.EditJournalResponse{
		Status: http.StatusCreated,
		Data: &pb.EditJournalData{
			Id:              req.Id,
			Name:            dbRes.Name,
			Description:     dbRes.Description,
			StartDate:       dbRes.StartDate,
			EndDate:         dbRes.EndDate,
			CreatedBy:       dbRes.CreatedBy,
			UsersSubscribed: dbRes.UsersSubscribed,
		},
	}, nil
}

func (s *Server) FindAllJournals(ctx context.Context, _ *pb.FindAllJournalsRequest) (*pb.FindAllJournalsResponse, error) {
	journals := make([]*pb.Journal, 0)

	if err := s.H.DB.NewSelect().Model(&journals).Column("id", "name", "description", "created_at", "start_date", "end_date", "created_by", "users_subscribed").Scan(ctx); err != nil {
		return &pb.FindAllJournalsResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	res := new(pb.FindAllJournalsResponse)

	for _, r := range journals {
		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (s *Server) FindOneJournal(ctx context.Context, req *pb.FindOneJournalRequest) (*pb.FindOneJournalResponse, error) {
	var journal models.Journal

	if err := s.H.DB.NewSelect().Model(&journal).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneJournalResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.Journal{
		Id:              journal.ID,
		Name:            journal.Name,
		Description:     journal.Description,
		StartDate:       journal.StartDate,
		EndDate:         journal.EndDate,
		CreatedBy:       journal.CreatedBy,
		CreatedAt:       journal.CreatedAt.String(),
		UsersSubscribed: journal.UsersSubscribed,
	}

	return &pb.FindOneJournalResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DeleteJournal(ctx context.Context, req *pb.DeleteJournalRequest) (*pb.DeleteJournalResponse, error) {
	if _, err := s.H.DB.NewDelete().Model(&models.Journal{}).Where("ID IN (?)", bun.In(req.Id)).Exec(ctx); err != nil {
		return &pb.DeleteJournalResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteJournalResponse{
		Status: http.StatusOK,
	}, nil
}

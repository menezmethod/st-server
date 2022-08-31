package services

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.TradeServiceServer
}

func (s *Server) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	var trade models.Trade

	trade.Comments = req.Comments
	trade.Direction = req.Direction
	trade.EntryPrice = req.EntryPrice
	trade.ExitPrice = req.ExitPrice
	trade.Instrument = req.Instrument
	trade.Market = req.Market
	trade.Outcome = req.Outcome
	trade.Quantity = req.Quantity
	trade.StopLoss = req.StopLoss
	trade.Strategy = req.Strategy
	trade.TakeProfit = req.TakeProfit
	trade.TimeClosed = req.TimeClosed.AsTime()
	trade.TimeExecuted = req.TimeExecuted.AsTime()

	if result := s.H.DB.Create(&trade); result.Error != nil {
		return &pb.CreateTradeResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateTradeResponse{
		Status: http.StatusCreated,
		Id:     trade.ID,
	}, nil
}

// EditTrade TODO PUT is working fine but we need this to work with PATCH instead
func (s *Server) EditTrade(ctx context.Context, req *pb.EditTradeRequest) (*pb.EditTradeResponse, error) {
	if result := s.H.DB.Model(models.Trade{}).Where("ID = ?", req.Id).Updates(models.TradePatch{
		ID:           &req.Id,
		Comments:     &req.Comments,
		Direction:    &req.Direction,
		EntryPrice:   &req.EntryPrice,
		ExitPrice:    &req.ExitPrice,
		Instrument:   &req.Instrument,
		Market:       &req.Market,
		Outcome:      &req.Outcome,
		Quantity:     &req.Quantity,
		StopLoss:     &req.StopLoss,
		Strategy:     &req.Strategy,
		TakeProfit:   &req.TakeProfit,
		TimeClosed:   req.TimeClosed.AsTime(),
		TimeExecuted: req.TimeExecuted.AsTime(),
	}); result.Error != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	var db models.Trade

	if result := s.H.DB.First(&db, req.Id); result.Error != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.EditTradeResponse{
		Status: http.StatusCreated,
		Data: &pb.EditTradeData{
			Id:           req.Id,
			Comments:     db.Comments,
			Direction:    db.Direction,
			EntryPrice:   db.EntryPrice,
			ExitPrice:    db.ExitPrice,
			Instrument:   db.Instrument,
			Market:       db.Market,
			Outcome:      db.Outcome,
			Quantity:     db.Quantity,
			StopLoss:     db.StopLoss,
			Strategy:     db.Strategy,
			TakeProfit:   db.TakeProfit,
			TimeClosed:   timestamppb.New(db.TimeClosed),
			TimeExecuted: timestamppb.New(db.TimeExecuted),
		},
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var trade models.Trade

	if result := s.H.DB.First(&trade, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:           trade.ID,
		Comments:     trade.Comments,
		Direction:    trade.Direction,
		EntryPrice:   trade.EntryPrice,
		ExitPrice:    trade.ExitPrice,
		Instrument:   trade.Instrument,
		Market:       trade.Market,
		Outcome:      trade.Outcome,
		Quantity:     trade.Quantity,
		StopLoss:     trade.StopLoss,
		Strategy:     trade.Strategy,
		TakeProfit:   trade.TakeProfit,
		TimeClosed:   timestamppb.New(trade.TimeClosed),
		TimeExecuted: timestamppb.New(trade.TimeExecuted),
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	if result := s.H.DB.First(&models.Trade{}, req.Id); result.Error != nil {
		return &pb.DeleteResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	s.H.DB.Where("ID IN (?)", req.Id).Delete(&models.Trade{})

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

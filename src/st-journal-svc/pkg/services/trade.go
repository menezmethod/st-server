package services

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"strconv"
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
		Instrument:   trade.Instrument,
		Market:       trade.Market,
		Outcome:      trade.Outcome,
		Strategy:     trade.Strategy,
		TimeClosed:   timestamppb.New(trade.TimeClosed),
		TimeExecuted: timestamppb.New(trade.TimeExecuted),
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var log models.TradeDeleteLog

	if result := s.H.DB.First(&models.Trade{}, req.Id); result.Error != nil {
		return &pb.DeleteResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	s.H.DB.Where("ID IN (?)", req.Id).Delete(&models.Trade{})

	for i := range req.Id {
		log.TradeId, _ = strconv.ParseUint(req.Id[i], 10, 32)
	}

	s.H.DB.Create(&log)

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

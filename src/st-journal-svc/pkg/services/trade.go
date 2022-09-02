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

func (s *Server) CreateTrade(_ context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	var trade models.Trade
	//TODO This needs validation
	trade.Comments = req.GetComments().Value
	trade.Direction = req.GetDirection().Value
	trade.EntryPrice = req.GetEntryPrice().Value
	trade.ExitPrice = req.GetExitPrice().Value
	trade.Instrument = req.GetInstrument().Value
	trade.Market = req.GetMarket().Value
	trade.Outcome = req.GetOutcome().Value
	trade.Quantity = req.GetQuantity().Value
	trade.StopLoss = req.GetStopLoss().Value
	trade.Strategy = req.GetStrategy().Value
	trade.TakeProfit = req.GetTakeProfit().Value
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

func (s *Server) EditTrade(_ context.Context, req *pb.EditTradeRequest) (*pb.EditTradeResponse, error) {
	var trade models.Trade

	if req.GetComments().Value != "" {
		trade.Comments = req.GetComments().Value
	}
	if req.GetDirection().GetValue() != "" {
		trade.Comments = req.GetDirection().GetValue()
	}
	if req.GetEntryPrice().Value != 0 {
		trade.EntryPrice = req.GetEntryPrice().Value
	}
	if req.GetExitPrice().Value != 0 {
		trade.ExitPrice = req.GetExitPrice().Value
	}
	if req.GetInstrument().Value != "" {
		trade.Instrument = req.GetInstrument().Value
	}
	if req.GetMarket().Value != "" {
		trade.Market = req.GetMarket().Value
	}
	if req.GetOutcome().Value != "" {
		trade.Outcome = req.GetOutcome().Value
	}
	if req.GetQuantity().Value != 0 {
		trade.Quantity = req.GetQuantity().Value
	}
	if req.GetStopLoss().Value != 0 {
		trade.StopLoss = req.GetStopLoss().Value
	}
	if req.GetStrategy().Value != "" {
		trade.Strategy = req.GetStrategy().Value
	}
	if req.GetTakeProfit().Value != 0 {
		trade.TakeProfit = req.GetTakeProfit().Value
	}
	if req.GetTimeClosed() != nil {
		trade.TimeClosed = req.GetTimeClosed().AsTime()
	}
	if req.GetTimeExecuted() != nil {
		trade.TimeExecuted = req.GetTimeExecuted().AsTime()
	}

	trade.ID = req.GetId()

	if result := s.H.DB.Updates(&trade); result.Error != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	var dbRes models.Trade

	if result := s.H.DB.First(&trade, req.Id); result.Error != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.EditTradeResponse{
		Status: http.StatusCreated,
		Data: &pb.EditTradeData{
			Id:           req.Id,
			Comments:     dbRes.Comments,
			Direction:    dbRes.Direction,
			EntryPrice:   dbRes.EntryPrice,
			ExitPrice:    dbRes.ExitPrice,
			Instrument:   dbRes.Instrument,
			Market:       dbRes.Market,
			Outcome:      dbRes.Outcome,
			Quantity:     dbRes.Quantity,
			StopLoss:     dbRes.StopLoss,
			Strategy:     dbRes.Strategy,
			TakeProfit:   dbRes.TakeProfit,
			TimeClosed:   timestamppb.New(dbRes.TimeClosed),
			TimeExecuted: timestamppb.New(dbRes.TimeExecuted),
		},
	}, nil
}

func (s *Server) FindOne(_ context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
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

func (s *Server) Delete(_ context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

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

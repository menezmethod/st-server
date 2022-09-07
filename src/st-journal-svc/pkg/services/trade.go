package services

import (
	"context"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"st-journal-svc/pkg/db"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

type Server struct {
	H db.DB
	pb.TradeServiceServer
}

func (s *Server) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
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

	if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		return &pb.CreateTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateTradeResponse{
		Status: http.StatusCreated,
		Id:     trade.ID,
	}, nil
}

func (s *Server) EditTrade(ctx context.Context, req *pb.EditTradeRequest) (*pb.EditTradeResponse, error) {
	var dbRes models.Trade
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

	if _, err := s.H.DB.NewUpdate().Model(&trade).Where("ID = ?", trade.ID).Exec(ctx); err != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
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

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var trade models.Trade

	if err := s.H.DB.NewSelect().Model(&trade).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
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
	if _, err := s.H.DB.NewDelete().Model(&models.Trade{}).Where("ID IN (?)", bun.In(req.Id)).Exec(ctx); err != nil {
		return &pb.DeleteResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

package services

import (
	"context"
	"github.com/uptrace/bun"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"time"
)

func (s *Server) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	var trade models.Trade

	trade.Comments = req.GetComments()
	trade.Direction = req.GetDirection()
	trade.EntryPrice = req.GetEntryPrice()
	trade.ExitPrice = req.GetExitPrice()
	trade.Journal = req.GetJournal()
	trade.BaseInstrument = req.GetBaseInstrument()
	trade.QuoteInstrument = req.GetQuoteInstrument()
	trade.Market = req.GetMarket()
	trade.Outcome = req.GetOutcome()
	trade.Quantity = req.GetQuantity()
	trade.StopLoss = req.GetStopLoss()
	trade.Strategy = req.GetStrategy()
	trade.TakeProfit = req.GetTakeProfit()
	trade.TimeClosed = req.GetTimeClosed()
	trade.TimeExecuted = req.GetTimeExecuted()
	trade.CreatedAt = time.Now()
	trade.CreatedBy = req.GetCreatedBy()

	if req.GetDirection() == "Long" && req.GetExitPrice()-req.GetEntryPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetDirection() == "Short" && req.GetExitPrice() > 0 && req.GetEntryPrice()-req.GetExitPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetExitPrice() == 0 {
		trade.Outcome = "TBD"
	} else {
		trade.Outcome = "Loss"
	}

	if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		return &pb.CreateTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateTradeResponse{
		Status: http.StatusCreated,
		Data: &pb.Trade{
			Id:              trade.ID,
			Comments:        trade.Comments,
			CreatedAt:       trade.CreatedAt.String(),
			CreatedBy:       trade.CreatedBy,
			Direction:       trade.Direction,
			EntryPrice:      trade.EntryPrice,
			ExitPrice:       trade.ExitPrice,
			Journal:         trade.Journal,
			BaseInstrument:  trade.BaseInstrument,
			QuoteInstrument: trade.QuoteInstrument,
			Market:          trade.Market,
			Outcome:         trade.Outcome,
			Quantity:        trade.Quantity,
			StopLoss:        trade.StopLoss,
			Strategy:        trade.Strategy,
			TakeProfit:      trade.TakeProfit,
			TimeClosed:      trade.TimeClosed,
			TimeExecuted:    trade.TimeExecuted,
		}}, nil
}

func (s *Server) EditTrade(ctx context.Context, req *pb.EditTradeRequest) (*pb.EditTradeResponse, error) {
	var dbRes models.Trade
	var trade models.Trade

	if req.GetComments() != "" {
		trade.Comments = req.GetComments()
	}
	if req.GetDirection() != "" {
		trade.Direction = req.GetDirection()
	}
	if req.GetEntryPrice() != 0 {
		trade.EntryPrice = req.GetEntryPrice()
	}
	if req.GetExitPrice() != 0 {
		trade.ExitPrice = req.GetExitPrice()
	}
	if req.GetJournal() != 0 {
		trade.Journal = req.GetJournal()
	}
	if req.GetBaseInstrument() != "" {
		trade.BaseInstrument = req.GetBaseInstrument()
	}
	if req.GetQuoteInstrument() != "" {
		trade.QuoteInstrument = req.GetQuoteInstrument()
	}
	if req.GetMarket() != "" {
		trade.Market = req.GetMarket()
	}
	if req.GetOutcome() != "" {
		trade.Outcome = req.GetOutcome()
	}
	if req.GetQuantity() != 0 {
		trade.Quantity = req.GetQuantity()
	}
	if req.GetStopLoss() != 0 {
		trade.StopLoss = req.GetStopLoss()
	}
	if req.GetStrategy() != "" {
		trade.Strategy = req.GetStrategy()
	}
	if req.GetTakeProfit() != 0 {
		trade.TakeProfit = req.GetTakeProfit()
	}
	if req.GetTimeClosed() != "" {
		trade.TimeClosed = req.GetTimeClosed()
	}
	if req.GetTimeExecuted() != "" {
		trade.TimeExecuted = req.GetTimeExecuted()
	}

	trade.ID = req.GetId()

	if req.GetDirection() == "Long" && req.GetExitPrice()-req.GetEntryPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetDirection() == "Short" && req.GetExitPrice() > 0 && req.GetEntryPrice()-req.GetExitPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetExitPrice() == 0 {
		trade.Outcome = "TBD"
	} else {
		trade.Outcome = "Loss"
	}

	if _, err := s.H.DB.NewUpdate().Model(&trade).Column("base_instrument", "comments", "direction",
		"entry_price", "exit_price", "journal", "market", "outcome", "quantity", "quote_instrument", "stop_loss",
		"strategy", "take_profit", "time_closed", "time_executed").Where("ID = ?", trade.ID).Exec(ctx); err != nil {
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
		Data: &pb.Trade{
			Id:              req.Id,
			Comments:        dbRes.Comments,
			Direction:       dbRes.Direction,
			EntryPrice:      dbRes.EntryPrice,
			ExitPrice:       dbRes.ExitPrice,
			Journal:         dbRes.Journal,
			BaseInstrument:  dbRes.BaseInstrument,
			QuoteInstrument: dbRes.QuoteInstrument,
			Market:          dbRes.Market,
			Outcome:         dbRes.Outcome,
			Quantity:        dbRes.Quantity,
			StopLoss:        dbRes.StopLoss,
			Strategy:        dbRes.Strategy,
			TakeProfit:      dbRes.TakeProfit,
			TimeClosed:      dbRes.TimeClosed,
			TimeExecuted:    dbRes.TimeExecuted,
		},
	}, nil
}

func (s *Server) FindAllTrades(ctx context.Context, _ *pb.FindAllTradesRequest) (*pb.FindAllTradesResponse, error) {
	trades := make([]*pb.Trade, 0)

	if err := s.H.DB.NewSelect().Model(&trades).Column("id", "base_instrument", "quote_instrument",
		"comments", "direction", "entry_price", "exit_price", "journal", "market", "outcome", "quantity",
		"stop_loss", "strategy", "take_profit", "time_executed", "time_closed", "created_at", "created_by").Scan(ctx); err != nil {
		return &pb.FindAllTradesResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	res := new(pb.FindAllTradesResponse)

	for _, r := range trades {
		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (s *Server) FindOneTrade(ctx context.Context, req *pb.FindOneTradeRequest) (*pb.FindOneTradeResponse, error) {
	var trade models.Trade

	if err := s.H.DB.NewSelect().Model(&trade).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneTradeResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.Trade{
		Id:              trade.ID,
		Comments:        trade.Comments,
		CreatedBy:       trade.CreatedBy,
		Direction:       trade.Direction,
		EntryPrice:      trade.EntryPrice,
		ExitPrice:       trade.ExitPrice,
		Journal:         trade.Journal,
		BaseInstrument:  trade.BaseInstrument,
		QuoteInstrument: trade.QuoteInstrument,
		Market:          trade.Market,
		Outcome:         trade.Outcome,
		Quantity:        trade.Quantity,
		StopLoss:        trade.StopLoss,
		Strategy:        trade.Strategy,
		TakeProfit:      trade.TakeProfit,
		TimeClosed:      trade.TimeClosed,
		TimeExecuted:    trade.TimeExecuted,
	}

	return &pb.FindOneTradeResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DeleteTrade(ctx context.Context, req *pb.DeleteTradeRequest) (*pb.DeleteTradeResponse, error) {
	if _, err := s.H.DB.NewDelete().Model(&models.Trade{}).Where("ID IN (?)", bun.In(req.Id)).Exec(ctx); err != nil {
		return &pb.DeleteTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteTradeResponse{
		Status: http.StatusOK,
	}, nil
}

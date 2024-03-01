package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"time"
)

func validateTrade(trade *models.Trade) string {
	timeExecuted, _ := time.Parse(time.RFC3339, trade.TimeExecuted)

	if trade.EntryPrice <= 0 {
		return "Entry Price must be greater than 0"
	}
	if trade.ExitPrice < 0 {
		return "Exit Price cannot be negative"
	}
	if trade.Quantity <= 0 {
		return "Quantity must be greater than 0"
	}
	if trade.StopLoss < 0 {
		return "Stop Loss cannot be negative"
	}
	if trade.TakeProfit < 0 {
		return "Take Profit cannot be negative"
	}
	if trade.Journal == 0 {
		return "Journal ID must be provided"
	}
	if trade.BaseInstrument == "" || trade.QuoteInstrument == "" {
		return "Both Base Instrument and Quote Instrument must be provided"
	}
	if trade.Market == "" {
		return "Market must be provided"
	}
	if trade.Strategy == "" {
		return "Strategy must be provided"
	}
	if trade.TimeClosed != "" {
		timeClosed, errClosed := time.Parse(time.RFC3339, trade.TimeClosed)
		if errClosed != nil {
			return "Time Closed must be a valid date"
		}
		if timeExecuted.After(timeClosed) {
			return "Time Executed cannot be after Time Closed"
		}
	}
	return ""
}

func populateTradeFromRequest(req *pb.CreateTradeRequest) models.Trade {
	return models.Trade{
		Comments:        req.GetComments(),
		Direction:       req.GetDirection(),
		EntryPrice:      req.GetEntryPrice(),
		ExitPrice:       req.GetExitPrice(),
		Journal:         req.GetJournal(),
		BaseInstrument:  req.GetBaseInstrument(),
		QuoteInstrument: req.GetQuoteInstrument(),
		Market:          req.GetMarket(),
		Outcome:         req.GetOutcome(),
		Quantity:        req.GetQuantity(),
		StopLoss:        req.GetStopLoss(),
		Strategy:        req.GetStrategy(),
		TakeProfit:      req.GetTakeProfit(),
		TimeClosed:      req.GetTimeClosed(),
		TimeExecuted:    req.GetTimeExecuted(),
		CreatedAt:       time.Now(),
		CreatedBy:       req.GetCreatedBy(),
	}
}

func createTradeResponse(trade models.Trade) *pb.CreateTradeResponse {
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
		},
	}
}

func (s *Server) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	trade := populateTradeFromRequest(req)
	if errMsg := validateTrade(&trade); errMsg != "" {
		return &pb.CreateTradeResponse{
			Status: http.StatusBadRequest,
			Error:  errMsg,
		}, nil
	}

	if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		return &pb.CreateTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return createTradeResponse(trade), nil
}

package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
	"time"

	"go.uber.org/zap"
)

func determineOutcome(req *pb.UpdateTradeRequest, trade *models.Trade) {
	if req.GetDirection() == "Long" && req.GetExitPrice()-req.GetEntryPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetDirection() == "Short" && req.GetExitPrice() > 0 && req.GetEntryPrice()-req.GetExitPrice() > 0.01 {
		trade.Outcome = "Win"
	} else if req.GetExitPrice() == 0 {
		trade.Outcome = "TBD"
	} else {
		trade.Outcome = "Loss"
	}
}

func updateTradeFieldsFromRequest(req *pb.UpdateTradeRequest, trade *models.Trade) {
	trade.Comments = req.GetComments()
	trade.Direction = req.GetDirection()
	trade.EntryPrice = req.GetEntryPrice()
	trade.ExitPrice = req.GetExitPrice()
	trade.Journal = req.GetJournal()
	trade.BaseInstrument = req.GetBaseInstrument()
	trade.QuoteInstrument = req.GetQuoteInstrument()
	trade.Market = req.GetMarket()
	trade.Quantity = req.GetQuantity()
	trade.StopLoss = req.GetStopLoss()
	trade.Strategy = req.GetStrategy()
	trade.TakeProfit = req.GetTakeProfit()
	trade.TimeClosed = req.GetTimeClosed()
	trade.TimeExecuted = req.GetTimeExecuted()
	trade.ID = req.GetId()
}

func createUpdateTradeResponse(trade models.Trade, status uint64, errorMessage string) *pb.UpdateTradeResponse {
	timestamp := time.Now().Format(time.RFC3339)
	level := "INFO"
	message := "Trade updated successfully"

	if errorMessage != "" {
		level = "ERROR"
		message = "Failed to update trade"
	}

	return &pb.UpdateTradeResponse{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Status:    status,
		Data: &pb.Trade{
			Id:              trade.ID,
			Comments:        trade.Comments,
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
		Error: errorMessage,
	}
}

func (s *Server) UpdateTrade(ctx context.Context, req *pb.UpdateTradeRequest) (*pb.UpdateTradeResponse, error) {
	s.Logger.Debug("Received UpdateTrade request", zap.Uint64("ID", req.GetId()))

	var trade models.Trade
	updateTradeFieldsFromRequest(req, &trade)
	determineOutcome(req, &trade)

	errorMsg := utils.ValidateTrade(&trade)
	if errorMsg != "" {
		s.Logger.Error("Validation failed for UpdateTrade", zap.Uint64("ID", req.GetId()), zap.String("error", errorMsg))
		return createUpdateTradeResponse(models.Trade{}, http.StatusBadRequest, errorMsg), nil
	}

	if _, err := s.H.DB.NewUpdate().Model(&trade).Where("ID = ?", trade.ID).Exec(ctx); err != nil {
		s.Logger.Error("Failed to update trade", zap.Uint64("ID", trade.ID), zap.Error(err))
		return createUpdateTradeResponse(models.Trade{}, http.StatusConflict, err.Error()), nil
	}

	var dbRes models.Trade
	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.GetId()).Scan(ctx); err != nil {
		s.Logger.Error("Failed to fetch updated trade", zap.Uint64("ID", req.GetId()), zap.Error(err))
		return createUpdateTradeResponse(models.Trade{}, http.StatusNotFound, err.Error()), nil
	}

	resp := createUpdateTradeResponse(dbRes, http.StatusCreated, "")
	utils.LogResponse(s.Logger, "UpdateTrade", resp, int(resp.Status))

	return resp, nil
}

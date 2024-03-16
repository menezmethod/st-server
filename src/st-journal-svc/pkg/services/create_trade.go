package services

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
	"time"
)

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

func createTradeResponse(trade models.Trade, status uint64) *pb.CreateTradeResponse {
	return &pb.CreateTradeResponse{
		Timestamp: time.Now().Format(time.RFC1123),
		Level:     "INFO",
		Message:   "Trade created successfully",
		Status:    status,
		Data: &pb.Trade{
			Id:              trade.ID,
			Comments:        trade.Comments,
			CreatedAt:       trade.CreatedAt.Format(time.RFC1123),
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
	s.Logger.Debug("Received CreateTrade request")

	trade := populateTradeFromRequest(req)
	errorMsg := utils.ValidateTrade(&trade)

	var resp *pb.CreateTradeResponse
	if errorMsg != "" {
		s.Logger.Error("Validation failed for CreateTrade", zap.String("error", errorMsg))
		resp = &pb.CreateTradeResponse{
			Timestamp: time.Now().Format(time.RFC1123),
			Level:     utils.GetStatusLevel(http.StatusBadRequest),
			Message:   errorMsg,
			Status:    http.StatusBadRequest,
			Error:     errorMsg,
		}
	} else if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		s.Logger.Error("Failed to insert trade", zap.Error(err))
		resp = &pb.CreateTradeResponse{
			Timestamp: time.Now().Format(time.RFC1123),
			Level:     utils.GetStatusLevel(http.StatusInternalServerError),
			Message:   "Failed to insert trade",
			Status:    http.StatusInternalServerError,
			Error:     err.Error(),
		}
	} else {
		resp = createTradeResponse(trade, http.StatusCreated)
	}

	utils.LogResponse(s.Logger, "CreateTrade", resp, int(resp.Status))

	return resp, nil
}

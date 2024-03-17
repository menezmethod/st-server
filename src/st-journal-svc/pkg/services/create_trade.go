package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
	"st-journal-svc/pkg/utils"
	"strings"
	"time"
)

func (s *Server) CreateTrade(ctx context.Context, req *pb.CreateTradeRequest) (*pb.CreateTradeResponse, error) {
	log := s.Logger.With(
		zap.String("action", "CreateTrade"),
		zap.Time("requestTime", time.Now()),
	)

	log.Debug("Received CreateTrade request")

	trade, err := s.populateTradeFromRequest(req)
	if err != nil {
		log.Error("Validation failed for CreateTrade", zap.String("error", err.Error()))
		return s.generateTradeResponse(http.StatusBadRequest, "Validation failed", err.Error(), nil), nil
	}

	if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		log.Error("Failed to insert trade", zap.Error(err))
		return s.generateTradeResponse(http.StatusInternalServerError, "Failed to insert trade", err.Error(), nil), nil
	}

	log.Info("Trade created successfully", zap.Any("trade", trade))
	return s.generateTradeResponse(http.StatusCreated, "Trade created successfully", "", &trade), nil
}

func (s *Server) populateTradeFromRequest(req *pb.CreateTradeRequest) (models.Trade, error) {
	trade := models.Trade{
		BaseInstrument:  req.GetBaseInstrument(),
		Direction:       req.GetDirection(),
		EntryPrice:      req.GetEntryPrice(),
		ExitPrice:       req.GetExitPrice(),
		Journal:         req.GetJournal(),
		Market:          req.GetMarket(),
		Outcome:         req.GetOutcome(),
		Quantity:        req.GetQuantity(),
		QuoteInstrument: req.GetQuoteInstrument(),
		StopLoss:        req.GetStopLoss(),
		Strategy:        req.GetStrategy(),
		TakeProfit:      req.GetTakeProfit(),
		TimeClosed:      req.GetTimeClosed(),
		TimeExecuted:    req.GetTimeExecuted(),
		CreatedAt:       time.Now(),
		CreatedBy:       req.GetCreatedBy(),
	}

	if err := s.Validator.Struct(trade); err != nil {
		var errMsgs []string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, e := range validationErrors {
				errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid for field %s", e.ActualTag(), e.Field()))
			}
		}
		return models.Trade{}, fmt.Errorf(strings.Join(errMsgs, ", "))
	}
	return trade, nil
}

func (s *Server) generateTradeResponse(statusCode int, message, errorDetail string, trade *models.Trade) *pb.CreateTradeResponse {
	resp := &pb.CreateTradeResponse{
		Timestamp: time.Now().String(),
		Level:     utils.GetStatusLevel(statusCode),
		Message:   message,
		Status:    uint64(statusCode),
	}

	if errorDetail != "" {
		resp.Error = errorDetail
	}

	if trade != nil {
		resp.Data = &pb.Trade{
			Id:              trade.ID,
			BaseInstrument:  trade.BaseInstrument,
			QuoteInstrument: trade.QuoteInstrument,
			Market:          trade.Market,
			Strategy:        trade.Strategy,
			EntryPrice:      trade.EntryPrice,
			ExitPrice:       trade.ExitPrice,
			Quantity:        trade.Quantity,
			StopLoss:        trade.StopLoss,
			TakeProfit:      trade.TakeProfit,
			Direction:       trade.Direction,
			Outcome:         trade.Outcome,
			TimeExecuted:    trade.TimeExecuted,
			TimeClosed:      trade.TimeClosed,
			CreatedAt:       timestamppb.New(trade.CreatedAt),
			CreatedBy:       trade.CreatedBy,
		}
	}

	utils.LogResponse(s.Logger, "CreateTrade", resp, statusCode)
	return resp
}

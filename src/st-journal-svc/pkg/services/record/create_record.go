package record

import (
	"context"
	"errors"
	"fmt"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/db"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

type Server struct {
	H db.DB
	pb.RecordServiceServer
	Logger    *zap.Logger
	Validator *validator.Validate
}

func (s *Server) CreateRecord(ctx context.Context, req *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	log := s.Logger.With(
		zap.String("action", "CreateRecord"),
		zap.Time("requestTime", time.Now()),
	)

	log.Debug("Received CreateRecord request")

	trade, err := s.populateTradeFromRequest(req)
	if err != nil {
		log.Error("Validation failed for CreateRecord", zap.String("error", err.Error()))
		return s.generateTradeResponse(http.StatusBadRequest, "Validation failed", err.Error(), nil), nil
	}

	if _, err := s.H.DB.NewInsert().Model(&trade).Exec(ctx); err != nil {
		log.Error("Failed to insert trade", zap.Error(err))
		return s.generateTradeResponse(http.StatusInternalServerError, "Failed to insert trade", err.Error(), nil), nil
	}

	log.Info("Record created successfully", zap.Any("trade", trade))
	return s.generateTradeResponse(http.StatusCreated, "Record created successfully", "", &trade), nil
}

func (s *Server) populateTradeFromRequest(req *pb.CreateRecordRequest) (models.Record, error) {
	trade := models.Record{
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
		return models.Record{}, fmt.Errorf(strings.Join(errMsgs, ", "))
	}
	return trade, nil
}

func (s *Server) generateTradeResponse(statusCode int, message, errorDetail string, trade *models.Record) *pb.CreateRecordResponse {
	resp := &pb.CreateRecordResponse{
		Timestamp: time.Now().String(),
		Level:     utils.GetStatusLevel(statusCode),
		Message:   message,
		Status:    uint64(statusCode),
	}

	if errorDetail != "" {
		resp.Error = errorDetail
	}

	if trade != nil {
		resp.Data = &pb.Record{
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

	utils.LogResponse(s.Logger, "CreateRecord", resp, statusCode)
	return resp
}

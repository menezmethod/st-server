package record

import (
	"context"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func determineOutcome(req *pb.UpdateRecordRequest, trade *models.Record) {
	if req.GetDirection() == "Long" && req.GetExitPrice()-req.GetEntryPrice() > 0.01 {
		trade.Outcome = "WIN"
	} else if req.GetDirection() == "Short" && req.GetExitPrice() > 0 && req.GetEntryPrice()-req.GetExitPrice() > 0.01 {
		trade.Outcome = "WIN"
	} else if req.GetExitPrice() == 0 {
		trade.Outcome = "TBD"
	} else {
		trade.Outcome = "LOSS"
	}
}

func UpdateRecordFieldsFromRequest(req *pb.UpdateRecordRequest, trade *models.Record) {
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

func createUpdateRecordResponse(trade models.Record, status uint64, errorMessage string) *pb.UpdateRecordResponse {
	timestamp := time.Now().Format(time.RFC3339)
	level := "INFO"
	message := "Record updated successfully"

	if errorMessage != "" {
		level = "ERROR"
		message = "Failed to update trade"
	}

	return &pb.UpdateRecordResponse{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
		Status:    status,
		Data: &pb.Record{
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

func (s *Server) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (*pb.UpdateRecordResponse, error) {
	s.Logger.Debug("Received UpdateRecord request", zap.Uint64("ID", req.GetId()))

	var trade models.Record
	UpdateRecordFieldsFromRequest(req, &trade)
	determineOutcome(req, &trade)

	errorMsg := utils.ValidateTrade(&trade)
	if errorMsg != "" {
		s.Logger.Error("Validation failed for UpdateRecord", zap.Uint64("ID", req.GetId()), zap.String("error", errorMsg))
		return createUpdateRecordResponse(models.Record{}, http.StatusBadRequest, errorMsg), nil
	}

	if _, err := s.H.DB.NewUpdate().Model(&trade).Where("ID = ?", trade.ID).Exec(ctx); err != nil {
		s.Logger.Error("Failed to update trade", zap.Uint64("ID", trade.ID), zap.Error(err))
		return createUpdateRecordResponse(models.Record{}, http.StatusConflict, err.Error()), nil
	}

	var dbRes models.Record
	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.GetId()).Scan(ctx); err != nil {
		s.Logger.Error("Failed to fetch updated trade", zap.Uint64("ID", req.GetId()), zap.Error(err))
		return createUpdateRecordResponse(models.Record{}, http.StatusNotFound, err.Error()), nil
	}

	resp := createUpdateRecordResponse(dbRes, http.StatusOK, "")
	utils.LogResponse(s.Logger, "UpdateRecord", resp, int(resp.Status))

	return resp, nil
}

package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

func determineOutcome(req *pb.EditTradeRequest, trade *models.Trade) {
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

func updateTradeFieldsFromRequest(req *pb.EditTradeRequest, trade *models.Trade) {
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

func (s *Server) EditTrade(ctx context.Context, req *pb.EditTradeRequest) (*pb.EditTradeResponse, error) {
	var trade models.Trade
	updateTradeFieldsFromRequest(req, &trade)
	determineOutcome(req, &trade)

	errorMsg := validateTrade(&trade)
	if errorMsg != "" {
		return &pb.EditTradeResponse{
			Status: http.StatusBadRequest,
			Error:  errorMsg,
		}, nil
	}

	if _, err := s.H.DB.NewUpdate().Model(&trade).Where("ID = ?", trade.ID).Exec(ctx); err != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	var dbRes models.Trade
	if err := s.H.DB.NewSelect().Model(&dbRes).Where("ID = ?", req.Id).Scan(ctx); err != nil {
		return &pb.EditTradeResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.EditTradeResponse{
		Status: http.StatusCreated,
		Data: &pb.Trade{
			Id:              dbRes.ID,
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

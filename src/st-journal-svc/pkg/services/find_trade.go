package services

import (
	"context"
	"net/http"
	"st-journal-svc/pkg/models"
	"st-journal-svc/pkg/pb"
)

func mapModelTradeToPBTrade(trade models.Trade) *pb.Trade {
	return &pb.Trade{
		Id:              trade.ID,
		BaseInstrument:  trade.BaseInstrument,
		QuoteInstrument: trade.QuoteInstrument,
		Comments:        trade.Comments,
		Direction:       trade.Direction,
		EntryPrice:      trade.EntryPrice,
		ExitPrice:       trade.ExitPrice,
		Journal:         trade.Journal,
		Market:          trade.Market,
		Outcome:         trade.Outcome,
		Quantity:        trade.Quantity,
		StopLoss:        trade.StopLoss,
		Strategy:        trade.Strategy,
		TakeProfit:      trade.TakeProfit,
		TimeExecuted:    trade.TimeExecuted,
		TimeClosed:      trade.TimeClosed,
		CreatedAt:       trade.CreatedAt.String(),
		CreatedBy:       trade.CreatedBy,
	}
}

func (s *Server) FindAllTrades(ctx context.Context, _ *pb.FindAllTradesRequest) (*pb.FindAllTradesResponse, error) {
	var modelTrades []models.Trade
	if err := s.H.DB.NewSelect().Model(&modelTrades).Scan(ctx); err != nil {
		return &pb.FindAllTradesResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	pbTrades := make([]*pb.Trade, len(modelTrades))
	for i, trade := range modelTrades {
		pbTrades[i] = mapModelTradeToPBTrade(trade)
	}

	return &pb.FindAllTradesResponse{Data: pbTrades}, nil
}

func (s *Server) FindOneTrade(ctx context.Context, req *pb.FindOneTradeRequest) (*pb.FindOneTradeResponse, error) {
	var trade models.Trade
	if err := s.H.DB.NewSelect().Model(&trade).Where("id = ?", req.Id).Scan(ctx); err != nil {
		return &pb.FindOneTradeResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.FindOneTradeResponse{
		Status: http.StatusOK,
		Data:   mapModelTradeToPBTrade(trade),
	}, nil
}

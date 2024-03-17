package services

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"

	"go.uber.org/zap"

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
		CreatedAt:       timestamppb.New(trade.CreatedAt),
		CreatedBy:       trade.CreatedBy,
	}
}

func (s *Server) FindAllTrades(ctx context.Context, _ *pb.FindAllTradesRequest) (*pb.FindAllTradesResponse, error) {
	s.Logger.Debug("Received request to find all trades")

	var modelTrades []models.Trade
	if err := s.H.DB.NewSelect().Model(&modelTrades).Scan(ctx); err != nil {
		s.Logger.Error("Failed to retrieve trades", zap.Error(err))
		return &pb.FindAllTradesResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	pbTrades := make([]*pb.Trade, len(modelTrades))
	for i, trade := range modelTrades {
		pbTrades[i] = mapModelTradeToPBTrade(trade)
	}

	s.Logger.Info("Successfully found trades", zap.Int("count", len(pbTrades)))
	return &pb.FindAllTradesResponse{
		Status: http.StatusOK,
		Data:   pbTrades,
	}, nil
}
func (s *Server) FindOneTrade(ctx context.Context, req *pb.FindOneTradeRequest) (*pb.FindOneTradeResponse, error) {
	s.Logger.Debug("Received request to find trade with ID", zap.Uint64("ID", req.Id))

	var trade models.Trade
	if err := s.H.DB.NewSelect().Model(&trade).Where("id = ?", req.Id).Scan(ctx); err != nil {
		s.Logger.Error("Failed to find trade", zap.Uint64("ID", req.Id), zap.Error(err))
		return &pb.FindOneTradeResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	s.Logger.Info("Successfully found trade", zap.Uint64("ID", trade.ID))
	return &pb.FindOneTradeResponse{
		Status: http.StatusOK,
		Data:   mapModelTradeToPBTrade(trade),
	}, nil
}

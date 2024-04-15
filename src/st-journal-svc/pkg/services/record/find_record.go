package record

import (
	"context"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"

	"go.uber.org/zap"
)

func mapModelTradeToPBTrade(trade models.Record) *pb.Record {
	return &pb.Record{
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

func (s *Server) ListRecords(ctx context.Context, _ *pb.FindAllRecordsRequest) (*pb.FindAllRecordsResponse, error) {
	s.Logger.Debug("Received request to find all trades")

	var modelTrades []models.Record
	if err := s.H.DB.NewSelect().Model(&modelTrades).Scan(ctx); err != nil {
		s.Logger.Error("Failed to retrieve trades", zap.Error(err))
		return &pb.FindAllRecordsResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	pbTrades := make([]*pb.Record, len(modelTrades))
	for i, trade := range modelTrades {
		pbTrades[i] = mapModelTradeToPBTrade(trade)
	}

	s.Logger.Info("Successfully found trades", zap.Int("count", len(pbTrades)))
	return &pb.FindAllRecordsResponse{
		Status: http.StatusOK,
		Data:   pbTrades,
	}, nil
}
func (s *Server) GetRecord(ctx context.Context, req *pb.FindOneRecordRequest) (*pb.FindOneRecordResponse, error) {
	s.Logger.Debug("Received request to find trade with ID", zap.Uint64("ID", req.Id))

	var trade models.Record
	if err := s.H.DB.NewSelect().Model(&trade).Where("id = ?", req.Id).Scan(ctx); err != nil {
		s.Logger.Error("Failed to find trade", zap.Uint64("ID", req.Id), zap.Error(err))
		return &pb.FindOneRecordResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	s.Logger.Info("Successfully found trade", zap.Uint64("ID", trade.ID))
	return &pb.FindOneRecordResponse{
		Status: http.StatusOK,
		Data:   mapModelTradeToPBTrade(trade),
	}, nil
}

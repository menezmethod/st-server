package record

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/db"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/util"
)

type Server struct {
	AuthServiceClient pb.AuthServiceClient
	H                 db.DB
	Logger            *zap.Logger
	Validator         *validator.Validate
	pb.RecordServiceServer
}

func (s *Server) CreateRecord(ctx context.Context, req *pb.CreateRecordRequest) (*pb.CreateRecordResponse, error) {
	if s.Logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	log := s.Logger.With(
		zap.String("action", "CreateRecord"),
		zap.Time("requestTime", time.Now()),
	)

	log.Debug("Received CreateRecord request")

	record := populateRecordFromRequest(req)
	errorMsg := util.ValidateRecord(&record)

	var resp *pb.CreateRecordResponse
	if errorMsg != "" {
		log.Error("Validation failed for CreateRecord", zap.String("error", errorMsg))
		resp = createRecordResponse(&record, http.StatusBadRequest, errorMsg)
	} else {
		if _, err := s.H.DB.NewInsert().Model(&record).Exec(ctx); err != nil {
			log.Error("Failed to insert record", zap.Error(err))
			resp = createRecordResponse(&record, http.StatusInternalServerError, "Failed to insert record")
		} else {
			log.Info("Record created successfully", zap.Any("record", record))
			resp = createRecordResponse(&record, http.StatusCreated, "")
		}
	}

	util.LogResponse(s.Logger, "CreateRecord", resp, int(resp.Status))

	return resp, nil
}

func populateRecordFromRequest(req *pb.CreateRecordRequest) models.Record {
	return models.Record{
		BaseInstrument:  req.GetBaseInstrument(),
		Comments:        req.GetComments(),
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
}

func createRecordResponse(record *models.Record, status uint64, errorMessage string) *pb.CreateRecordResponse {
	timestamp := time.Now().Format(time.RFC3339)
	response := &pb.CreateRecordResponse{
		Timestamp: timestamp,
		Level:     util.GetStatusLevel(int(status)),
		Status:    status,
	}

	if errorMessage != "" {
		response.Error = errorMessage
	} else {
		response.Message = "Record created successfully"
		if record.ID != 0 {
			response.Data = &pb.Record{
				Id:              record.ID,
				CreatedAt:       timestamppb.New(record.CreatedAt),
				BaseInstrument:  record.BaseInstrument,
				Comments:        record.Comments,
				Direction:       record.Direction,
				EntryPrice:      record.EntryPrice,
				ExitPrice:       record.ExitPrice,
				Journal:         record.Journal,
				Market:          record.Market,
				Outcome:         record.Outcome,
				Quantity:        record.Quantity,
				QuoteInstrument: record.QuoteInstrument,
				StopLoss:        record.StopLoss,
				Strategy:        record.Strategy,
				TakeProfit:      record.TakeProfit,
				TimeClosed:      record.TimeClosed,
				TimeExecuted:    record.TimeExecuted,
			}
		}
	}

	return response
}

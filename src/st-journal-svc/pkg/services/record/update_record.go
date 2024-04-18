package record

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc/metadata"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/pb"
	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/utils"
)

func determineOutcome(req *pb.UpdateRecordRequest, record *models.Record) {
	if req.GetDirection() == "Long" && req.GetExitPrice()-req.GetEntryPrice() > 0.01 {
		record.Outcome = "WIN"
	} else if req.GetDirection() == "Short" && req.GetExitPrice() > 0 && req.GetEntryPrice()-req.GetExitPrice() > 0.01 {
		record.Outcome = "WIN"
	} else if req.GetExitPrice() == 0 {
		record.Outcome = "TBD"
	} else {
		record.Outcome = "LOSS"
	}
}

func updateRecordFieldsFromRequest(req *pb.UpdateRecordRequest, record *models.Record) {
	record.BaseInstrument = req.GetBaseInstrument()
	record.Comments = req.GetComments()
	record.Direction = req.GetDirection()
	record.EntryPrice = req.GetEntryPrice()
	record.ExitPrice = req.GetExitPrice()
	record.ID = req.GetId()
	record.Journal = req.GetJournal()
	record.LastUpdatedBy = req.GetLastUpdatedBy()
	record.Market = req.GetMarket()
	record.Quantity = req.GetQuantity()
	record.QuoteInstrument = req.GetQuoteInstrument()
	record.StopLoss = req.GetStopLoss()
	record.Strategy = req.GetStrategy()
	record.TakeProfit = req.GetTakeProfit()
	record.TimeClosed = req.GetTimeClosed()
	record.TimeExecuted = req.GetTimeExecuted()
}

func createUpdateRecordResponse(record *models.Record, status uint64, errorMessage string) *pb.UpdateRecordResponse {
	timestamp := time.Now().Format(time.RFC3339)
	response := &pb.UpdateRecordResponse{
		Timestamp: timestamp,
		Level:     utils.GetStatusLevel(int(status)),
		Status:    status,
	}

	if errorMessage != "" {
		response.Error = errorMessage
	} else {
		response.Message = "Record updated successfully"
		if record.ID != 0 {
			response.Data = &pb.Record{
				BaseInstrument:  record.BaseInstrument,
				Comments:        record.Comments,
				Direction:       record.Direction,
				EntryPrice:      record.EntryPrice,
				ExitPrice:       record.ExitPrice,
				Id:              record.ID,
				Journal:         record.Journal,
				LastUpdatedBy:   record.LastUpdatedBy,
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

func (s *Server) UpdateRecord(ctx context.Context, req *pb.UpdateRecordRequest) (*pb.UpdateRecordResponse, error) {
	s.Logger.Debug("Received UpdateRecord request", zap.String("RecordID", strconv.FormatUint(req.GetId(), 10)))

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		s.Logger.Warn("No metadata received with request")
		return createUpdateRecordResponse(&models.Record{}, http.StatusBadRequest, "no metadata received with request"), nil
	}

	userIDStrs, ok := md["user-id"]
	if !ok || len(userIDStrs) == 0 {
		s.Logger.Warn("User ID not provided in metadata")
		return createUpdateRecordResponse(&models.Record{}, http.StatusUnauthorized, "user-id not provided in metadata"), nil
	}

	loggedInUserID, err := strconv.ParseUint(userIDStrs[0], 10, 64)
	if err != nil {
		s.Logger.Error("Invalid user ID format", zap.Error(err))
		return createUpdateRecordResponse(&models.Record{}, http.StatusUnauthorized, err.Error()), nil
	}

	s.Logger.Info("Authenticated user", zap.Uint64("UserID", loggedInUserID))

	authRes, err := s.AuthServiceClient.FindOneUser(ctx, &pb.FindOneUserRequest{Id: loggedInUserID})
	if err != nil {
		s.Logger.Error("Failed to get user from auth service", zap.Error(err))
		return createUpdateRecordResponse(&models.Record{}, http.StatusInternalServerError, err.Error()), nil
	}

	if authRes == nil || authRes.Data == nil {
		s.Logger.Error("Invalid response from auth service")
		return createUpdateRecordResponse(&models.Record{}, http.StatusInternalServerError, "invalid response from auth service"), nil
	}

	s.Logger.Info("User role retrieved", zap.String("Role", authRes.Data.Role))

	var existingRecord models.Record
	if err := s.H.DB.NewSelect().Model(&existingRecord).Where("ID = ?", req.GetId()).Scan(ctx); err != nil {
		s.Logger.Error("Failed to retrieve record from database", zap.Error(err))
		return createUpdateRecordResponse(&models.Record{}, http.StatusInternalServerError, "failed to retrieve record"), nil
	}

	if existingRecord.CreatedBy != loggedInUserID && authRes.Data.Role != "ADMIN" {
		s.Logger.Error("Unauthorized attempt to update record", zap.Uint64("RecordID", req.GetId()), zap.Uint64("AttemptedByUserID", loggedInUserID))
		return createUpdateRecordResponse(&models.Record{}, http.StatusForbidden, "unauthorized to update this record"), nil
	}

	record := &models.Record{}
	updateRecordFieldsFromRequest(req, record)
	determineOutcome(req, record)

	errorMsg := utils.ValidateRecord(record)
	if errorMsg != "" {
		s.Logger.Error("Validation failed for UpdateRecord", zap.String("error", errorMsg))
		return createUpdateRecordResponse(&models.Record{}, http.StatusBadRequest, errorMsg), nil
	}

	if _, err := s.H.DB.NewUpdate().Model(&record).Where("ID = ?", record.ID).Exec(ctx); err != nil {
		s.Logger.Error("Failed to update record in database", zap.Error(err))
		return createUpdateRecordResponse(&models.Record{}, http.StatusInternalServerError, "failed to update record"), nil
	}

	s.Logger.Info("Record updated successfully", zap.Uint64("RecordID", record.ID))
	return createUpdateRecordResponse(record, http.StatusOK, ""), nil
}

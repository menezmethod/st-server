package utils

import (
	"encoding/json"

	"github.com/menezmethod/st-server/src/st-journal-svc/pkg/models"
	"go.uber.org/zap"
)

func ValidateJournal(journal *models.Journal) string {
	if journal.Name == "" {
		return "Name cannot be empty"
	}
	if journal.Description == "" {
		return "Description cannot be empty"
	}
	return ""
}

func ValidateRecord(trade *models.Record) string {
	if trade.EntryPrice <= 0 {
		return "Entry Price must be greater than 0"
	}
	if trade.ExitPrice < 0 {
		return "Exit Price cannot be negative"
	}
	if trade.Quantity <= 0 {
		return "Quantity must be greater than 0"
	}
	if trade.StopLoss < 0 {
		return "Stop Loss cannot be negative"
	}
	if trade.TakeProfit < 0 {
		return "Take Profit cannot be negative"
	}
	if trade.Journal == 0 {
		return "Journal ID must be provided"
	}
	if trade.BaseInstrument == "" || trade.QuoteInstrument == "" {
		return "Both Base Instrument and Quote Instrument must be provided"
	}
	if trade.Market == "" {
		return "Market must be provided"
	}
	if trade.Strategy == "" {
		return "Strategy must be provided"
	}
	return ""
}

func LogResponse(logger *zap.Logger, action string, resp interface{}, statusCode int) {
	respJSON, err := json.Marshal(resp)
	if err != nil {
		logger.Error("Failed to marshal response",
			zap.String("action", action),
			zap.Error(err),
		)
		return
	}

	level := GetStatusLevel(statusCode)

	switch level {
	case "ERROR":
		logger.Error(action+" response",
			zap.String("response", string(respJSON)),
			zap.Int("status", statusCode),
		)
	case "WARNING":
		logger.Warn(action+" response",
			zap.String("response", string(respJSON)),
			zap.Int("status", statusCode),
		)
	case "INFO":
		logger.Info(action+" response",
			zap.String("response", string(respJSON)),
			zap.Int("status", statusCode),
		)
	default:
		logger.Debug(action+" response",
			zap.String("response", string(respJSON)),
			zap.Int("status", statusCode),
		)
	}
}

func GetStatusLevel(statusCode int) string {
	switch {
	case statusCode >= 500:
		return "ERROR"
	case statusCode >= 400:
		return "WARNING"
	case statusCode >= 200:
		return "INFO"
	default:
		return "DEBUG"
	}
}

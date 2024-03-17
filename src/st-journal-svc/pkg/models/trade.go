package models

import (
	"time"
)

type Trade struct {
	ID              uint64    `json:"id" bun:",pk,autoincrement"`
	BaseInstrument  string    `json:"baseInstrument" validate:"required"`
	Comments        string    `json:"comments" validate:"max=500"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       string    `json:"createdBy" validate:"required"`
	Direction       string    `json:"direction" validate:"required,oneof=BUY SHORT"`
	EntryPrice      float32   `json:"entryPrice" validate:"gt=0"`
	ExitPrice       float32   `json:"exitPrice" validate:"gte=0"`
	Journal         uint64    `json:"journal" validate:"required"`
	Market          string    `json:"market" validate:"required"`
	Outcome         string    `json:"outcome" validate:"omitempty,oneof=WIN LOSS BREAK-EVEN"`
	Quantity        float32   `json:"quantity" validate:"gt=0"`
	QuoteInstrument string    `json:"quoteInstrument" validate:"required"`
	StopLoss        float32   `json:"stopLoss" validate:"gte=0"`
	Strategy        string    `json:"strategy" validate:"required"`
	TakeProfit      float32   `json:"takeProfit" validate:"gte=0"`
	TimeClosed      string    `json:"timeClosed" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	TimeExecuted    string    `json:"timeExecuted" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
}

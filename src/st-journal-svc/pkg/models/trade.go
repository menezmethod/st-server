package models

import (
	"time"
)

type Trade struct {
	ID              uint64    `json:"id,omitempty" bun:",pk,autoincrement"`
	BaseInstrument  string    `json:"baseInstrument,omitempty"`
	Comments        string    `json:"comments,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
	Direction       string    `json:"direction,omitempty,omitempty"`
	EntryPrice      float32   `json:"entryPrice,omitempty"`
	ExitPrice       float32   `json:"exitPrice,omitempty"`
	Journal         uint64    `json:"journal,omitempty"`
	Market          string    `json:"market,omitempty"`
	Outcome         string    `json:"outcome,omitempty"`
	Quantity        float32   `json:"quantity,omitempty"`
	QuoteInstrument string    `json:"quoteInstrument,omitempty"`
	StopLoss        float32   `json:"stopLoss,omitempty"`
	Strategy        string    `json:"strategy,omitempty"`
	TakeProfit      float32   `json:"takeProfit,omitempty"`
	TimeClosed      string    `json:"timeClosed,omitempty"`
	TimeExecuted    string    `json:"timeExecuted,omitempty"`
}

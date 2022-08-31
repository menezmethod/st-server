package models

import (
	"time"
)

type TradePatch struct {
	ID           *uint64   `json:"id,omitempty" gorm:"primaryKey"`
	Comments     *string   `json:"comments,omitempty"`
	Direction    *string   `json:"direction,omitempty,omitempty"`
	EntryPrice   *float32  `json:"entryPrice,omitempty"`
	ExitPrice    *float32  `json:"exitPrice,omitempty"`
	Instrument   *string   `json:"instrument,omitempty"`
	Market       *string   `json:"market,omitempty"`
	Outcome      *string   `json:"outcome,omitempty"`
	Quantity     *float32  `json:"quantity,omitempty"`
	StopLoss     *float32  `json:"stopLoss,omitempty"`
	Strategy     *string   `json:"strategy,omitempty"`
	TakeProfit   *float32  `json:"takeProfit,omitempty"`
	TimeClosed   time.Time `json:"timeClosed,omitempty"`
	TimeExecuted time.Time `json:"timeExecuted,omitempty"`
}

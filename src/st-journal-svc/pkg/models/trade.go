package models

import (
	"time"
)

type Trade struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	Instrument   string    `json:"instrument"`
	Comments     string    `json:"comments"`
	Direction    string    `json:"direction"`
	EntryPrice   float32   `json:"entryPrice"`
	ExitPrice    float32   `json:"exitPrice"`
	Market       string    `json:"market"`
	Outcome      string    `json:"outcome"`
	Quantity     float32   `json:"quantity"`
	StopLoss     float32   `json:"stopLoss"`
	Strategy     string    `json:"strategy"`
	TakeProfit   float32   `json:"takeProfit"`
	TimeClosed   time.Time `json:"timeClosed"`
	TimeExecuted time.Time `json:"timeExecuted"`
}

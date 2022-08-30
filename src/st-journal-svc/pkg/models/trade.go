package models

import (
	"time"
)

type Trade struct {
	ID           uint64    `json:"id" gorm:"primaryKey"`
	Instrument   string    `json:"instrument"`
	Comments     string    `json:"comments"`
	Direction    string    `json:"direction"`
	EntryPrice   uint64    `json:"entryPrice"`
	ExitPrice    uint64    `json:"exitPrice"`
	Market       string    `json:"market"`
	Outcome      string    `json:"outcome"`
	Quantity     uint32    `json:"quantity"`
	StopLoss     uint64    `json:"stopLoss"`
	Strategy     string    `json:"strategy"`
	TakeProfit   uint64    `json:"takeProfit"`
	TimeClosed   time.Time `json:"timeClosed"`
	TimeExecuted time.Time `json:"timeExecuted"`
}

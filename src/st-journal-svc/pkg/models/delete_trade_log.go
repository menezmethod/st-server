package models

type TradeDeleteLog struct {
	Id      uint64 `json:"id"`
	TradeId uint64 `json:"tradeId"`
}

package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Symbol          string  `json:"symbol" gorm:"unique"`
	CompanyName     string  `json:"company_name"`
	TotalShares     int     `json:"total_shares"`       
	AvailableShares int     `json:"available_shares"`   
	IPOPrice        float64 `json:"ipo_price"`
}

type UserStock struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	StockSymbol string `json:"stock_symbol"`
	Quantity int `json:"quantity"`
}

type BuyIPORequest struct {
	UserID   uint   `json:"user_id"`
	Symbol   string `json:"symbol"`
	Quantity int    `json:"quantity"`
}

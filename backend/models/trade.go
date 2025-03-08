package models

import (
	"gorm.io/gorm"
)

type TradeRequest struct {
	BuyerId   int    `json:"buyer_id" binding:"required"`
	SellerId  int    `json:"seller_id" binding:"required"`
	Symbol    string `json:"symbol" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gte=1"`
	BuyPrice  float64 `json:"buy_price" binding:"required"`
}

type TradeResponse struct {
	Success bool   `json:"success" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type Trade struct {
	gorm.Model
	BuyerID   int    `json:"buyer_id" binding:"required"`
	SellerID  int    `json:"seller_id" binding:"required"`
	Symbol    string `json:"symbol" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gte=1"`
}

package db

import (
	"backend/models"
)

func InsertTradeRequest(tradeRequest *models.TradeRequest) error {
	result := DB.Create(tradeRequest)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
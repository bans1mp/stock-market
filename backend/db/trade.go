package db

import (
	"backend/dto"
	"backend/models"
)

func InsertTradeRequest(tradeRequest *models.TradeRequest) error {
	trade := dto.ConvertTradeRequestToTrade(tradeRequest)
	result := DB.Create(trade)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
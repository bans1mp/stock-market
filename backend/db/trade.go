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

func UpdateBalance(tradeRequest *models.TradeRequest) error {
	var buyerBalanceInfo models.UserBalance
	DB.Model(&buyerBalanceInfo).Where("user_id = ?", tradeRequest.BuyerId).First(&buyerBalanceInfo)
	buyerBalanceInfo.ReservedBalance -= tradeRequest.Price * float64(tradeRequest.Quantity)
	// balance add outstanding amount
	res := DB.Model(&buyerBalanceInfo).Updates(buyerBalanceInfo)
	if res.Error != nil {
		return res.Error
	}

	var sellerBalanceInfo models.UserBalance
	DB.Model(&sellerBalanceInfo).Where("user_id = ?", tradeRequest.SellerId).First(&sellerBalanceInfo)
	sellerBalanceInfo.Balance += tradeRequest.Price * float64(tradeRequest.Quantity)
	res = DB.Model(&sellerBalanceInfo).Updates(sellerBalanceInfo)
	if res.Error != nil {
		return res.Error
	}
	
	return nil
}
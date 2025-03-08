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
	// trade hui hai price pr
	// buyer dera tha buyPrice mtlb usne reserved me quantity*buyPrice dala hoga
	// ab uske kate kitne? (buyPrice - price)*quantity ye uske balance me daalo cos extra hai
	buyerBalanceInfo.Balance += (tradeRequest.BuyPrice-tradeRequest.Price) * float64(tradeRequest.Quantity)
	// balance add outstanding amount
	res := DB.Model(&buyerBalanceInfo).Updates(buyerBalanceInfo)
	if res.Error != nil {
		return res.Error
	}

	var sellerBalanceInfo models.UserBalance
	DB.Model(&sellerBalanceInfo).Where("user_id = ?", tradeRequest.SellerId).First(&sellerBalanceInfo)
	sellerBalanceInfo.Balance += tradeRequest.Price * float64(tradeRequest.Quantity)
	// trade price pr hui hai seller never reserved any money so no need to update
	res = DB.Model(&sellerBalanceInfo).Updates(sellerBalanceInfo)
	if res.Error != nil {
		return res.Error
	}
	
	return nil
}
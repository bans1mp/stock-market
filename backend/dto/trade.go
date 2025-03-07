package dto

import (
	"backend/models"
	trade "backend/proto"
)

// ConvertOrderRequestToProto converts an OrderRequest struct to a Proto OrderRequest
func ConvertTradeRequestFromProto(req *trade.TradeRequest) *models.TradeRequest {
	return &models.TradeRequest{
		BuyerId:   int(req.BuyerId),
		SellerId:  int(req.SellerId),
		Symbol:    req.Symbol,
		Price:     req.Price,
		Quantity:  int(req.Quantity),
	}
}

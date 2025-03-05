package dto

import (
	"backend/models"
	trade "backend/proto"
)

// ConvertOrderRequestToProto converts an OrderRequest struct to a Proto OrderRequest
func ConvertOrderRequestToProto(req *models.OrderRequest) *trade.OrderRequest {
	return &trade.OrderRequest{
		UserId:    int32(req.UserId),
		Symbol:    req.Symbol,
		Price:     req.Price,
		Quantity:  int32(req.Quantity),
		OrderType: req.OrderType,
	}
}

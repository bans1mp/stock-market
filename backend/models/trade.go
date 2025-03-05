package models

type OrderRequest struct {
	UserId    int     `json:"user_id" binding:"required"`
	Symbol    string  `json:"symbol" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	OrderType string  `json:"order_type" binding:"required"`
}

type OrderResponse struct {
	Status string `json:"status" binding:"required"`
	Id     int    `json:"id" binding:"required"`
}
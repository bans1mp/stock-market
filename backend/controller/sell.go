package controller

import (
	"backend/db"
	"backend/dto"
	"backend/models"
	"backend/pkg/matching_engine"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PlaceSellOrder(c *gin.Context) {
	orderRequest := models.OrderRequest{}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateSellOrder(&orderRequest)
	if err != nil {
		return 
	}

	// Deduct stocks from the seller
	var userStock models.UserStock
	db.DB.Where("user_id = ? AND stock_symbol = ?", orderRequest.UserId, orderRequest.Symbol).First(&userStock)
	userStock.Quantity -= orderRequest.Quantity
	db.DB.Save(&userStock)

	// convert input into protobuf message
	orderRequestProto := dto.ConvertOrderRequestToProto(&orderRequest)
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	orderResponseProto, err := matching_engine.Client.PlaceOrder(ctx, orderRequestProto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": orderResponseProto.Message, "status": orderResponseProto.Success})
}

func validateSellOrder(orderRequest *models.OrderRequest) error {
	if orderRequest.Symbol == "" {
		return errors.New("symbol not present")
	}
	if orderRequest.OrderType != "sell" {
		return errors.New("wrong order type")
	}

	// Check if the user has enough stocks to sell
	var userStock models.UserStock
	err := db.DB.Where("user_id = ? AND stock_symbol = ?", orderRequest.UserId, orderRequest.Symbol).First(&userStock).Error
	if err != nil {
		return errors.New("user does not own this stock")
	}

	if userStock.Quantity < orderRequest.Quantity {
		return errors.New("not enough stocks to sell")
	}
	return nil
}
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

func PlaceBuyOrder(c *gin.Context) {
	orderRequest := models.OrderRequest{}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validateBuyOrder(&orderRequest)
	if err != nil {
		return 
	}
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

func validateBuyOrder(orderRequest *models.OrderRequest) error {
	if orderRequest.Symbol == "" {
		return errors.New("symbol not present")
	}
	if orderRequest.OrderType != "buy" {
		return errors.New("wrong order type")
	}
	moneySpent := orderRequest.Price * float64(orderRequest.Quantity)

	var userBalanceInfo models.UserBalance
	db.DB.Where("user_id = ?", orderRequest.UserId).First(&userBalanceInfo)

	if moneySpent > userBalanceInfo.Balance {
		return errors.New("insufficient balance")
	}

	userBalanceInfo.Balance -= moneySpent
	userBalanceInfo.ReservedBalance += moneySpent

	db.DB.Model(&userBalanceInfo).Updates(userBalanceInfo)
	return nil
}
package controller

import (
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
	return nil
}
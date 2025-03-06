package controller

import (
	"backend/dto"
	"backend/models"
	"backend/pkg/matching_engine"
	"context"
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
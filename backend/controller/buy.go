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

func PlaceBuyOrder(c *gin.Context) {
	orderRequest := models.OrderRequest{}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert input into protobuf message
	orderRequestProto := dto.ConvertOrderRequestToProto(&orderRequest)
	
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	matching_engine.Client.PlaceOrder(ctx, orderRequestProto)
	
}
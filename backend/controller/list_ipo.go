package controller

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListIPO(c *gin.Context) {
	req := models.ListIPORequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

	db.DB.Create(&models.Stock{
		Symbol: req.Symbol,
		TotalShares: req.Quantity,
		AvailableShares: req.Quantity,
		IPOPrice: req.IPOPrice,
	})

	c.JSON(http.StatusOK, gin.H{"message": "IPO listed successfully"})
}
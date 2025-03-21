package controller

import (
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStocks(c *gin.Context) {
	var stocks []models.Stock

	// Query all stocks
	if err := db.DB.Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stocks"})
		return
	}
	fmt.Print(stocks)
	// Return stocks as JSON
	c.JSON(http.StatusOK, stocks)
}

func GetStockPrices(c *gin.Context) {
	symbol := c.Param("symbol")
	// fetch all entries for the given stock symbol
	var stockPrices []models.StockPrice
	db.DB.Where("stock_symbol = ?", symbol).Order("created_at DESC").Find(&stockPrices)
	c.JSON(http.StatusOK, stockPrices)

}
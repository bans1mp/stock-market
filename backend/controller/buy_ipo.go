package controller

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyIPO(c *gin.Context) {
	req := models.BuyIPORequest{}

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

	var user models.User
    if err := db.DB.Where("id = ?", req.UserID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    // Get stock info
    var stock models.Stock
    if err := db.DB.Where("symbol = ?", req.Symbol).First(&stock).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
        return
    }

   // Calculate total cost
    totalCost := float64(req.Quantity) * stock.IPOPrice

    // Check if user has enough balance
    var userBalance models.UserBalance
    if err := db.DB.Where("user_id = ?", req.UserID).First(&userBalance).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User balance not found"})
        return
    }

    if userBalance.Balance < totalCost {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
        return
    }

    // Check if enough IPO shares are available
    if stock.AvailableShares < req.Quantity {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough IPO shares available"})
        return
    }

    // Deduct IPO shares
    stock.AvailableShares -= req.Quantity
    db.DB.Save(&stock)

    // Deduct balance from user
    userBalance.Balance -= totalCost
    db.DB.Save(&userBalance)

    // Check if the user already owns this stock
    var userStock models.UserStock
    if err := db.DB.Where("user_id = ? AND stock_symbol = ?", req.UserID, stock.ID).First(&userStock).Error; err != nil {
        // User does not own the stock yet, create a new entry
        userStock = models.UserStock{
            UserID:   req.UserID,
            StockSymbol:  stock.Symbol,
            Quantity: req.Quantity,
        }
        db.DB.Create(&userStock)
    } else {
        // User already owns the stock, update quantity
        userStock.Quantity += req.Quantity
        db.DB.Save(&userStock)
    }

    c.JSON(http.StatusOK, gin.H{"message": "IPO shares bought successfully"})
}

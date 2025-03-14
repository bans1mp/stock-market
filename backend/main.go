package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	"backend/pkg/matching_engine"
	"backend/service"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	fmt.Println("Database connected: ", database)

	matching_engine.Init()
	defer matching_engine.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Change to match your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	go service.StartServer()

	auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    // auth.GET("/profile", controller.GetProfile)
	auth.POST("/buy", controller.PlaceBuyOrder)
	auth.POST("/sell", controller.PlaceSellOrder)
	auth.POST("/buy-ipo", controller.BuyIPO)
	auth.GET("/get-stocks", controller.GetStocks)
	r.POST("list-ipo", controller.ListIPO)
	// introduce scripts to trade automatically

	r.Run(":8080")

}
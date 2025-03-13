package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	"backend/pkg/matching_engine"
	"backend/service"
	"fmt"
	"log"

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

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	go service.StartServer()

	auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    // auth.GET("/profile", controller.GetProfile)
	auth.POST("/buy", controller.PlaceBuyOrder)
	auth.POST("/sell", controller.PlaceSellOrder)

	// to do check if the person has the stocks that he wants to sell
	// introduce portfolio
	// check how companies list their stock (introduce IPO initially the company has x shares and it wants to sell)
	// introduce scripts to trade automatically

	r.Run(":8080")

}
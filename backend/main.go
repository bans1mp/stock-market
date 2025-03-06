package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	"backend/pkg/matching_engine"
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

	auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    // auth.GET("/profile", controller.GetProfile)
	auth.POST("/buy", controller.PlaceBuyOrder)
	auth.POST("/sell", controller.PlaceSellOrder)

	r.Run(":8080")

}
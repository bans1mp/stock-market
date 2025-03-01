package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
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

	r := gin.Default()

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    // auth.GET("/profile", controller.GetProfile)

	r.Run(":8080")

}
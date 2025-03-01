package main

import (
	"backend/controller"
	"backend/db"
	"backend/middleware"
	trade "backend/proto"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	fmt.Println("Database connected: ", database)

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error creating GRPC connection to server")
	}
	defer conn.Close()

	client := trade.NewMatchingEngineClient(conn)

	req := &trade.OrderRequest{
		UserId: 1,
		Symbol: "AAPL",
		Price: 200,
		Quantity: 10,
		OrderType: "BUY",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.PlaceOrder(ctx, req)
	if err != nil {
		log.Fatalf("Error calling PlaceOrder: %v", err)
	}

	fmt.Println("Server Response:", res)

	r := gin.Default()

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())
    // auth.GET("/profile", controller.GetProfile)

	r.Run(":8080")

}
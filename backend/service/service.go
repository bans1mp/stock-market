package service

import (
	"backend/db"
	"backend/dto"
	"backend/models"
	pb "backend/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Backend struct {
	pb.UnimplementedBackendServer
}

func (s *Backend) ExecuteTrade(ctx context.Context, request *pb.TradeRequest) (*pb.TradeResponse, error) {
	tradeRequest := dto.ConvertTradeRequestFromProto(request)

	err := db.InsertTradeRequest(tradeRequest) 
	if err != nil {
		return &pb.TradeResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	err = db.UpdateBalance(tradeRequest)
	if err != nil {
		return &pb.TradeResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	err = db.UpdateStockPrices(tradeRequest)
	if err != nil {
		return &pb.TradeResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Add stocks to the buyer
	var userStock models.UserStock
	if err := db.DB.Where("user_id = ? AND stock_id = ?", tradeRequest.BuyerId, tradeRequest.Symbol).First(&userStock).Error; err != nil {
		// Buyer does not own this stock yet, create a new entry
		userStock = models.UserStock{
			UserID:   uint(tradeRequest.BuyerId),
			StockSymbol:  tradeRequest.Symbol,
			Quantity: int(tradeRequest.Quantity),
		}
		db.DB.Create(&userStock)
	} else {
		// Buyer already owns the stock, update quantity
		userStock.Quantity += int(tradeRequest.Quantity)
		db.DB.Save(&userStock)
	}

	return &pb.TradeResponse{
		Success: true,
		Message: "Trade executed successfully",
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBackendServer(s, &Backend{})

	log.Println("gRPC server listening on port 8081")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
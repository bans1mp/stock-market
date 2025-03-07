package service

import (
	"backend/dto"
	"backend/db"
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
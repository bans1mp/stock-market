package service

import (
	pb "backend/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type TradeServer struct {
	pb.UnimplementedBackendServer
}

func (s *TradeServer) ExecuteTrade(ctx context.Context, request *pb.TradeRequest) (*pb.TradeResponse, error) {
	return &pb.TradeResponse{
		Success: true,
		Message: "Trade executed successfully",
	}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()


	log.Println("gRPC server listening on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
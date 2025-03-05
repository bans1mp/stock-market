package matching_engine

import (
	trade "backend/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Connection *grpc.ClientConn
var Client trade.MatchingEngineClient

func Init() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error creating GRPC connection to server")
	}

	Connection = conn 
	Client = trade.NewMatchingEngineClient(conn)
}

func Close() {
	if Connection != nil {
		Connection.Close() 
	}
}
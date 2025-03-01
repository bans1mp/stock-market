#include <iostream>
#include <memory>
#include <string>
#include <grpcpp/grpcpp.h>
#include "../proto/trade.grpc.pb.h"  // Include gRPC generated headers
#include <hiredis/hiredis.h>

using namespace std ;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
using trade::OrderRequest;
using trade::OrderResponse;
using trade::MatchingEngine;

// Implement the MatchingEngine gRPC service
class MatchingEngineImpl final : public MatchingEngine::Service {
public:
    Status PlaceOrder(ServerContext* context, const OrderRequest* request, OrderResponse* response) override {
        cout << "Received order from user: " << request->user_id() << endl;
        cout << "Order type: " << request->order_type() << ", Quantity: " << request->quantity() << endl;

        // process order here

        response->set_message("Order Placed");
        response->set_success(true);

        cout << "Order processed\n";
        return Status::OK;
    }
};

// Start the gRPC server
void RunServer() {
    string server_address("0.0.0.0:50051");
    MatchingEngineImpl service;

    ServerBuilder builder;
    builder.AddListeningPort(server_address, grpc::InsecureServerCredentials());
    builder.RegisterService(&service);
    unique_ptr<Server> server(builder.BuildAndStart());

    cout << "Matching Engine Server listening on " << server_address << endl;
    server->Wait();
}

redisContext* connectRedis() {
    redisContext* redis = redisConnect("127.0.0.1", 6379);
    if(redis == NULL || redis->err) {
        cerr<<"Redis connection error!\n"<<redis->err<<"\n" ;
        return NULL ;
    }
    return redis ;
}

int main() {
    redisContext* conn = connectRedis();
    if(conn) {
        cout<<"Connection successful\n" ;
    }
    RunServer();
    return 0;
}

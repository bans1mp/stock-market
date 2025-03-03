#include "server.h"
#include "order_book.h"
#include <iostream>

using namespace std;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;

Status MatchingEngineImpl::PlaceOrder(ServerContext* context, const trade::OrderRequest* request, trade::OrderResponse* response) {
    cout << "Received order from user: " << request->user_id() << endl;
    cout << "Order type: " << request->order_type() << ", Quantity: " << request->quantity() << endl;

    string orderType = request->order_type();
    
    if (orderType == "BUY") {
        executeBuyOrder(request->price(), request->quantity(), request->user_id(), request->symbol());
    } 
    else if (orderType == "SELL") {
        executeSellOrder(request->price(), request->quantity(), request->user_id(), request->symbol());
    } 
    else {
        response->set_message("Invalid Order type");
        response->set_success(false);
    }

    response->set_message("Order Placed");
    response->set_success(true);

    return Status::OK;
}

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

// need to maintain the order book
// for buy orders store sorted set with {x,x} 
// whenever a sell order comes with cost y, check for least x >= y and get that value from the set, check in the map of that value to find another sorted set
// sorted on the basis of time, keep subtracting till you get required condition
// similarly with sell orders
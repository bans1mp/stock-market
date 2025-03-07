// client.cpp
#include "client.h"
#include <iostream>

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using trade::Backend;
using trade::TradeRequest;
using trade::TradeResponse;

TradeClient::TradeClient(std::shared_ptr<Channel> channel)
    : stub_(Backend::NewStub(channel)) {}

void TradeClient::ExecuteTrade(int buyer_id, int seller_id, const std::string& symbol, double price, int quantity) {
    TradeRequest request;
    request.set_buyer_id(buyer_id);
    request.set_seller_id(seller_id);
    request.set_symbol(symbol);
    request.set_price(price);
    request.set_quantity(quantity);

    TradeResponse response;
    ClientContext context;

    // Call ExecuteTrade on the gRPC server
    Status status = stub_->ExecuteTrade(&context, request, &response);

    if (status.ok()) {
        std::cout << "Trade executed: " << response.message() << std::endl;
    } else {
        std::cerr << "RPC failed: " << status.error_message() << std::endl;
    }
}

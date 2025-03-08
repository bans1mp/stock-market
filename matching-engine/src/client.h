// client.h
#ifndef CLIENT_H
#define CLIENT_H

#include <grpcpp/grpcpp.h>
#include "trade.grpc.pb.h"

class TradeClient {
public:
    TradeClient(std::shared_ptr<grpc::Channel> channel);
    void ExecuteTrade(int buyer_id, int seller_id, const std::string& symbol, double price, int quantity, int buyPrice);

private:
    std::unique_ptr<trade::Backend::Stub> stub_;
};

#endif // CLIENT_H

#ifndef SERVER_H
#define SERVER_H

#include <grpcpp/grpcpp.h>
#include "../proto/trade.grpc.pb.h"

class MatchingEngineImpl final : public trade::MatchingEngine::Service {
public:
    grpc::Status PlaceOrder(grpc::ServerContext* context, const trade::OrderRequest* request, trade::OrderResponse* response) override;
};

void RunServer();

#endif // SERVER_H

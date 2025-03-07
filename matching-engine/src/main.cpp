#include "server.h"
#include "redis.h"
#include "order_book.h"
#include <iostream>

int main() {
    InitGrpcClient();
    RunServer();
    return 0;
}

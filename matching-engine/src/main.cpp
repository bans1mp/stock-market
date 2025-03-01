#include "server.h"
#include "redis.h"
#include <iostream>

int main() {
    redisContext* conn = connectRedis();
    if(conn) {
        std::cout << "Connected to Redis successfully!" << std::endl;
    }
    RunServer();
    return 0;
}

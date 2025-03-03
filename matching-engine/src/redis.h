#ifndef REDIS_H
#define REDIS_H

#include <hiredis/hiredis.h>

class RedisClient {
private:
    redisContext* conn;
    RedisClient();  // Private constructor for Singleton pattern

public:
    static RedisClient& getInstance();  // Returns the single instance
    redisContext* getConnection();  // Gets the Redis connection
    ~RedisClient();
};

#endif // REDIS_H

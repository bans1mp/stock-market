#include "redis.h"
#include <iostream>

using namespace std;

redisContext* connectRedis() {
    redisContext* redis = redisConnect("127.0.0.1", 6379);
    if (redis == NULL || redis->err) {
        cerr << "Redis connection error! Code: " << (redis ? redis->errstr : "Unknown") << endl;
        return NULL;
    }
    return redis;
}



#include "redis.h"
#include <iostream>

using namespace std;

RedisClient::RedisClient() {
    conn = redisConnect("127.0.0.1", 6379);
    if (conn == NULL || conn->err) {
        cerr << "Connection error: " << (conn ? conn->errstr : "NULL context") << endl;
        conn = nullptr;
    }
}

RedisClient& RedisClient::getInstance() {
    static RedisClient instance;
    return instance;
}

redisContext* RedisClient::getConnection() {
    return conn;
}

RedisClient::~RedisClient() {
    if (conn) {
        redisFree(conn);
    }
}

#include "redis.h"
#include <iostream>

using namespace std;

class RedisClient {
private:
    redisContext* conn ;

    RedisClient(){
        conn = redisConnect("127.0.0.1", 6379);
        if (conn == NULL || conn->err) {
            cerr << "Connection error: " << (conn ? conn->errstr : "NULL context") << std::endl;
            conn = nullptr;
        }
    }

public:
    static RedisClient& getInstance(){
        static RedisClient instance ;
        return instance;
    }

    redisContext* getConnection() {
        return conn;
    }

    ~RedisClient() {
        if(conn) {
            redisFree(conn);
        }
    }
};

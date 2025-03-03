#include "order_book.h"
#include "redis.h"
#include <hiredis/hiredis.h>
#include <iostream>
#include <chrono>

using namespace std;

void executeBuyOrder(double buyPrice, int quantity, int userID, string symbol) {
    cout << "Processing Buy Order: Price = " << buyPrice << ", Quantity = " << quantity << endl;

    auto now = chrono::system_clock::now();
    auto epoch = chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count();

    string sellRedisKey = "sell_orders_" + symbol;
    string buyRedisKey = "buy_orders_" + symbol;
    string value = userID + "_" + to_string(quantity) + '_' + to_string(epoch); 

    redisContext* conn = connectRedis();
    if (conn == NULL || conn->err) {
        printf("Connection error: %s\n", conn ? conn->errstr : "NULL context");
        return ;
    }
    // can break into class

    int reqQuantity = quantity ;

    while(reqQuantity > 0) {
        redisReply* reply = (redisReply*)redisCommand(conn, "ZRANGE %s 0 0 WITHSCORES", sellRedisKey);
        if(reply == NULL || reply -> type != REDIS_REPLY_ARRAY) {
            printf("An error occured\n");
            break ;
        }

        if (reply->elements == 0){
            break ;
        }

        char* order = reply->element[0]->str;
        double price = atof(reply->element[1]->str);

        if (price > buyPrice){
            break ;
        }

        char sellerUserID[50];
        int availableQuantity;
        long timestamp ;
        sscanf(order, "%d_%d_%ld", &userID, &availableQuantity, &timestamp);

        if (availableQuantity > reqQuantity) {
            int remainingQuantity = availableQuantity-reqQuantity;
            char new_order[100];
            snprintf(new_order, sizeof(new_order), "%d_%d_%ld", userID, remainingQuantity, timestamp);
            // send order via protobuf to golang for processing SQL

            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", sellRedisKey, order);
            redisReply *addReply = (redisReply *)redisCommand(conn, "ZADD %s %d %s", sellRedisKey, price, new_order);

            reqQuantity = 0 ;
        } else {
            // send order via protobuf to golang for processing SQL
            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", sellRedisKey, order);
            reqQuantity -= availableQuantity ;
        }

    }

    if(reqQuantity > 0){
        char new_buy_order[100];
        snprintf(new_buy_order, sizeof(new_buy_order), "%d_%d_%ld", userID, reqQuantity, epoch);
        redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", buyRedisKey.c_str(), buyPrice, new_buy_order);
    }

    return;

}

void executeSellOrder(double sellPrice, int quantity, int userID, string symbol) {
    cout << "Processing Sell Order: Price = " << sellPrice << ", Quantity = " << quantity << endl;

    auto now = chrono::system_clock::now();
    auto epoch = chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count();

    string sellRedisKey = "sell_orders_" + symbol;
    string buyRedisKey = "buy_orders_" + symbol;
    string value = to_string(userID) + "_" + to_string(quantity) + '_' + to_string(epoch); 

    redisContext* conn = connectRedis();
    if (conn == NULL || conn->err) {
        printf("Connection error: %s\n", conn ? conn->errstr : "NULL context");
        return;
    }

    int reqQuantity = quantity;

    while (reqQuantity > 0) {
        redisReply* reply = (redisReply*)redisCommand(conn, "ZREVRANGE %s 0 0 WITHSCORES", buyRedisKey.c_str());
        if (reply == NULL || reply->type != REDIS_REPLY_ARRAY) {
            printf("An error occurred\n");
            break;
        }

        if (reply->elements == 0) {
            break;
        }

        char* order = reply->element[0]->str;
        double price = atof(reply->element[1]->str);

        if (price < sellPrice) {
            break;
        }

        int buyerUserID, availableQuantity;
        long timestamp;
        sscanf(order, "%d_%d_%ld", &buyerUserID, &availableQuantity, &timestamp);

        if (availableQuantity > reqQuantity) {
            int remainingQuantity = availableQuantity - reqQuantity;
            char new_order[100];
            snprintf(new_order, sizeof(new_order), "%d_%d_%ld", buyerUserID, remainingQuantity, timestamp);
            // send order via protobuf to golang for processing SQL

            redisReply* delReply = (redisReply*)redisCommand(conn, "ZREM %s %s", buyRedisKey.c_str(), order);
            redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", buyRedisKey.c_str(), price, new_order);

            reqQuantity = 0;
        } else {
            // send order via protobuf to golang for processing SQL
            redisReply* delReply = (redisReply*)redisCommand(conn, "ZREM %s %s", buyRedisKey.c_str(), order);
            reqQuantity -= availableQuantity;
        }
    }

    if (reqQuantity > 0) {
        char new_sell_order[100];
        snprintf(new_sell_order, sizeof(new_sell_order), "%d_%d_%ld", userID, reqQuantity, epoch);
        redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", sellRedisKey.c_str(), sellPrice, new_sell_order);
    }

    return;
}

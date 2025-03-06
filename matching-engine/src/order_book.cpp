#include "order_book.h"
#include "redis.h"
#include <hiredis/hiredis.h>
#include <bits/stdc++.h>
#include <chrono>

using namespace std;

void executeBuyOrder(double buyPrice, int quantity, int userID, string symbol) {
    cout << "Processing Buy Order: Price = " << buyPrice << ", Quantity = " << quantity << endl;

    auto now = chrono::system_clock::now();
    auto epoch = chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count();

    string sellRedisKey = "sell_orders_" + symbol;
    string buyRedisKey = "buy_orders_" + symbol;

    redisContext* conn = RedisClient::getInstance().getConnection();
    if (conn == NULL || conn->err) {
        printf("Connection error: %s\n", conn ? conn->errstr : "NULL context");
        return ;
    }

    int reqQuantity = quantity ;
    vector<string> selfOrders ;

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

        int availableQuantity, sellerUserID;
        long timestamp ;
        sscanf(order, "%ld_%d_%d", &timestamp, &sellerUserID, &availableQuantity);

        if(sellerUserID == userID){
            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", sellRedisKey, order);
            stringstream ss;
            ss << "ZADD " << sellRedisKey << " " << price << " " << order;
            selfOrders.push_back(ss.str().c_str());
            continue;
        }

        if (availableQuantity > reqQuantity) {
            int remainingQuantity = availableQuantity-reqQuantity;
            char new_order[100];
            snprintf(new_order, sizeof(new_order), "%ld_%d_%d", timestamp, userID, remainingQuantity);
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
        snprintf(new_buy_order, sizeof(new_buy_order), "%ld_%d_%d", epoch, userID, reqQuantity);
        redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", buyRedisKey.c_str(), buyPrice, new_buy_order);
    }

    for(auto x : selfOrders){
        redisCommand(conn, x.c_str());
    }

    return;

}

void executeSellOrder(double sellPrice, int quantity, int userID, string symbol) {
    cout << "Processing Sell Order: Price = " << sellPrice << ", Quantity = " << quantity << endl;

    auto now = chrono::system_clock::now();
    auto epoch = chrono::duration_cast<chrono::seconds>(now.time_since_epoch()).count();

    string sellRedisKey = "sell_orders_" + symbol;
    string buyRedisKey = "buy_orders_" + symbol; 

    redisContext* conn = RedisClient::getInstance().getConnection();
    if (conn == NULL || conn->err) {
        printf("Connection error: %s\n", conn ? conn->errstr : "NULL context");
        return;
    }

    int reqQuantity = quantity;
    vector<string> selfOrders ;

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
        sscanf(order, "%ld_%d_%d", &timestamp, &buyerUserID, &availableQuantity);

        if(buyerUserID == userID){
            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", buyRedisKey.c_str(), order);
            stringstream ss;
            ss << "ZADD " << sellRedisKey << " " << price << " " << order;
            selfOrders.push_back(ss.str().c_str());
            continue;
        }

        if (availableQuantity > reqQuantity) {
            int remainingQuantity = availableQuantity - reqQuantity;
            char new_order[100];
            snprintf(new_order, sizeof(new_order), "%ld_%d_%d", timestamp, buyerUserID, remainingQuantity);
            // send order via protobuf to golang for processing SQL

            redisReply* delReply = (redisReply*)redisCommand(conn, "ZREM %s %s", buyRedisKey.c_str(), order);
            redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", buyRedisKey.c_str(), price, new_order);
            cout<<"ZADD %s %f %s"<<buyRedisKey.c_str()<<price<<new_order<<"\n" ;

            reqQuantity = 0;
        } else {
            // send order via protobuf to golang for processing SQL
            redisReply* delReply = (redisReply*)redisCommand(conn, "ZREM %s %s", buyRedisKey.c_str(), order);
            reqQuantity -= availableQuantity;
        }
    }

    if (reqQuantity > 0) {
        char new_sell_order[100];
        snprintf(new_sell_order, sizeof(new_sell_order), "%ld_%d_%d", epoch, userID, reqQuantity);
        redisReply* addReply = (redisReply*)redisCommand(conn, "ZADD %s %f %s", sellRedisKey.c_str(), sellPrice, new_sell_order);
    }

    for(auto x : selfOrders){
        redisCommand(conn, x.c_str());
    }

    return;
}

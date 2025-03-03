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

    string redisKey = "sell_orders_" + symbol;
    string value = userID + "_" + to_string(quantity) + '_' + to_string(epoch); // add timestamp

    redisContext* conn = connectRedis();
    if (conn == NULL || conn->err) {
        printf("Connection error: %s\n", conn ? conn->errstr : "NULL context");
        return ;
    }
    // can break into class

    int reqQuantity = quantity ;

    while(reqQuantity > 0) {
        redisReply* reply = (redisReply*)redisCommand(conn, "ZRANGE %s 0 0 WITHSCORES", redisKey);
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

            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", redisKey, order);
            redisReply *addReply = (redisReply *)redisCommand(conn, "ZADD %s %d %s", redisKey, price, new_order);

            reqQuantity = 0 ;
        } else {
            // send order via protobuf to golang for processing SQL
            redisReply *delReply = (redisReply *)redisCommand(conn, "ZREM %s %s", redisKey, order);
            reqQuantity -= availableQuantity ;
        }

    }

    if(reqQuantity > 0){
        // add the rest of the order to the buy order book
    }

    return;

}

// void executeSellOrder(int price, int quantity) {
//     cout << "Processing Sell Order: Price = " << price << ", Quantity = " << quantity << endl;

//     // Try to find a buy order that satisfies the sell order
//     auto it = buyOrders.lower_bound(price);
//     while (it != buyOrders.begin() && quantity > 0) {
//         --it;
//         int buyPrice = it->first;
//         set<int>& buyTimes = it->second;

//         if (!buyTimes.empty()) {
//             // Remove the earliest placed order
//             buyTimes.erase(buyTimes.begin());
//             quantity--;
//             cout << "Matched with Buy Order at " << buyPrice << endl;
//         }

//         // Remove price level if no orders left
//         if (buyTimes.empty()) {
//             it = buyOrders.erase(it);
//         }
//     }

//     // If there's still remaining quantity, add to sellOrders
//     if (quantity > 0) {
//         sellOrders[price].insert(time(nullptr));
//         cout << "Added remaining sell order to book at " << price << endl;
//     }
// }

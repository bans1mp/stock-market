#include "order_book.h"
#include "redis.h"
#include <hiredis/hiredis.h>
#include <iostream>

using namespace std;

void executeBuyOrder(int price, int quantity, int userID, string symbol) {
    cout << "Processing Buy Order: Price = " << price << ", Quantity = " << quantity << endl;

    string redisKey = "sell_orders_" + symbol;
    string value = userID + ":" + to_string(quantity) ;

    redisContext* conn = connectRedis();

    // iterate from smallert

}

void executeSellOrder(int price, int quantity) {
    cout << "Processing Sell Order: Price = " << price << ", Quantity = " << quantity << endl;

    // Try to find a buy order that satisfies the sell order
    auto it = buyOrders.lower_bound(price);
    while (it != buyOrders.begin() && quantity > 0) {
        --it;
        int buyPrice = it->first;
        set<int>& buyTimes = it->second;

        if (!buyTimes.empty()) {
            // Remove the earliest placed order
            buyTimes.erase(buyTimes.begin());
            quantity--;
            cout << "Matched with Buy Order at " << buyPrice << endl;
        }

        // Remove price level if no orders left
        if (buyTimes.empty()) {
            it = buyOrders.erase(it);
        }
    }

    // If there's still remaining quantity, add to sellOrders
    if (quantity > 0) {
        sellOrders[price].insert(time(nullptr));
        cout << "Added remaining sell order to book at " << price << endl;
    }
}

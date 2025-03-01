#ifndef ORDER_BOOK_H
#define ORDER_BOOK_H

#include <map>
#include <set>
#include <string>

using namespace std;

// Order book structure (maps price to sorted set of order timestamps)
extern map<int, set<int>> buyOrders;
extern map<int, set<int>> sellOrders;

// Function declarations
void executeBuyOrder(int price, int quantity);
void executeSellOrder(int price, int quantity);

#endif // ORDER_BOOK_H

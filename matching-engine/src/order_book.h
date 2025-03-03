#ifndef ORDER_BOOK_H
#define ORDER_BOOK_H

#include <map>
#include <set>
#include <string>

using namespace std;

// Function declarations
void executeBuyOrder(double buyPrice, int quantity, int userID, string symbol);
void executeSellOrder(double buyPrice, int quantity, int userID, string symbol);

#endif // ORDER_BOOK_H

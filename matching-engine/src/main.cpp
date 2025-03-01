#include <hiredis/hiredis.h>
#include<bits/stdc++.h>

using namespace std ;

redisContext* connectRedis() {
    redisContext* redis = redisConnect("127.0.0.1", 6379);
    if(redis == NULL || redis->err) {
        cerr<<"Redis connection error!\n"<<redis->err<<"\n" ;
        return NULL ;
    }
    return redis ;
}

int main() {
    redisContext* conn = connectRedis();
    if(conn) {
        cout<<"Connection successful\n" ;
    }

    // need to maintain the order book
    // for buy orders store sorted set with {x,x} 
    // whenever a sell order comes with cost y, check for least x >= y and get that value from the set, check in the map of that value to find another sorted set
    // sorted on the basis of time, keep subtracting till you get required condition

    // similarly with sell orders
    
    return 0 ;
}


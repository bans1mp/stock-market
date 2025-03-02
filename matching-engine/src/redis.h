#ifndef REDIS_H
#define REDIS_H

#include <hiredis/hiredis.h>
#include <string>

using namespace std ;

// Connect to Redis
redisContext* connectRedis();

#endif // REDIS_H

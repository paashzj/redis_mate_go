#!/bin/bash

nohup $REDIS_HOME/redis-server $REDIS_HOME/redis.conf >$REDIS_HOME/redis.log 2>$REDIS_HOME/redis_error.log &
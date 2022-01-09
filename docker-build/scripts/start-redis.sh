#!/bin/bash

mkdir $REDIS_HOME/logs
nohup $REDIS_HOME/redis-server $REDIS_HOME/redis.conf >>$REDIS_HOME/logs/redis.stdout.log 2>>$REDIS_HOME/logs/redis.stderr.log &


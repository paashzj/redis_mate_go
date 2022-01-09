#!/bin/bash

export REMOTE_MODE=false
mkdir $REDIS_HOME/logs
nohup $REDIS_HOME/mate/redis_mate >>$REDIS_HOME/logs/redis_mate.stdout.log 2>>$REDIS_HOME/logs/redis_mate.stderr.log


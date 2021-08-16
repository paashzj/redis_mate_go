#!/bin/bash

export REMOTE_MODE=false
nohup $REDIS_HOME/mate/redis_mate >$REDIS_HOME/redis_mate.log 2>$REDIS_HOME/redis_mate_error.log
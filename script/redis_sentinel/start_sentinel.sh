#!/bin/bash

nohup ./redis-sentinel redis-sentinel_26479.conf > 26479.log 2>&1  &
nohup ./redis-sentinel redis-sentinel_26480.conf > 26480.log 2>&1  &
nohup ./redis-sentinel redis-sentinel_26481.conf > 26481.log 2>&1  &

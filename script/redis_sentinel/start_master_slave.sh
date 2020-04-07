#!/bin/bash

cd 26379
../redis-server redis.conf
cd ../26380
../redis-server redis.conf
cd ../26381
../redis-server redis.conf


#!/bin/bash


echo "select 1\r\nFLUSHALL" | ./redis-cli -h 127.0.0.1 -p 26379
#echo "select 1\r\nxadd connect_stream * hello world" | ./redis-cli -h 127.0.0.1 -p 26379
# echo "select 1\r\nxadd disconn_stream * hello world" | ./redis-cli -h 127.0.0.1 -p 26379
# echo "select 1\r\nxgroup create disconn_stream disgroup $" | ./redis-cli -h 127.0.0.1 -p 26379
#echo "select 1\r\nxgroup create connect_stream congroup $" | ./redis-cli -h 127.0.0.1 -p 26379




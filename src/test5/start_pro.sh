#!/bin/bash
#./test4 -redisAddr "127.0.01:26481,127.0.01:26480, 127.0.01:26479" -master "master001" -consumer "zoned1"
if [ ! -n "$1" ]; then
	echo "please input msg"
	exit
fi
./test5 -redisAddr "127.0.01:26481,127.0.01:26480, 127.0.01:26479" -master "master001" -msg "$1"

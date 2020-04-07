#!/bin/sh
# prepare
# Usage:./startup.sh

 #init redis-cluster
for port in `echo 6380 6381 6382 6383 6384 6385`
do
	echo    'daemonize yes\n
		 port '$port'\n
		 cluster-node-timeout 5000\n
		 save ""\n
		 cluster-enabled yes\n
		 cluster-config-file '$port'.conf' | redis-server -
done

echo "yes"  | redis-cli --cluster create 127.0.0.1:6380 127.0.0.1:6381 127.0.0.1:6382 127.0.0.1:6383 127.0.0.1:6384 127.0.0.1:6385 --cluster-replicas 1

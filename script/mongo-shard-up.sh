#!/bin/sh
set -ex
# remove test-data dir
rm -rf test-data

# create config server (single member replica set)
mkdir -p test-data/cfg0
mongod --configsvr --replSet cfg0 --dbpath test-data/cfg0 --logpath test-data/config1 --bind_ip localhost --port 27019 & sleep 5
mongo --port 27019 --eval 'rs.initiate({_id:"cfg0", configsvr: true, members: [{"_id":1, "host":"localhost:27019"}]})'

# create replica set for shard 0
mkdir -p test-data/db01
mkdir -p test-data/db02
mkdir -p test-data/db03
mongod --shardsvr --replSet rs0 --dbpath test-data/db01 --logpath test-data/shard0 --bind_ip localhost --port 27018 & sleep 3
mongod --shardsvr --replSet rs0 --dbpath test-data/db02 --logpath test-data/shard1 --bind_ip localhost --port 27028 & sleep 3
mongod --shardsvr --replSet rs0 --dbpath test-data/db03 --logpath test-data/shard2 --bind_ip localhost --port 27038 & sleep 3
mongo --port 27018 --eval 'rs.initiate({_id : "rs0", members: [{ _id : 0, host : "localhost:27018" }, { _id : 1, host : "localhost:27028" }, { _id : 2, host : "localhost:27038" }]})'

# create replica set for shard 1
mkdir -p test-data/db11
mkdir -p test-data/db12
mkdir -p test-data/db13
mongod --shardsvr --replSet rs1 --dbpath test-data/db11 --logpath test-data/shard3 --bind_ip localhost --port 27118 & sleep 3
mongod --shardsvr --replSet rs1 --dbpath test-data/db12 --logpath test-data/shard4 --bind_ip localhost --port 27128 & sleep 3
mongod --shardsvr --replSet rs1 --dbpath test-data/db13 --logpath test-data/shard5 --bind_ip localhost --port 27138 & sleep 3
mongo --port 27118 --eval 'rs.initiate({_id : "rs1", members: [{ _id : 0, host : "localhost:27118" }, { _id : 1, host : "localhost:27128" }, { _id : 2, host : "localhost:27138" }]})'

# start mongos on default port 27017
mongos --configdb cfg0/localhost:27019 --logpath test-data/mongos1 --bind_ip localhost --port 26007 & sleep 3
# add shard
mongo --host localhost --port 26007 --eval 'sh.addShard("rs0/localhost:27018")'
mongo --host localhost --port 26007 --eval 'sh.addShard("rs1/localhost:27118")'

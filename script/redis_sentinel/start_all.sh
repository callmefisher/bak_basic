#!/bin/bash
./start_master_slave.sh
./start_sentinel.sh
sleep 2
ps -aef | grep redis

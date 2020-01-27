#!/bin/bash

./bin/server &
SERVER_PID=$!

time ./bin/client -messages=200 -size=5 -server 127.0.0.1 -port 8081

kill -9 $SERVER_PID
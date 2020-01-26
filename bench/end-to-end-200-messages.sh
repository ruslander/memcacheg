#!/bin/bash

./bin/server &
SERVER_PID=$!

time ./bin/client -messages=200 -size=5

kill -9 $SERVER_PID
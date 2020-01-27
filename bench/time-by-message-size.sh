#!/bin/bash

for i in 10 100 300 500 900 1200 1500 1800 2000;
do
  sizeOfPayload=$i

  echo ""
  echo "Workload message size $i (bytes)"

  ./bin/server > /dev/null 2>&1 &
  SERVER_PID=$!

  time ./bin/client -messages=200 -size=$sizeOfPayload -server 127.0.0.1 -port 8081

  kill -9 $SERVER_PID

done;


#!/bin/bash

minArg=$2
maxArg=$4
step=64

echo "Workloads from $minArg to $maxArg"

for i in 10 100 1000 10000 100000;
do
  sizeOfPayload=$((1000*1024*1024/$i))

  echo ""
  echo "Workload $i $sizeOfPayload"

  ./bin/server > /dev/null 2>&1 &
  SERVER_PID=$!

  time ./bin/client -messages=200 -size=$sizeOfPayload > /dev/null 2>&1

  kill -9 $SERVER_PID

done;


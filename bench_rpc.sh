#!/usr/bin/env bash

if [ -d "kitex-benchmark" ]; then
  echo "already download sonic"
else
  # add di
  git clone http://github.com/zdyj3170101136/kitex-benchmark
fi

cd scripts

./benchmark_grpc.sh

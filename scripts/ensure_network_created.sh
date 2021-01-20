#!/bin/bash

SUBSCRIPT=$1
# echo "create network $SUBSCRIPT"

# check if network is created or not
output=$(docker inspect $SUBSCRIPT 2>&1)
ret=$?

if [[ $ret -ne 0 ]]; then
  # this means the network doesn't exist yet, we will create one:
  echo Creating network $SUBSCRIPT
  docker network create -d bridge $SUBSCRIPT >/dev/null
fi

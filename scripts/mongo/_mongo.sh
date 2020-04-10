#!/bin/bash

SUBSCRIPT=$1

cd ./scripts/mongo

case $SUBSCRIPT in
  start)
    ./start.sh
    ;;
  stop)
    ./stop.sh
    ;;
  uri)
    echo "mongodb://"$LOCAL_MONGODB_CONTAINER_NAME":27017"
    ;;
  *)
    echo \> Wrong usage: mongo subscript should be one of \"start\", \"stop\".
    echo \> Example: ./script.sh mongo start
esac


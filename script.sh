#!/bin/bash

# This script is to run scripts specified in ./scripts folder

SUBSCRIPT=$1

argsWithoutSubscript=
for ARG in $*
  do
    if [ $ARG != $SUBSCRIPT ]
      then
        argsWithoutSubscript+=" $ARG"
    fi
done

export LOCAL_NETWORK=local-network
export LOCAL_MONGODB_CONTAINER_NAME=local-mongo

./scripts/ensure_network_created.sh $LOCAL_NETWORK

# This is to make all echo from now on to be printed in green
printf "$(tput setaf 2)"

case $SUBSCRIPT in
  dev)
    ./scripts/dev.sh
    ;;
  build)
    ./scripts/build.sh $argsWithoutSubscript
    ;;
  mongo)
    ./scripts/mongo/_mongo.sh $argsWithoutSubscript
    ;;
  clean)
    rm -rf build/fresh/tmp
    rm -f main
    docker system prune -f
    echo Cleaned
    ;;
  *)
    echo \> Wrong usage: currently support [\"dev\", \"build\", \"mongo\"]
esac

echo

# End of green echo
printf "$(tput sgr0)"

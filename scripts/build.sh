#!/bin/bash

# This script is to be run from root of project

SUBSCRIPT=$1

case $SUBSCRIPT in
  main)
    go build -o main .
    echo Done
    ;;
  image)
    docker build -t go-rest-template -f build/docker/Dockerfile .
    ;;
  *)
    echo \> Wrong usage: build subscript should be one of \"main\" or \"image\" \(docker image\).
    echo \> Example: ./script.sh build image
esac


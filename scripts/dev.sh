#!/bin/bash

# This script is to be run from root of project

docker build -t go-rest-template:dev -f build/docker/Dockerfile.dev .
docker run -it --rm -p 5000:5000 -v $(pwd):/go-rest-template --network $LOCAL_NETWORK go-rest-template:dev

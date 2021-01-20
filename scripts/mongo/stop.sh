#!/bin/bash

# found container, including active one only
FOUND_RUNNING_CONTAINER_ID=$(docker ps -qf "name=^${LOCAL_MONGODB_CONTAINER_NAME}$")
# found container id, including inactive one
FOUND_CONTAINER_ID=$(docker ps -aqf "name=^${LOCAL_MONGODB_CONTAINER_NAME}$")

if [[ -n ${FOUND_RUNNING_CONTAINER_ID} ]]; then
  echo "> Stopping local MongoDB ..."
  docker stop ${LOCAL_MONGODB_CONTAINER_NAME} >/dev/null
  echo "> Local MongoDB has been stopped"
elif [[ -n ${FOUND_CONTAINER_ID} ]]; then
  echo "> Local MongoDB is stopped already"
else 
  echo "> Local MongoDB has not been started yet"
fi
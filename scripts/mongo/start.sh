#!/bin/bash

# found container, including active one only
FOUND_RUNNING_CONTAINER_ID=$(docker ps -qf "name=^${LOCAL_MONGODB_CONTAINER_NAME}$")
# found container id, including inactive one
FOUND_CONTAINER_ID=$(docker ps -aqf "name=^${LOCAL_MONGODB_CONTAINER_NAME}$")

if [[ -n ${FOUND_RUNNING_CONTAINER_ID} ]]; then
  echo "> Local MongoDB already started"
elif [[ -n ${FOUND_CONTAINER_ID} ]]; then
  echo "> Found a local MongoDB container, restarting ..."
  docker start ${LOCAL_MONGODB_CONTAINER_NAME} >/dev/null
  echo "> Restarted existing local MongoDB"
else
  echo "> No local MongoDB container found"
  echo "> Starting a new local MongoDB ..."
  docker run -d -v /var/lib/mongo --name ${LOCAL_MONGODB_CONTAINER_NAME} -p 27017:27017 --network $LOCAL_NETWORK mongo >/dev/null
  echo "> Local MongoDB has been started"
fi
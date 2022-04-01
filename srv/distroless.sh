#!/bin/bash 
PROJ_NAME=srv
PROJ_VERSION=0.0.1

docker build \
--build-arg PROJ_NAME=$PROJ_NAME \
-t $PROJ_NAME:$PROJ_VERSION .

docker run -it -d -v ~/:/source -p 100:80 --name $PROJ_NAME $PROJ_NAME:$PROJ_VERSION

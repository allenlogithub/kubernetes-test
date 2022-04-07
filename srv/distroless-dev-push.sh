#!/bin/bash
DOCKERHUB_ACCOUNT=outlier11111001000
PROJ_NAME=srv
PROJ_VERSION=0.0.1

docker build \
--build-arg PROJ_NAME=$PROJ_NAME \
-t $DOCKERHUB_ACCOUNT/$PROJ_NAME:$PROJ_VERSION .

docker push $DOCKERHUB_ACCOUNT/$PROJ_NAME
#!/bin/bash
DOCKERHUB_ACCOUNT=outlier11111001000
PROJ_NAME=srv
PROJ_VERSION=0.0.1

docker pull $DOCKERHUB_ACCOUNT/$PROJ_NAME:$PROJ_VERSION

docker run -it -d -v ~/:/source -p 100:80 --name $PROJ_NAME \
--env MYSQL_ADDRESS="172.17.0.2:3306" \
--env MYSQL_NET="tcp" \
--env MYSQL_USER="root" \
--env MYSQL_PASSWORD="my-secret-pw" \
--env MYSQL_DBNAME="dev" \
$DOCKERHUB_ACCOUNT/$PROJ_NAME:$PROJ_VERSION

#!/usr/bin/env bash

source setenv.sh

# Criar rede
echo "Criando a rede $DOCKER_NETWORK..."
docker network ls | grep $DOCKER_NETWORK
if [ "$?" != 0 ]; then
   docker network create $DOCKER_NETWORK
fi

# Mysql
echo "Subindo o mysql..."
docker run -d --name mysqldb --network $DOCKER_NETWORK  \
-p 3306:3306 \
-e MYSQL_USER=${MYSQL_USER} \
-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
-e MYSQL_DATABASE=${MYSQL_DATABASE} \
mysql:5.7

# Message API
echo "Subindo o go-auth..."
docker run -d --name go-auth --network $DOCKER_NETWORK \
-p 8181:8080 \
-e MYSQL_USER=${MYSQL_USER} \
-e MYSQL_PASSWORD=${MYSQL_PASSWORD} \
-e MYSQL_HOSTNAME=${MYSQL_HOSTNAME} \
-e MYSQL_DATABASE=${MYSQL_DATABASE} \
-e MYSQL_PORT=${MYSQL_PORT} \
-e TZ=America/Sao_Paulo \
marceloagmelo/go-auth

# Listando os containers
docker ps

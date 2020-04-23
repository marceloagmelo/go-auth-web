#!/usr/bin/env bash

source setenv.sh

# Mysql
echo "Finalizando o mysql..."
docker rm -f $MYSQL_HOSTNAME

# API
echo "Finalizando o ${APP_API_NAME}..."
docker rm -f ${APP_API_NAME}

# Aplicação
echo "Finalizando o ${APP_NAME}..."
docker rm -f ${APP_NAME}

# Remover rede
echo "Removendo a rede message-net..."
docker network rm $DOCKER_NETWORK

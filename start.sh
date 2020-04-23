#!/usr/bin/env bash

source setenv.sh

# Verificar rede
echo "Verificando se existe a rede $DOCKER_NETWORK..."
docker network ls | grep $DOCKER_NETWORK
if [ "$?" != 0 ]; then
   echo "Rede $DOCKER_NETWORK não existe!"
   exit 0
fi

echo "Subindo o ${APP_NAME}..."
docker run -d --name ${APP_NAME} --network $DOCKER_NETWORK  \
-p 7070:8080 \
-e API_SERVICE_URL=${API_SERVICE_URL} \
-e TZ=America/Sao_Paulo \
${DOCKER_REGISTRY}/${APP_NAME}

# Listando os containers
docker ps
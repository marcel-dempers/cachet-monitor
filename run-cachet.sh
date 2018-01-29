#!/bin/bash

export MSYS_NO_PATHCONV=1

#echo "starting postgres..."
#docker run -d --rm --name postgres -e POSTGRES_USER=postgres -e PGDATA=/var/lib/postgresql/data/pgdata -e POSTGRES_PASSWORD=postgres -v $DIR/data:/var/lib/postgresql/data -d postgres:9.5

echo 'PostGreSQL Host:'
read db_host

echo 'PostGreSQL Username:'
read username

echo 'PostGreSQL Password:'
read -s password

docker rm --force cachet

echo "starting cachet..."
docker run -it --rm --name cachet \
    -p 8000:8000 \
    -e DB_DRIVER=pgsql \
    -e DB_HOST=$db_host \
    -e DB_DATABASE="postgres" \
    -e PGSSLMODE="require" \
    -e DB_USERNAME=$username \
    -e DB_PASSWORD=$password \
    -e APP_KEY=base64:p6+P8N3ZB2BJKoucim0bis48jB6dPqhnitpt5bihN4A= \
    cachethq/docker:2.3.13
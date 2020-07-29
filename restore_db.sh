#!/bin/sh
# Setting up env variables:
set -a
. config/.env
set +a

MYSQL_CONTAINER_ID=$(docker ps -aqf "name=bible_db")
MYSQL_DUMP_PATH="/var/lib/mysql"

echo "Copying backup files to the container..."
docker cp backup/aa.sql ${MYSQL_CONTAINER_ID}:${MYSQL_DUMP_PATH}/aa.sql
docker cp backup/nvi.sql ${MYSQL_CONTAINER_ID}:${MYSQL_DUMP_PATH}/nvi.sql
docker cp backup/acf.sql ${MYSQL_CONTAINER_ID}:${MYSQL_DUMP_PATH}/acf.sql

echo "Starting backup restore process..."
docker exec -it ${MYSQL_CONTAINER_ID} bash -c "mysql --max_allowed_packet=1048576000 --user=root --password=${MYSQL_PASSWORD} --database=${MYSQL_DATABASE} < ${MYSQL_DUMP_PATH}/aa.sql"
docker exec -it ${MYSQL_CONTAINER_ID} bash -c "mysql --max_allowed_packet=1048576000 --user=root --password=${MYSQL_PASSWORD} --database=${MYSQL_DATABASE} < ${MYSQL_DUMP_PATH}/nvi.sql"
docker exec -it ${MYSQL_CONTAINER_ID} bash -c "mysql --max_allowed_packet=1048576000 --user=root --password=${MYSQL_PASSWORD} --database=${MYSQL_DATABASE} < ${MYSQL_DUMP_PATH}/acf.sql"

echo "Removing backup files from the container..."
docker exec -it ${MYSQL_CONTAINER_ID} sh -c "rm ${MYSQL_DUMP_PATH}/aa.sql"
docker exec -it ${MYSQL_CONTAINER_ID} sh -c "rm ${MYSQL_DUMP_PATH}/nvi.sql"
docker exec -it ${MYSQL_CONTAINER_ID} sh -c "rm ${MYSQL_DUMP_PATH}/acf.sql"

echo "Done."
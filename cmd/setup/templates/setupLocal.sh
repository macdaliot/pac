#!/bin/sh

PROJECT=[psi[.projectName]]
SERVICE=sample
docker network create pac-$PROJECT
docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --name elastic --network pac-$PROJECT elasticsearch:7.1.0 
docker run -d -p 27017:27017 --name mongo --network pac-$PROJECT mongo:3.6
docker run -d -p 5601:5601 -e "ELASTICSEARCH_HOSTS=http://elastic:9200" --name kibana --network pac-$PROJECT kibana:7.1.0

cat << EOF > services/$SERVICE/.env
ENVIRONMENT=local
# Log level values 'fatal', 'error', 'warn', 'info', 'debug', 'trace' or 'silent'.
LOG_LEVEL=info
MONGO_CONN_STRING=mongodb://mongo:27017/$PROJECT
PROJECTNAME=$PROJECT
ES_ENDPOINT=http://elastic:9200
EOF

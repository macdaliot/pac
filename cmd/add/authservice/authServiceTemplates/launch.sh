#! /bin/bash

# Ensure AWS credentials were provided
AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)
if [ ! -z "$AWS_ACCESS_KEY_ID" ] || [ ! -z "$AWS_SECRET_ACCESS_KEY" ]; then
  # Delete service container, if running
  SERVICE_CONTAINER_PID=$(docker ps -aqf Name=pac-auth-service)
  if [ ! -z "$SERVICE_CONTAINER_PID" ]; then
    docker stop $SERVICE_CONTAINER_PID >> /dev/null
    docker rm $SERVICE_CONTAINER_PID >> /dev/null
    echo "Stopped existing service container"
  fi

  # NPM install
  if [ ! -d node_modules/ ]; then
    npm i
  fi

  # Compile, build, and run
  npx tsc
  docker build -t pac-auth-service .
  CONTAINER=`docker run --name pac-auth-service -p 3000:3000  -d pac-auth-service`
  docker logs -f $CONTAINER
else
  echo "ABORTED LAUNCH: The AWS keys were not found. Try running 'aws configure'"
fi

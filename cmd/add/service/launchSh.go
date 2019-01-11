package service

import (
  "runtime"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

// CreateLaunchSh creates a templated shell script to launch the service
func CreateLaunchSh(filePath string, config map[string]string) {
  const template = `#! /bin/bash

# Ensure AWS credentials were provided
AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)
if [ ! -z "$AWS_ACCESS_KEY_ID" ] || [ ! -z "$AWS_SECRET_ACCESS_KEY" ]; then
  # Start DynamoDB if not running
  if docker port pac-db-local >>/dev/null 2>/dev/null ; then
    echo "Local DynamoDB already exists, skipping DynamoDB creation"
  else
    docker run --name pac-db-local -d -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb
  fi

  # Delete service container, if running
  SERVICE_CONTAINER_PID=$(docker ps -aqf Name=pac-.*-service)
  if [ ! -z "$SERVICE_CONTAINER_PID" ]; then
    docker stop $SERVICE_CONTAINER_PID >> /dev/null
    docker rm $SERVICE_CONTAINER_PID >> /dev/null
    echo "Stopped existing service container"
  fi

  # NPM install
  if [ ! -d node_modules/ ]; then
    npm i
  fi

  # Attempt creation of DynamoDB table
  aws dynamodb create-table --cli-input-json file://dynamoConfig.json --endpoint-url http://localhost:8000 >> /dev/null 2> /dev/null

  # Compile, build, and run
  npx tsc
  docker build -t pac-{{.serviceName}}-service .
  docker run --name pac-{{.serviceName}}-service -p 3000:3000 --link pac-db-local -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -d pac-{{.serviceName}}-service
  echo "DONE! Launched microservice locally"
else
  echo "ABORTED LAUNCH: The AWS keys were not found. Try running 'aws configure'"
fi
`
  files.CreateFromTemplate(filePath, template, config)
  if runtime.GOOS == "windows" {
    files.ChangePermissions(str.Concat(".\\", config["serviceName"], "\\launch.sh"), 0755)
  } else {
    files.ChangePermissions(str.Concat("./", config["serviceName"], "/launch.sh"), 0755)
  }
}

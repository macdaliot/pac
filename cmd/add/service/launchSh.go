package service

import (
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

// CreateLaunchSh creates a templated shell script to launch the service
func CreateLaunchSh(filePath string, config map[string]string) {
  const template = `#! /bin/bash

AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)
if [ ! -z "$AWS_ACCESS_KEY_ID" ] || [ ! -z "$AWS_SECRET_ACCESS_KEY" ]; then
  if [ ! -d node_modules/ ]; then
    npm i
  fi
  docker stop pac-{{.serviceName}}-service || true
  docker rm pac-{{.serviceName}}-service || true
  npx tsc server.ts
  docker build -t pac-{{.serviceName}}-service .
  docker run --name pac-{{.serviceName}}-service -p 3000:3000 --link pac-db-local -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -d pac-{{.serviceName}}-service
  echo "DONE! Launched microservice locally"
else
  echo "ABORTED LAUNCH: The AWS keys were not found. Try running 'aws configure'"
fi
`
  files.CreateFromTemplate(filePath, template, config)
  files.ChangePermissions(str.Concat("./", config["serviceName"], "/launch.sh"), 0755)
}

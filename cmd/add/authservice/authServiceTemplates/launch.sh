#! /bin/bash

# Ensure AWS credentials were provided
if [[ $(uname) =~ MINGW.* ]]; then
  AWS_ACCESS_KEY_ID=$(aws.cmd --profile default configure get aws_access_key_id)
  AWS_SECRET_ACCESS_KEY=$(aws.cmd --profile default configure get aws_secret_access_key)
else
  AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
  AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)
fi
if [ ! -z "$AWS_ACCESS_KEY_ID" ] || [ ! -z "$AWS_SECRET_ACCESS_KEY" ]; then
  # Get project name
  PROJECT_NAME=$(echo $(sed -e 's/\(.*\)\/.*/\1/' -e 's/\(.*\)\/.*/\1/' -e 's/.*\///g' <<< $(pwd)))

  # Create Docker network if it does not exist
  if docker network inspect pac-$PROJECT_NAME >>/dev/null 2>/dev/null; then
    echo "Docker network already exists, skipping network creation"
  else
    docker network create pac-$PROJECT_NAME >>/dev/null
    if [ $(echo $?) -eq 0 ]; then
      echo "Created Docker network pac-$PROJECT_NAME"
    else
      echo "ERROR: Something went wrong creating your Docker network. Are you in the necessary Docker groups?"
      exit 2
    fi
  fi

  # Delete current service's container, if it exists
  SERVICE_CONTAINER_PID=$(docker ps -aqf Name=pac-$PROJECT_NAME-{{.serviceName}})
  if [ ! -z "$SERVICE_CONTAINER_PID" ]; then
    docker stop $SERVICE_CONTAINER_PID >> /dev/null
    docker rm $SERVICE_CONTAINER_PID >> /dev/null
    echo "Stopped existing service container"
  fi

  # Ensure all microservices are launched
  SERVICE_DIR=$(echo $(sed -e 's/\(.*\)\/.*/\1/' <<< $(pwd)))
  for DIR in ../*; do
    if [ ! $DIR == "../haproxy.cfg" ]; then
      SERVICE_NAME=$(echo $(sed -e 's/..\///g' <<< $DIR))
      if docker port pac-$PROJECT_NAME-$SERVICE_NAME >>/dev/null 2>/dev/null; then
        echo "$SERVICE_NAME microservice already launched locally"
      else
        pushd $DIR >>/dev/null
        # NPM install if no node_modules folder
        if [ ! -d node_modules/ ]; then
          echo "Installing NPM modules for the $SERVICE_NAME microservice..."
          npm i
          echo "Finished installing NPM modules for the $SERVICE_NAME microservice"
        fi

        # Compile, build, and launch
        npx tsc
        echo "Started building $SERVICE_NAME Docker image"
        docker build -t pac-$PROJECT_NAME-$SERVICE_NAME .
        docker run --name pac-$PROJECT_NAME-$SERVICE_NAME --network pac-$PROJECT_NAME -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -d pac-$PROJECT_NAME-$SERVICE_NAME
        if [ $(echo $?) -eq 0 ]; then
          echo "The $SERVICE_NAME microservice was successfully launched locally"
        else
          echo "ERROR: Something went wrong creating the $SERVICE_NAME microservice Docker container"
          exit 2
        fi
        popd >>/dev/null
      fi
    fi
  done

  # Start HaProxy if not running, reload config if running
  HAPROXY_PORT=3000
  if docker port pac-proxy-local >>/dev/null 2>/dev/null; then
    echo "Local microservice proxy already exists, skipping proxy creation"
    docker kill -s HUP pac-proxy-local >>/dev/null
    if [ $(echo $?) -eq 0 ]; then
      echo "Reloaded local microservice proxy configuration"
    else
      echo "ERROR: Something went wrong reloading the microservice proxy configuration"
      exit 2
    fi
  else
    HAPROXY_CONFIG_PATH=$(echo $(pwd) | sed -e 's/\(.*\)\/.*/\1/')
    HAPROXY_CONFIG_PATH=$(sed -e 's/^\/c/C:/g' <<< $HAPROXY_CONFIG_PATH)
    docker run --name pac-proxy-local --network pac-$PROJECT_NAME -p $HAPROXY_PORT:$HAPROXY_PORT -v $HAPROXY_CONFIG_PATH:/usr/local/etc/haproxy:ro -d haproxy >>/dev/null
    if [ $(echo $?) -eq 0 ]; then
      echo "Launched local microservice proxy running on port $HAPROXY_PORT"
    else
      echo "ERROR: Something went wrong launching the local microservice proxy Docker container. Is port $HAPROXY_PORT open?"
      exit 2
    fi
  fi
  echo "All microservices are now available at port $HAPROXY_PORT:"
  echo "    i.e. http://localhost:$HAPROXY_PORT/api/$SERVICE_NAME"
else
  echo "ABORTED LAUNCH: The AWS keys were not found. Try running 'aws configure'"
fi
#!/bin/bash

JOB_NAMES=(
  "entrypoint"
  "front-end"
  "release-through-staging"
  "release-to-production"
  "services"
)

for JOB_NAME in "${JOB_NAMES[@]}"; do
  java -jar jenkins-cli.jar -s https://jenkins.{{.projectName}}.pac.pyramidchallenges.com -auth pyramid:systems create-job $JOB_NAME < $JOB_NAME.xml
done

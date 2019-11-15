#!/bin/bash

# $1 - The environment name to deploy to

# If the bucket does not exist, wait 2.5 minutes for Terraform to create it
# This command should only end in a sleep on the first run of a new environment
aws s3 ls | grep "$1.[psi[.projectName]].pac.pyramidchallenges.com" >> /dev/null && echo "Bucket exists" || sleep 150

# Sync the files in the newly build `dist` folder with the environment's S3 bucket
aws s3 sync dist s3://$1.[psi[.projectName]].pac.pyramidchallenges.com --acl public-read

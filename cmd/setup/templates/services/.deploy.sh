#! /bin/bash

# $1 - The environment name to deploy to

## Download Terraform
TF_VERSION="0.11.14"

## Navigate to Terraform directory
cd ../terraform/$1

## Initialize Terraform
terraform init -input=false

## Save Terraform plan
terraform plan -out=tfplan -input=false

## Apply Terraform plan
terraform apply -input=false tfplan

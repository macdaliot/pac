#! /bin/bash

## Download Terraform
TF_VERSION="0.11.13"


## Navigate to Terraform directory
cd ./svc/terraform

## Initialize terraform
terraform init -input=false

## Save tf plan
terraform plan -out=tfplan -input=false

## Apply tf plan
terraform apply -input=false tfplan

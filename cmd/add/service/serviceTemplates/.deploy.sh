#! /bin/bash

## Download Terraform
TF_VERSION="0.11.13"

TIMESTAMP=`date +%s`

## Navigate to Terraform directory
cd ../terraform

## Initialize terraform
terraform init -input=false

## Save tf plan
terraform plan -out=tfplan_$TIMESTAMP -input=false

## Apply tf plan
terraform apply -input=false tfplan_$TIMESTAMP

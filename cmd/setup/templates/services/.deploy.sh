#! /bin/bash

## Download Terraform
TF_VERSION="0.11.13"

if [ $(find . -maxdepth 1 -type d | wc -l) -gt 2 ]; then
  ## Navigate to Terraform directory
  cd terraform

  ## Initialize terraform
  terraform init -input=false

  ## Save tf plan
  terraform plan -out=tfplan -input=false

  ## Apply tf plan
  terraform apply -input=false tfplan
fi

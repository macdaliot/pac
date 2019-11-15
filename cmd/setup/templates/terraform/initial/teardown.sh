#! /bin/bash

rm terraform.tfstate terraform.tfstate.backup
terragrunt init
sed -i 's/backend "s3"/backend "local"/g' providers.tf
terraform init -force-copy
terraform destroy -auto-approve

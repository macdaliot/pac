#!/bin/bash

# Create AWS resources using the local backend configuration
terraform init
terraform apply -input=false -auto-approve

# Switch to S3 Terraform backend configuration (using Terragrunt, see `terragrunt.hcl` for more information)
sed -i 's/backend "local"/backend "s3"/g' providers.tf

# Migrate Terraform state to S3 backend where it can be picked up by the project
terragrunt init -force-copy
terragrunt apply -input=false -auto-approve

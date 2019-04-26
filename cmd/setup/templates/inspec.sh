#!/bin/bash

## Download Terraform
TF_VERSION="0.11.13"

# download and install InSpec
#wget https://packages.chef.io/files/stable/inspec/3.9.3/el/7/inspec-3.9.3-1.el7.x86_64.rpm \
#rpm -i inspec-3.9.3-1.el7.x86_64.rpm

#install inspec-iggy plugin
#inspec plugin install inspec-iggy

## Navigate to Terraform directory
cd terraform

# remove any previous artifcats of terraform and inspec
rm -rf .terraform
rm -f pac_state

## Initialize terraform
terraform init -input=false

# download Terraform state file
terraform state pull > pac_state


#generate InSpec profile and generate tests based on Terraform files, create "pac" profile
inspec terraform generate  --overwrite --tfstate pac_state --name pac

# exec InSpec
# output in json format
# target aws platform
# use "pac" profile
# redirect output to file
inspec exec --reporter=junit -t aws:// pac > inspec_terraform_report.xml

# do something to get results into Sonarqube(?)
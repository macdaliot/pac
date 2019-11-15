#
# http://www.terraform.io/docs/backends/config.html
#
# ADD NOTES HERE
#

terraform {
  backend "local" {}
  required_version = "0.12.7"
}

#
# http://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so Terraform can talk to AWS
#
# Note: `region` is not listed as required in the documentation, but Terraform
#   will prompt us for a region if it is not set here
#
provider "aws" {
  region  = "us-east-1"
  version = "~>2.21"
}

provider "template" {
  version = "2.1"
}

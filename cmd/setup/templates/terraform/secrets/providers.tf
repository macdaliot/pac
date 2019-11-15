#
# http://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.[psi[.projectName]].[psi[.hostedZone]]"
    key    = "tfstate/dev/secrets"
    region = "[psi[.region]]"
  }
}

#
# http://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
# `region` is not listed as required in the documentation, but we are prompted
#   for a value if a `region` is not specified here
#
provider "aws" {
  region = "[psi[.region]]"
  version = "~>2.21"
}

data "terraform_remote_state" "bootstrap" {
  backend = "s3"
  config = {
    bucket = "terraform.[psi[.projectName]].[psi[.hostedZone]]"
    key    = "bootstrap/terraform.tfstate"
    region = "[psi[.region]]"
  }
}

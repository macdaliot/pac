#
# http://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.[psi[.projectName]].pac.pyramidchallenges.com"
    key    = "tfstate/[psi[.env]]/ssl"
    region = "[[ .region ]]"
  }

  required_version = "0.12.7"
}

#
# http://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "[[ .region ]]"
  version = "~>2.21"
}

provider "template" {
  version = "~>2.1"
}

#
# http://www.terraform.io/docs/providers/random/index.html
#
provider "random" {
  version = "~>2.1"
}

data "terraform_remote_state" "dns" {
  backend = "s3"
  config = {
    bucket = "terraform.[psi[.projectName]].pac.pyramidchallenges.com"
    key    = "tfstate/[psi[.env]]/dns"
    region = "[[ .region ]]"
  }
}

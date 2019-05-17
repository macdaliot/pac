#
# http://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.yoshi.pac.pyramidchallenges.com"
    key    = "state/development"
    region = "us-east-2"
  }
}

#
# http://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "us-east-2"

  version = "1.60"
}

provider "template" {
  version = "2.1"
}

#
# http://www.terraform.io/docs/providers/random/index.html
#
provider "random" {
  version = "2.1"
}

// Reference DNS infor in different S3 key
data "terraform_remote_state" "dns" {
  backend = "s3"

  config {
    bucket = "terraform.yoshi.pac.pyramidchallenges.com"
    key    = "state/development/dns"
    region = "us-east-2"
  }
}
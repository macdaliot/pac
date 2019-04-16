#
# https://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.{{ .projectName }}.pac.pyramidchallenges.com"
    key    = "state/development"
    region = "{{ .region }}"
  }
}

#
# https://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "{{ .region }}"

  version = "{{ .terraformAWSVersion }}"
}

provider "template" {
  version = "{{ .terraformTemplateVersion }}"
}
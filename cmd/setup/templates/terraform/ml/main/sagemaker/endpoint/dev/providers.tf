terraform {
  backend "s3" {
    bucket = "terraform.{{ .projectName }}.{{ hostedZone }}"
    key    = "tfstate/dev/sagemaker/development"
    region = "{{ .region }}"
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
  region = "{{ .region }}"

  version = "~>2.21"
}
#
# https://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.{{ .projectName }}.pac.pyramidchallenges.com"
    key    = "state/development"
    region = "{{ .Region }}"
    dynamodb_table = "pac-{{ .projectName }}-terraform-locking-table"
  }
}

#
# https://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "{{ .Region }}"

  version = "{{ .AWSVersion }}"
}

provider "template" {
  version = "{{ .TemplateVersion }}"
}

resource "aws_dynamodb_table" "{{ .projectName }}_dynamodb_table" {
  name           = "pac-{{ .projectName }}-terraform-locking-table"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "LockID"
  stream_view_type = "NEW_AND_OLD_IMAGES"
  stream_enabled = true

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name        = "pac-{{ .projectName }}-terraform-locking-table"
  }
}
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "us-east-2"
}

resource "aws_vpc" "vpc" {
  cidr_block = "${var.cidr_block}"

  tags = {
    "Name" = "${var.project_name}-vpc"
  }
}
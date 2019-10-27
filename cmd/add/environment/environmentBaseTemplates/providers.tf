#
# http://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.{{ .projectFqdn }}"
    key    = "tfstate/{{ .env }}/{{ .environmentName }}"
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
  region  = "{{ .region }}"
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

data "terraform_remote_state" "bootstrap" {
  backend = "s3"
  config = {
    bucket = "terraform.{{.projectFqdn}}"
    key    = "bootstrap/terraform.tfstate"
    region = "{{.region}}"
  }
}

data "terraform_remote_state" "dns" {
  backend = "s3"
  config = {
    bucket = "terraform.{{.projectFqdn}}"
    key    = "tfstate/dev/dns"
    region = "{{.region}}"
  }
}

data "terraform_remote_state" "ssl" {
  backend = "s3"
  config = {
    bucket = "terraform.{{.projectFqdn}}"
    key    = "tfstate/dev/ssl"
    region = "{{.region}}"
  }
}

data "terraform_remote_state" "management" {
  backend = "s3"
  config = {
    bucket = "terraform.{{.projectFqdn}}"
    key    = "tfstate/dev/management_vpc"
    region = "{{.region}}"
  }
}

data "terraform_remote_state" "secrets" {
  backend = "s3"
  config = {
    bucket = "terraform.{{.projectFqdn}}"
    key    = "tfstate/dev/secrets"
    region = "{{.region}}"
  }
}

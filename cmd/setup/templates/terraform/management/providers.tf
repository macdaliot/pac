#
# http://www.terraform.io/docs/backends/config.html
#
# S3 bucket to store infrastructure state
#
terraform {
  backend "s3" {
    bucket = "terraform.{{ .projectName }}.{{ .hostedZone }}"
    key    = "tfstate/{{ .env }}/management_vpc"
    region = "{{ .region }}"
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
  region = "{{ .region }}"
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

provider "acme" {
  server_url = "https://acme-v02.api.letsencrypt.org/directory"
}

data "terraform_remote_state" "bootstrap" {
  backend = "s3"
  config = {
    bucket = "terraform.[[.projectName]].{{.hostedZone}}"
    key    = "bootstrap/terraform.tfstate"
    region = "{{.region}}"
  }
}

data "terraform_remote_state" "dns" {
  backend = "s3"
  config = {
    bucket = "terraform.{{ .projectName }}.{{ .hostedZone }}"
    key    = "tfstate/{{ .env }}/dns"
    region = "{{ .region }}"
  }
}

data "terraform_remote_state" "ssl" {
  backend = "s3"
  config = {
    bucket = "terraform.[[.projectName]].{{.hostedZone}}"
    key    = "tfstate/dev/ssl"
    region = "{{.region}}"
  }
}

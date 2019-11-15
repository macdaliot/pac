locals {
  hosted_zone  = "pac.pyramidchallenges.com"
  project_name = "bdsopracthree"
}

remote_state {
  backend = "s3"
  config = {
    bucket     = "terraform.${local.project_name}.${local.hosted_zone}"
    encrypt    = true
    key        = "${path_relative_to_include()}/bootstrap/terraform.tfstate"
    kms_key_id = "alias/pac/${local.project_name}"
    region     = "us-east-1"
  }
}

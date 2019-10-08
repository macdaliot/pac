#----------------------------------------------------------------------------------------------------------------------
# BACKEND
#----------------------------------------------------------------------------------------------------------------------

# Note: Terraform does not allow variables in the terraform block, so change the {{.projectName}} and {{.region}}
# to their appropriate values.

terraform {
  backend "local" {}
}

#
# http://www.terraform.io/docs/backends/config.html
#
# To store the state in S3.
#
# 1. Comment out the first terraform block, then uncomment the second terraform block.
# 2. Run 'terrafform init'. 
#    Terraform will notice the backend configuration change and ask you if you want to push state to S3.
# 3. Say "yes"
# 4. Done. Terraform is now storing the state in S3.
#
# Note: These resources will need be destroyed after teardown of the management VPC.
#
# terraform {
#   backend "s3" {
#     bucket         = "terraform.{{.projectName}}.{{.hostedZone}}"
#     dynamodb_table = "{{.projectName}}_tf_lock_table"
#     encrypt        = true
#     key            = "tfstate/bootstrap"
#     kms_key_id     = "alias/pac/{{.projectName}}"
#     region         = "{{.region}}"
#   }

#    required_version = "0.12.7"
# }

#----------------------------------------------------------------------------------------------------------------------
# SSH ENCRYPTION & AWS KEY PAIR
#----------------------------------------------------------------------------------------------------------------------
module "kms_key" {
    source = "./modules/kms"

    project_name = var.project_name
}

# Create tls key pair
module "management_tls" {
  source = "./modules/tls"
}

output "management_tls" {
  value = module.management_tls.tls
}

# Create AWS key pair
module "keypair" {
  source = "./modules/keypair"

  project_name       = var.project_name
  public_key_openssh = module.management_tls.tls.public_key_openssh
}

output "default_keypair" {
  value = module.keypair.default.key_name
}

#----------------------------------------------------------------------------------------------------------------------
# TERRAFORM STATE S3 BUCKET
#----------------------------------------------------------------------------------------------------------------------

module "terraform_bucket" {
  source = "./modules/s3"

  project_name = "terraform.${var.project_name}.${var.hosted_zone}"
  key_id       = module.kms_key.kms_key.id
  region       = var.region
}

# STATE LOCKING TABLE
module "terraform_lock_table" {
  source = "./modules/dynamodb"

  attribute_name = "LockID"
  project_name   = var.project_name
  table_name     = "${var.project_name}_tf_lock_table"
}

#----------------------------------------------------------------------------------------------------------------------
# EC2 INSTANCE
#----------------------------------------------------------------------------------------------------------------------

# TODO: finish the EC2 module
# module "bootstrap_instance" {
#   source = "./modules/ec2"

#   ami                  = data.aws_ami.amazon-linux-2.id
#   availability_zone    = var.availability_zone
#   cpu_core_count       = var.cpu_core_count
#   cpu_threads_per_core = var.cpu_threads_per_core
#   host_id              = var.host_id
#   placement_group      = var.placement_group
#   tenancy              = var.tenancy
# }

#----------------------------------------------------------------------------------------------------------------------
# AMI
#----------------------------------------------------------------------------------------------------------------------

# this should get the latest Amazon Linux 2 ami id for the current region
data "aws_ami" "amazon-linux-2" {
  most_recent = true

  filter {
    name   = "owner-alias"
    values = ["amazon"]
  }

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }

  owners = ["amazon"]
}

#----------------------------------------------------------------------------------------------------------------------
# PROVIDER
#----------------------------------------------------------------------------------------------------------------------
#
# http://www.terraform.io/docs/providers/aws/index.html
#
# AWS provider so terraform can talk to AWS
#
provider "aws" {
  # not listed as require in documentation but will be asked for it if not set
  region = "{{.region}}"

  version = "~>2.21"
}
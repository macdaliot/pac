locals {
  worker_instance_name = "${var.project_name}-worker"
}

module "kms_key" {
  source     = "./modules/kms"
  alias_name = "alias/pac/${var.project_name}"
}

module "terraform_state_bucket" {
  source      = "./modules/s3"
  bucket_name = "terraform.${var.project_name}.${var.hosted_zone}"
  key_id      = module.kms_key.key.id
  region      = var.region
}

module "worker_instance" {
  source         = "./modules/ec2"
  ami_id         = data.aws_ami.amazon-linux.id
  disk_space     = 42
  key_name       = module.worker_keypair.name
  name           = local.worker_instance_name
  size           = "t3.medium"
  startup_script = data.template_file.user_data_script.rendered
}

module "worker_keypair" {
  source   = "./modules/keypair"
  key_name = local.worker_instance_name
}

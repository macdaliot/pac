output "kms_key" {
  value = module.kms_key
}

output "terraform_state_bucket" {
  value = module.terraform_state_bucket
}

output "worker_instance" {
  value = module.worker_instance
}

output "worker_keypair" {
  value = module.worker_keypair
}

output "worker_private_key" {
  value = module.worker_keypair.private_key
}

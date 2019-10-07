# Creates an SSH keypair that is stored only in Terraform state
# Taint this resource to generate a new keypair
# Resources that use this key pair would also need to get the keypair updated
resource "tls_private_key" "jumpbox" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

# Creates a keypair in AWS with the provided public signature
resource "aws_key_pair" "jumpbox" {
  key_name = "jumpbox-${var.project_name}"

  # This is what gets added to the authorized_keys file
  public_key = tls_private_key.jumpbox.public_key_openssh
}

output "private_key" {
  value = tls_private_key.jumpbox.private_key_pem
  sensitive = true
}

output "public_key" {
  value = tls_private_key.jumpbox.public_key_pem
  sensitive = true
}

output "public_openssh" {
  value = tls_private_key.jumpbox.public_key_openssh
  sensitive = true
}
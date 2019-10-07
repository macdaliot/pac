# Creates an SSH keypair that is stored only in Terraform state
# Taint this resource to generate a new keypair
# Resources that use this key pair would also need to get the keypair updated
resource "tls_private_key" "management" {
  algorithm = "RSA"
  rsa_bits  = 4096
}
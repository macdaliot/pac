# Creates an SSH keypair that is stored only in Terraform state
# Taint this resource to generate a new keypair
# Resources that use this key pair would also need to get the keypair updated

#
# https://www.terraform.io/docs/providers/aws/r/key_pair.html
#
# Creates a keypair in AWS with the provided public signature
#
resource "aws_key_pair" "key" {
  key_name   = var.key_name
  public_key = tls_private_key.key.public_key_openssh
}

resource "tls_private_key" "key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

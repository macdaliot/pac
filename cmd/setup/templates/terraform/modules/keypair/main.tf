# Creates a keypair in AWS with the provided public signature
resource "aws_key_pair" "keypair" {
  key_name = "jumpbox-${var.project_name}"

  # This is what gets added to the authorized_keys file
  public_key = var.public_key_openssh
}
#
# https://www.terraform.io/docs/providers/aws/r/instance.html
#
resource "aws_instance" "instance" {
  ami           = var.ami_id
  instance_type = var.size
  key_name      = var.key_name
  root_block_device {
    volume_size = var.disk_space
  }
  security_groups = [aws_security_group.allow_ssh.name]
  user_data       = var.startup_script
  tags = {
    Name = var.name
  }
}

#
# https://www.terraform.io/docs/providers/aws/r/security_group.html
#
resource "aws_security_group" "allow_ssh" {
  description = "Managed by Terraform for ${var.name} instance"
  name        = var.name
  ingress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 22
    protocol    = "tcp"
    to_port     = 22
  }
  egress {
    cidr_blocks = ["0.0.0.0/0"]
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
  }
  tags = {
    Name = var.name
  }
}

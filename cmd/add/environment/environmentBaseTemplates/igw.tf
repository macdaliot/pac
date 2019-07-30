# IGW for the public subnet
resource "aws_internet_gateway" "application_igw" {
  vpc_id = aws_vpc.application_vpc.id

  tags = {
    Name = "${var.project_name} Jumpbox"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

output "application_igw_id" {
  value = aws_internet_gateway.application_igw.id
}


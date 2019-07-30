# IGW for the public subnet
resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.management_vpc.id

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

output "igw_id" {
  value = aws_internet_gateway.gw.id
}
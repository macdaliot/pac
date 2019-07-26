# IGW for the public subnet
resource "aws_internet_gateway" "application_igw" {
  vpc_id = aws_vpc.application_vpc.id
}

output "application_igw_id" {
  value = aws_internet_gateway.application_igw.id
}


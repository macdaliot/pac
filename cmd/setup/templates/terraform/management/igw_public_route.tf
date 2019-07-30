# Route the public subnet traffic through the IGW
resource "aws_route" "management_internet_access" {
  route_table_id         = aws_vpc.management_vpc.main_route_table_id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.gw.id

  # Tags not supported
}

output "public_igw_route_id" {
  value = aws_route.management_internet_access.id
}


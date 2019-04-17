# Route the public subnet traffic through the IGW
resource "aws_route" "internet_access" {
  route_table_id         = "${var.vpc_main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${var.internet_gateway_id}"
}
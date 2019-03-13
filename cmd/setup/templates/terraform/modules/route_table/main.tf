resource "aws_route_table" "route_table" {
  count  = "${var.az_count}"
  vpc_id = "${var.vpc_id}"

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = "${element(var.nat_gateway_elastic_ips, count.index)}"
  }
}
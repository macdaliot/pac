resource "aws_nat_gateway" "gw" {
  count         = "${var.az_count}"
  subnet_id     = "${element(var.public_subnets, count.index)}"
  allocation_id = "${element(var.elastic_ips, count.index)}"
}
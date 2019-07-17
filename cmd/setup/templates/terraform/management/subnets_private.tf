resource "aws_subnet" "private" {
  count       = "${length(data.aws_availability_zones.available.names)}"
  cidr_block    = "${cidrsubnet(aws_vpc.management_vpc.cidr_block, 8, count.index)}"
  availability_zone = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id      = "${aws_vpc.management_vpc.id}"

  tags {
    name = "${var.project_name}-private-${count.index}"
  }
}

output "private_subnets" {
  value = "${aws_subnet.private.*.cidr_block}"
}

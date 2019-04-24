resource "aws_route_table" "management_vpc" {
  vpc_id = "${aws_vpc.management_vpc.id}"

  route {
    cidr_block = "${aws_vpc.application_vpc.cidr_block}"
    vpc_peering_connection_id = "${aws_vpc_peering_connection.pc.id}"
  }

  tags = {
    Name = "${var.project_name} vpc peering route table"
  }
}

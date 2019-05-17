resource "aws_vpc_peering_connection" "pc" {
  peer_vpc_id   = "${aws_vpc.management_vpc.id}"
  vpc_id        = "${aws_vpc.application_vpc.id}" 
  auto_accept   = true

  # accepter {
  #   allow_remote_vpc_dns_resolution = true
  # }

  # requester {
  #   allow_remote_vpc_dns_resolution = true
  # }

  tags = {
    Name = "${var.project_name} management/application vpc peering"
  }
}

output "vpc_peering_status" {
    value = "${aws_vpc_peering_connection.pc.accept_status}"
}

output "vpc_peering_id" {
    value = "${aws_vpc_peering_connection.pc.id}"
}
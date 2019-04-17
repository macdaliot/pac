
#
# https://www.terraform.io/docs/providers/aws/r/vpc_peering.html
# Peer VPCs
#
resource "aws_vpc_peering_connection" "pc" {
  peer_vpc_id   = "${var.peer_vpc_id}"
  vpc_id        = "${var.vpc_id}" 
  auto_accept   = true

  # accepter {
  #   allow_remote_vpc_dns_resolution = true
  # }

  # requester {
  #   allow_remote_vpc_dns_resolution = true
  # }

  tags = {
    Name = "management/application vpc peering"
  }
}
# Creates the route53 hosted zone and NS records for the project#
# https://www.terraform.io/docs/providers/aws/r/route53_zone.html
#
data "aws_route53_zone" "primary" {
  name = "${var.hosted_zone}"
  private_zone = false
}

resource "aws_route53_zone" "main" {
  name = "${var.project_name}.${var.hosted_zone}"
}

resource "aws_route53_record" "ns" {
  zone_id = "${data.aws_route53_zone.primary.zone_id}"
  name    = "${var.project_name}.${var.hosted_zone}"
  type    = "NS"
  ttl     = "30" # default 30, why so long?

  records = [
    "${aws_route53_zone.main.name_servers.0}",
    "${aws_route53_zone.main.name_servers.1}",
    "${aws_route53_zone.main.name_servers.2}",
    "${aws_route53_zone.main.name_servers.3}"
  ]
}

#Fetch AZs in the current region
data "aws_availability_zones" "available" {}

resource "aws_vpc" "management" {
  cidr_block = "10.0.0.0/16"

  tags {
    name = "${var.project_name}-management-vpc"
  }
}

# Create var.az_count private subnets, each in a different AZ
resource "aws_subnet" "private" {
  count             = "${var.az_count}"
  cidr_block        = "${cidrsubnet(aws_vpc.management.cidr_block, 8, count.index)}"
  availability_zone = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id            = "${aws_vpc.management.id}"

  tags {
    name = "${var.project_name}-private-${count.index}"
  }
}

resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}"
}

module "peer_vpcs" {
  source = "./modules/vpc/peering_connection"

  peer_vpc_id = "${aws_vpc.management.id}"
  vpc_id      = "${aws_vpc.application.id}"
}

resource "aws_route_table" "management_vpc" {
  vpc_id = "${aws_vpc.management.id}"

  route {
    cidr_block = "${aws_vpc.application.cidr_block}"
    vpc_peering_connection_id = "${module.peer_vpcs.id}"
  }

  tags = {
    Name = "${var.project_name} vpc peering route table"
  }
}
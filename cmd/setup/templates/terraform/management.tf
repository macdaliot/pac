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

resource "aws_vpc" "application" {
  cidr_block = "10.1.0.0/16"

  tags {
    name = "${var.project_name}-application-vpc"
  }
}

# Create var.az_count public subnets, each in a different AZ
resource "aws_subnet" "public" {
  count                   = "${var.az_count}"
  cidr_block              = "${cidrsubnet(aws_vpc.application.cidr_block, 8, var.az_count + count.index)}"
  availability_zone       = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id                  = "${aws_vpc.application.id}"
  map_public_ip_on_launch = false

  tags {
    name = "${var.project_name}-public-${count.index}"
  }
}

module "peer_vpcs" {
  source = "./modules/vpc/peering_connection"

  peer_vpc_id = "${aws_vpc.management.id}"
  vpc_id      = "${aws_vpc.application.id}"
}

# Add route from application vpc to management vpc using vpc peering connection
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

# Add route from management vpc to application vpc using vpc peering connection
resource "aws_route_table" "application_vpc" {
  vpc_id = "${aws_vpc.application.id}"

  route {
    cidr_block = "${aws_vpc.management.cidr_block}"
    vpc_peering_connection_id = "${module.peer_vpcs.id}"
  }

  tags = {
    Name = "${var.project_name} vpc peering route table"
  }
}

# IGW for the public subnet
resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.application.id}"
}

# Route the public subnet traffic through the IGW
resource "aws_route" "internet_access" {
  route_table_id         = "${aws_vpc.application.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.gw.id}"
}

resource "aws_alb" "main" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.lb.id}"]
}

variable "cnames" {
  default = ["api", "jenkins", "selenium", "sonarqube"]
}

#
# https://www.terraform.io/docs/providers/aws/r/route53_record.html
#
#
# These records are created in application VPC instead of the management VPC because they requie a load balancer which
# is created in the application VPC per design requirements to speed up setup
#
resource "aws_route53_record" "record" {
  

  count   = "${length(var.cnames)}"
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "${element(var.cnames,count.index)}"
  type    = "CNAME"
  ttl     = "60"
  records = ["${aws_alb.main.dns_name}"]
}


resource "aws_ecs_cluster" "main" {
  name = "${var.project_name}"
}
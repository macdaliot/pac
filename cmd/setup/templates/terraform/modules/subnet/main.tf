# Fetch AZs in the current region
data "aws_availability_zones" "available" {}

# Create var.az_count private subnets, each in a different AZ
resource "aws_subnet" "private" {
  count                   = "${length(data.aws_availability_zones.available.names)}"
  cidr_block              = "${cidrsubnet(var.vpc_cidr_block, 8, count.index)}"
  availability_zone       = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id                  = "${var.vpc_id}"
}

resource "aws_subnet" "public" {
  count                   = "${length(data.aws_availability_zones.available.names)}"
  cidr_block              = "${cidrsubnet(var.vpc_cidr_block, 8, var.az_count + count.index)}"
  availability_zone       = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id                  = "${var.vpc_id}"
  map_public_ip_on_launch = "${var.map_public_ip_on_launch}"
}
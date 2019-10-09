# Create var.az_count public subnets, each in a different AZ
resource "aws_subnet" "public" {
  count = length(data.aws_availability_zones.available.names)
  cidr_block = cidrsubnet(
    aws_vpc.application_vpc.cidr_block,
    8,
    length(data.aws_availability_zones.available.names) + count.index,
  )
  availability_zone       = data.aws_availability_zones.available.names[count.index]
  vpc_id                  = aws_vpc.application_vpc.id
  map_public_ip_on_launch = false

  tags = {
    Name = "${var.project_name}-public-${count.index}"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

output "public_subnet" {
  value = aws_subnet.public
}

output "public_subnets" {
  value = aws_subnet.public.*.cidr_block
}


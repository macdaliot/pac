resource "aws_vpc" "management_vpc" {
  cidr_block = var.management_cidr_block
  tags = {
    Name = "${var.project_name}-management-vpc"
  }
}

output "management_vpc_cidr_block" {
  value = aws_vpc.management_vpc.cidr_block
}

output "management_vpc_id" {
  value = aws_vpc.management_vpc.id
}


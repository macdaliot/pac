resource "aws_vpc" "application_vpc" {
  cidr_block = "${var.application_cidr_block}"

  tags {
    Name = "${var.project_name}-${var.environment_name}-vpc"
  }
}

output "application_vpc_cidr_block" {
  value = "${aws_vpc.application_vpc.cidr_block}"
}

output "application_vpc_id" {
  value = "${aws_vpc.application_vpc.id}"
}

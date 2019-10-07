#----------------------------------------------------------------------------------------------------------------------
# VPC
#----------------------------------------------------------------------------------------------------------------------
resource "aws_vpc" "management_vpc" {
  cidr_block = var.management_cidr_block

  tags = {
    Name = "${var.project_name}-management-vpc"
    pac-project-name = var.project_name
    environment      = "management"
  }
}

#----------------------------------------------------------------------------------------------------------------------
# S3 VPC ENDPOINT
#----------------------------------------------------------------------------------------------------------------------
resource "aws_vpc_endpoint" "s3" {
  vpc_id       = "${aws_vpc.management_vpc.id}"
  service_name = "com.amazonaws.${var.region}.s3"
}

#----------------------------------------------------------------------------------------------------------------------
# VPC FLOW LOGS
#----------------------------------------------------------------------------------------------------------------------
module "management_vpc_flow_log" {
  source = "../modules/vpc_flow_log"

  iam_role_arn  = aws_iam_role.vpc_flow_log.arn
  log_group_arn = aws_cloudwatch_log_group.practiceone_log_group.arn
  project_name  = var.project_name
  vpc_id        = aws_vpc.management_vpc.id
}

#----------------------------------------------------------------------------------------------------------------------
# OUTPUTS
#----------------------------------------------------------------------------------------------------------------------

output "management_vpc_cidr_block" {
  value = aws_vpc.management_vpc.cidr_block
}

output "management_vpc_id" {
  value = aws_vpc.management_vpc.id
}
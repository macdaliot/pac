resource "aws_vpc" "application_vpc" {
  cidr_block = var.application_cidr_block
  enable_dns_support  = "true"
  enable_dns_hostnames= "true"

  tags = {
    Name = "${var.project_name}-${var.environment_name}-vpc"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

resource "aws_vpc_endpoint" "dynamodb" {
  vpc_id = aws_vpc.application_vpc.id
  service_name = "com.amazonaws.${var.region}.dynamodb"
  route_table_ids = [aws_vpc.application_vpc.main_route_table_id]

  tags = {
    Name = "${var.project_name} dynamodb vpc endpoint"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

resource "aws_vpc_endpoint" "s3" {
  vpc_id = aws_vpc.application_vpc.id
  service_name = "com.amazonaws.${var.region}.s3"
  route_table_ids = [aws_vpc.application_vpc.main_route_table_id]
  
  tags = {
    Name = "${var.project_name} S3 vpc endpoint"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

resource "aws_vpc_endpoint" "sagemaker_api" {
  vpc_id       = "${aws_vpc.application_vpc.id}"
  service_name = "com.amazonaws.${var.region}.sagemaker.api"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = "true"
  subnet_ids = [
    aws_subnet.public[0].id,
    aws_subnet.public[1].id,
    aws_subnet.public[2].id
  ]
  security_group_ids = [
    aws_security_group.application_lb.id
  ]
}

resource "aws_vpc_endpoint" "sagemaker_runtime" {
  vpc_id       = "${aws_vpc.application_vpc.id}"
  service_name = "com.amazonaws.${var.region}.sagemaker.runtime"
  vpc_endpoint_type = "Interface"
  private_dns_enabled = "true"
  subnet_ids = [
    aws_subnet.public[0].id,
    aws_subnet.public[1].id,
    aws_subnet.public[2].id
  ]
  security_group_ids = [
    aws_security_group.application_lb.id
  ]

  tags = {
    Name = "${var.project_name}"
  }
}

#----------------------------------------------------------------------------------------------------------------------
# VPC FLOW LOGS
#----------------------------------------------------------------------------------------------------------------------
module "development_vpc_flow_log" {
  source = "../modules/vpc_flow_log"

  iam_role_arn  = data.terraform_remote_state.management.outputs.vpc_flow_log_role.arn
  log_group_arn = data.terraform_remote_state.management.outputs.practiceone_log_group_arn
  project_name  = var.project_name
  vpc_id        = aws_vpc.application_vpc.id
}

#----------------------------------------------------------------------------------------------------------------------
# OUTPUTS
#----------------------------------------------------------------------------------------------------------------------
output "application_vpc_cidr_block" {
  value = aws_vpc.application_vpc.cidr_block
}

output "application_vpc_id" {
  value = aws_vpc.application_vpc.id
}

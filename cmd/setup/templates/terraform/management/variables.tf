#----------------------------------------------------------------------------------------------------------------------
# AWS
#----------------------------------------------------------------------------------------------------------------------
data "aws_availability_zones" "available" {}

variable "region" {
    default = "{{ .region }}"
}

#----------------------------------------------------------------------------------------------------------------------
# CLOUDWATCH
#----------------------------------------------------------------------------------------------------------------------

variable "cwl_retention" {
    description = "Cloudwatch Log rention time in days"
    default = 1
}

#----------------------------------------------------------------------------------------------------------------------
# DNS
#----------------------------------------------------------------------------------------------------------------------
variable "hosted_zone" {
    description = "DNS primary zone. Value can be found in terraform.tfvars"
}

#----------------------------------------------------------------------------------------------------------------------
# ELASTIC CONTAINER SERVICE
#----------------------------------------------------------------------------------------------------------------------

variable "execution_role_arn" {
    description = "Role with polices to execute ECS tasks and access Systems Manager Parameter Store to retrieve secrets"
    default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-{{ .env }}-task-execution"
}

variable "task_role_arn" {
    description = "Role with EC2, S3, and ECS access policies for ECS tasks"
    default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-{{ .env }}-jenkins"
}

variable "app_count" {
    description = "The number instances of a task to keep running at all times"
    default = 1 # the tasks are in the management VPC of which there is only one
}

#----------------------------------------------------------------------------------------------------------------------
# ELASTICSEARCH
#----------------------------------------------------------------------------------------------------------------------
variable "es_version" {
  default = "7.1"
}

variable "es_instance_type" {
  default = "r4.large.elasticsearch"
}

variable "es_automated_snapshot_start_hour" {
  default = 23
}

#----------------------------------------------------------------------------------------------------------------------
# JUMPBOX / BASTION HOST
#----------------------------------------------------------------------------------------------------------------------
# variable "end_user_cidr" {
#     default = "{{ .endUserIP }}/32"
#     description = "The IP address that will be added to the jumpbox security group to allow access to the Kibana endpoint."
# }

variable "jumpbox_ami" { default = "amzn2-ami-hvm-2.0.20190508-x86_64-gp2" }

#----------------------------------------------------------------------------------------------------------------------
# PROJECT
#----------------------------------------------------------------------------------------------------------------------
variable "project_name" {
    description = "project name"
}

variable "env" {
    description = "build environment (i.e.: dev, stage production)"
    default = "dev"
}

#----------------------------------------------------------------------------------------------------------------------
# SSH KEYPAIR
#----------------------------------------------------------------------------------------------------------------------
variable "enable_keypair_output" {
    default = "false"
}

#----------------------------------------------------------------------------------------------------------------------
# VPC
#----------------------------------------------------------------------------------------------------------------------
variable "management_cidr_block" {
    default = "{{ .awsManagementVpcCidrBlock }}"
}

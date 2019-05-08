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
    default = "pac.pyramidchallenges.com"
}

#----------------------------------------------------------------------------------------------------------------------
# ELASTIC CONTAINER SERVICE
#----------------------------------------------------------------------------------------------------------------------

variable "execution_role_arn" {
    description = "Role with polices to execute ECS tasks and access Systems Manager Parameter Store to retrieve secrets"
    # default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-{{ .env }}-task-execution"
    default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-dev-task-execution"
}

variable "task_role_arn" {
    description = "Role with EC2, S3, and ECS access policies for ECS tasks"
    # default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-{{ .env }}-jenkins"
    default = "arn:aws:iam::{{ .awsID }}:role/{{ .projectName }}-dev-jenkins"
}

variable "app_count" {
    description = "The number instances of a task to keep running at all times"
    default = 1 # the tasks are in the management VPC of which there is only one
}

#----------------------------------------------------------------------------------------------------------------------
# ELASTICSEARCH
#----------------------------------------------------------------------------------------------------------------------
variable "enable_elasticsearch" {
  description = "Whether or not to create Elasticsearch service"
  default = "false"
}

variable "es_version" {
    default = "6.5"
}
variable "es_instance_type" {
    default = "r4.large.elasticsearch"
}

variable "es_automated_snapshot_start_hour" {
    default = 23
}

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
# S3
#----------------------------------------------------------------------------------------------------------------------

variable "enable_demo_bucket" {
    default = "false"
}

variable "enable_integration_bucket" {
    default = "true"
}

#----------------------------------------------------------------------------------------------------------------------
# VPC
#----------------------------------------------------------------------------------------------------------------------
variable "management_cidr_block" {
    default = "10.101.0.0/16"
}

variable "application_cidr_block" {
    default = "10.201.0.0/16"
}

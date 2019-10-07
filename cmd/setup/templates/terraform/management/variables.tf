#----------------------------------------------------------------------------------------------------------------------
# ACME
#----------------------------------------------------------------------------------------------------------------------
variable "acme_registration_email" { default = "labs@pyramidsystems.com" }

#----------------------------------------------------------------------------------------------------------------------
# AMI IMAGES
#----------------------------------------------------------------------------------------------------------------------
# this should get the latest Amazon Linux 2 ami id for the current region
# example: data.aws_ami.amazon-linux-2.id
#
data "aws_ami" "amazon-linux-2" {
 most_recent = true


 filter {
   name   = "owner-alias"
   values = ["amazon"]
 }


 filter {
   name   = "name"
   values = ["amzn2-ami-hvm*"]
 }

 owners = ["amazon"]
}

data "aws_ami" "ubuntu" {
    most_recent = true

    filter {
        name   = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-xenial-16.04-amd64-server-*"]
    }

    filter {
        name   = "virtualization-type"
        values = ["hvm"]
    }

    owners = ["099720109477"] # Canonical
}

#----------------------------------------------------------------------------------------------------------------------
# AWS
#----------------------------------------------------------------------------------------------------------------------
data "aws_availability_zones" "available" {}

variable "account_id" {
    description = "AWS Account Id"
}

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

variable "kibana_username" {}

#----------------------------------------------------------------------------------------------------------------------
# JUMPBOX / BASTION HOST
#----------------------------------------------------------------------------------------------------------------------
variable "enable_jumpbox" {
    default = false
    description = "Enable the ec2 instance used to connect to resources inside the VPC not directly accessible via the Internet."
}

variable "end_user_cidr" {
    default = "0.0.0.0/0"
    description = "The IP address that will be added to the jumpbox security group to allow access to the Kibana endpoint."
}

#----------------------------------------------------------------------------------------------------------------------
# PROJECT
#----------------------------------------------------------------------------------------------------------------------
variable "env" {
    description = "build environment (i.e.: dev, stage production)"
    default = "dev"
}

variable "project_name" {
    description = "project name"
}

variable project_fqdn {
  description = "Fully-qualified domain name of the project"
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
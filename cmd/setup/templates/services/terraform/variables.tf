#----------------------------------------------------------------------------------------------------------------------
# AWS
#----------------------------------------------------------------------------------------------------------------------
data "aws_availability_zones" "available" {}

variable "region" {
    default = "{{ .region }}"
}


#----------------------------------------------------------------------------------------------------------------------
# DNS
#----------------------------------------------------------------------------------------------------------------------
variable "hosted_zone" {
    default = "pac.pyramidchallenges.com"
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
# JUMPBOX / BASTION HOST
#----------------------------------------------------------------------------------------------------------------------
variable "enable_jumpbox" {
    default = "false"
    description = "Enable the ec2 instance used to connect to resources inside the VPC not directly accessible via the Internet."
}

variable "end_user_cidr" {
    default = "{{ .endUserIP }}/32"
    description = "The IP address that will be added to the jumpbox security group to allow access to the Kibana endpoint."
}

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
# VPC
#----------------------------------------------------------------------------------------------------------------------

variable "application_cidr_block" {
    default = "{{ .awsApplicationVpcCidrBlock }}"
}

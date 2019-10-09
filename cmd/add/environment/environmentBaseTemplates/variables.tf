#----------------------------------------------------------------------------------------------------------------------
# AWS
#----------------------------------------------------------------------------------------------------------------------
data "aws_availability_zones" "available" {}

variable "region" {
  description = "AWS region to deploy to"
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
  description = "Active AWS hosted zone that records will be added to for the project"
}

#----------------------------------------------------------------------------------------------------------------------
# ELASTICSEARCH
#----------------------------------------------------------------------------------------------------------------------
variable "enable_elasticsearch" {
  description = "Whether or not to create Elasticsearch service"
}

variable "es_version" {
  default = "6.5"
  description = "Version of ElasticSearch service (only used if enable_elasticsearch is set to 'true')"
}

variable "es_instance_type" {
  description = "AWS instance size to be used for ElasticSearch (only used if enable_elasticsearch is set to 'true')"
}

variable "es_automated_snapshot_start_hour" {
  default = 23
  description = "Hour (in 24-hour time) that ElasticSearch will create a snapshot"
}

#----------------------------------------------------------------------------------------------------------------------
# JUMPBOX / BASTION HOST
#----------------------------------------------------------------------------------------------------------------------
variable "enable_jumpbox" {
  description = "Enable the ec2 instance used to connect to resources inside the VPC not directly accessible via the Internet."
}

variable "end_user_cidr" {
  description = "The IP address that will be added to the jumpbox security group to allow access to the Kibana endpoint."
}

variable "jumpbox_ami" {
  default = "amzn2-ami-hvm-2.0.20190508-x86_64-gp2"
  description = "Amazon Machine Image for jumpboxes"
}

#----------------------------------------------------------------------------------------------------------------------
# PROJECT
#----------------------------------------------------------------------------------------------------------------------
variable "project_name" {
  description = "The project name"
}

variable "project_fqdn" {
  description = "The fully qualified domain name of the project"
}

variable "environment_name" {
  description = "This environment's name / the name of this VPC"
}

variable "environment_abbr" {
  description = "This environment's abbreviation (used in namespacing AWS resources)"
}

output "environment_abbr" {
  value = var.environment_abbr
}

variable "tracing_active" {}

#----------------------------------------------------------------------------------------------------------------------
# VPC
#----------------------------------------------------------------------------------------------------------------------

variable "application_cidr_block" {
  default = "10.1.0.0/16"
  description = "CIDR block for this application VPC"
}

#----------------------------------------------------------------------------------------------------------------------
# DocumentDb
#----------------------------------------------------------------------------------------------------------------------
variable "enable_documentdb" {
  description = "Whether or not to create Amazon DocumentDb"
}
variable "documentdb_user" {
  description = "master username for DocumentDB cluster"
}
variable "documentdb_password" {
  description = "master password for DocumentDB cluster"
}

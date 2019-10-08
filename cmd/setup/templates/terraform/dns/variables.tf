variable "project_name" {
    default = "{{.projectName}}"
}

variable "hosted_zone" {
    default = "pac.pyramidchallenges.com"
}

variable "region" {
    default = "{{.region}}"
}

variable "environment_name" {}

#----------------------------------------------------------------------------------------------------------------------
# VPC
#----------------------------------------------------------------------------------------------------------------------
variable "management_cidr_block" {
    default = "{{.awsManagementVpcCidrBlock}}"
}

variable "application_cidr_block" {
    default = "{{.awsApplicationVpcCidrBlock}}"
}
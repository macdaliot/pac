variable "aws_access_key_id" {
  description = "AWS access key. Should not be set in terraform.tfvars to be kept a secret"
}

variable "aws_account_number" {
  description = "AWS accout number. Should not be set in terraform.tfvars to be kept a secret"
}

variable "aws_secret_access_key" {
  description = "AWS secret access key. Should not be set in terraform.tfvars to be kept a secret"
}

variable "github_auth" {
  description = "URL encoded version of GitHub username and password. Should not be set in terraform.tfvars to be kept a secret"
}

variable "github_password" {
  description = "Password to the PyramidSystemsInc GitHub organization. Should not be set in terraform.tfvars to be kept a secret"
}

variable "github_username" {
  description = "Username to the PyramidSystemsInc GitHub organization. Should not be set in terraform.tfvars to be kept a secret"
}

variable "hosted_zone" {
  description = "Pre-existing hosted zone domain where the deployed application will become a subdomain"
}

variable "jenkins_password" {
  description = "Password to login to the Jenkins instance. Should not be set in terraform.tfvars to be kept a secret"
}

variable "jenkins_username" {
  description = "Username to login to the Jenkins instance. Should not be set in terraform.tfvars to be kept a secret"
}

variable "jwt_issuer" {
  description = ""
}

variable "jwt_secret" {
  description = ""
}

variable "iam_user_password" {
  description = "Default password for a new IAM user. Should not be set in terraform.tfvars to be kept a secret"
}

variable "management_cidr_block" {
  description = "CIDR block for the VPC"
}

variable "postgres_password" {
  description = ""
}

variable "project_name" {
  description = "User-friendly name of the application. Must be all lowercase with no symbols (used in the URL of the application)"
}

variable "psi_aws_access_key_id" {
  description = "AWS access key for the pyramidsystemsinc account. Should not be set in terraform.tfvars to be kept a secret"
}

variable "psi_aws_account_number" {
  description = "AWS accout number for the pyramidsystemsinc account. Should not be set in terraform.tfvars to be kept a secret"
}

variable "psi_aws_secret_access_key" {
  description = "AWS secret access key for the pyramidsystemsinc account. Should not be set in terraform.tfvars to be kept a secret"
}

variable "region" {
  description = "AWS region where the application is to be deployed (i.e. us-east-1)"
}

variable "sonar_jdbc_password" {
  description = ""
}

variable "sonar_jdbc_username" {
  description = ""
}

variable "sonar_password" {
  description = ""
}

variable "sonar_secret" {
  description = "Secret token for SonarQube. Should not be set in terraform.tfvars to be kept a secret"
}

variable "sonar_username" {
  description = ""
}

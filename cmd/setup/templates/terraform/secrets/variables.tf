variable "aws_access_key_id" {}
variable "aws_secret_access_key" {}
variable "github_password" {}
variable "github_username" {}
variable "jenkins_password" {}
variable "jenkins_username" {}
variable "jwt_issuer" {}
variable "jwt_secret" {}
variable "postgres_password" {}
variable "sonar_jdbc_password" {}
variable "sonar_jdbc_username" {}
variable "sonar_password" {}
variable "sonar_secret" {}
variable "sonar_username" {}
variable "project_name" {
  default = "[psi[.projectName]]"
}

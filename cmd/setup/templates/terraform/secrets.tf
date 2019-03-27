#
# https://www.terraform.io/docs/providers/random/index.html
#
provider "random" {
  version = "2.1"
}

variable "secrets" {
  default = [
    "jwt_secret",
    "postgres_password",
    "sonar_jdbc_password",
    "sonar_jdbc_username",
  ]
}

#
# https://www.terraform.io/docs/providers/random/r/string.html
#
resource "random_string" "password" {
  count = "${length(var.secrets)}"
  length = 20
  special = false
}

#
# https://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "jwt_issuer" {
  name  = "JWT_ISSUER"
  description = "JWT Issuer"
  type  = "SecureString"
  value = "urn:pacAuth"
}

resource "aws_ssm_parameter" "jwt_secret" {
  name  = "JWT_SECRET"
  description = "Test password"
  type  = "SecureString"
  value = "${random_string.password.0.result}"
}
resource "aws_ssm_parameter" "postgres_password" {
  name  = "POSTGRES_PASSWORD"
  description = "PostGRES password"
  type  = "SecureString"
  value = "${random_string.password.1.result}"
}
resource "aws_ssm_parameter" "sonar_jdbc_password" {
  name  = "SONAR_JDBC_PASSWORD"
  description = "Sonar JDBC password"
  type  = "SecureString"
  value = "sonar"
}

resource "aws_ssm_parameter" "sonar_jdbc_url" {
  name  = "SONAR_JDBC_URL"
  description = "Sonar JDBC URL"
  type  = "SecureString"
  value = "jdbc:postgresql://localhost/sonar"
}

resource "aws_ssm_parameter" "sonar_jdbc_username" {
  name  = "SONAR_JDBC_USERNAME"
  description = "Sonar JDBC Username"
  type  = "SecureString"
  value = "${random_string.password.3.result}"
}

output "jwt_issuer" {
  value = "${aws_ssm_parameter.jwt_issuer.value}"
}

output "jwt_secret" {
  value = "${aws_ssm_parameter.jwt_secret.value}"
}

output "postgres_password" {
  value = "${aws_ssm_parameter.postgres_password.value}"
}


output "sonar_jdbc_username" {
  value = "${aws_ssm_parameter.sonar_jdbc_username.value}"
}

output "sonar_jdbc_password" {
  value = "${aws_ssm_parameter.sonar_jdbc_password.value}"
}

output "sonar_jdbc_url" {
  value = "${aws_ssm_parameter.sonar_jdbc_url.value}"
}
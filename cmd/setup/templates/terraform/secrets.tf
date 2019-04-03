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
    "sonar_secret",
    "jenkins_password",
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
  name  = "/pac/${var.project_name}/jwt/issuer"
  description = "JWT Issuer"
  type  = "SecureString"
  value = "urn:pacAuth"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "jwt_secret" {
  name  = "/pac/${var.project_name}/jwt/secret"
  description = "Test password"
  type  = "SecureString"
  value = "${random_string.password.0.result}"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}
resource "aws_ssm_parameter" "postgres_password" {
  name  = "/pac/${var.project_name}/postgres/password"
  description = "PostGRES password"
  type  = "SecureString"
  value = "${random_string.password.1.result}"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}
resource "aws_ssm_parameter" "sonar_jdbc_password" {
  name  = "/pac/${var.project_name}/sonar/sonar_jdbc_password"
  description = "Sonar JDBC password"
  type  = "SecureString"
  value = "sonar"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "sonar_jdbc_url" {
  name  = "/pac/${var.project_name}/sonar/sonar_jdbc_url"
  description = "Sonar JDBC URL"
  type  = "SecureString"
  value = "jdbc:postgresql://localhost/sonar"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "sonar_jdbc_username" {
  name  = "/pac/${var.project_name}/sonar/sonar_jdbc_username"
  description = "Sonar JDBC Username"
  type  = "SecureString"
  value = "${random_string.password.3.result}"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "sonar_secret" {
  name  = "/pac/${var.project_name}/sonar/secret"
  description = "Sonar Secret for Jenkins"
  type  = "SecureString"
  value = "${random_string.password.4.result}"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "jenkins_username" {
  name  = "/pac/${var.project_name}/jenkins/username"
  description = "Jenkins default user"
  type  = "SecureString"
  value = "pyramid"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

resource "aws_ssm_parameter" "jenkins_password" {
  name  = "/pac/${var.project_name}/jenkins/password"
  description = "Jenkins default password"
  type  = "SecureString"
  #value = "${random_string.password.5.result}"
  value = "systems"
  key_id = "${data.aws_kms_key.project_key.key_id}"
}

output "jwt_issuer" {
  value = "${aws_ssm_parameter.jwt_issuer.value}"
}

output "jwt_secret" {
  value = "${aws_ssm_parameter.jwt_secret.value}"
}

# output "postgres_password" {
#   value = "${aws_ssm_parameter.postgres_password.value}"
# }


# output "sonar_jdbc_username" {
#   value = "${aws_ssm_parameter.sonar_jdbc_username.value}"
# }

# output "sonar_jdbc_password" {
#   value = "${aws_ssm_parameter.sonar_jdbc_password.value}"
# }

# output "sonar_jdbc_url" {
#   value = "${aws_ssm_parameter.sonar_jdbc_url.value}"
# }

output "jenkins_password" {
  value = "${aws_ssm_parameter.jenkins_password.value}"
}
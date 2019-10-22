#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "aws_access_key_id" {
  name        = "/pac/aws/access_key_id"
  description = "AWS Access Key"
  type        = "SecureString"
  value       = var.aws_access_key_id
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "aws_secret_access_key" {
  name        = "/pac/aws/secret_access_key"
  description = "AWS Secret Access Key"
  type        = "SecureString"
  value       = var.aws_secret_access_key
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "github_password" {
  name        = "/pac/github/password"
  description = "GitHub Password"
  type        = "SecureString"
  value       = var.github_password
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "github_username" {
  name        = "/pac/github/username"
  description = "GitHub Username"
  type        = "SecureString"
  value       = var.github_username
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "jenkins_password" {
  name        = "/pac/${var.project_name}/jenkins/password"
  description = "Jenkins Password"
  type        = "SecureString"
  value       = var.jenkins_password
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "jenkins_username" {
  name        = "/pac/${var.project_name}/jenkins/username"
  description = "Jenkins Username"
  type        = "SecureString"
  value       = var.jenkins_username
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "jwt_issuer" {
  name        = "/pac/${var.project_name}/jwt/issuer"
  description = "JWT Issuer"
  type        = "SecureString"
  value       = var.jwt_issuer
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "jwt_secret" {
  name        = "/pac/${var.project_name}/jwt/secret"
  description = "JWT Secret"
  type        = "SecureString"
  value       = var.jwt_secret
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "postgres_password" {
  name        = "/pac/${var.project_name}/postgres/password"
  description = "PostgreSQL Password"
  type        = "SecureString"
  value       = var.postgres_password
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
} 

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "sonar_jdbc_password" {
  name        = "/pac/${var.project_name}/sonar/jdbc_password"
  description = "Sonar JDBC Password"
  type        = "SecureString"
  value       = var.sonar_jdbc_password
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
} 

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "sonar_jdbc_username" {
  name        = "/pac/${var.project_name}/sonar/jdbc_username"
  description = "Sonar JDBC Password"
  type        = "SecureString"
  value       = var.sonar_jdbc_username
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
} 

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "sonar_password" {
  name        = "/pac/${var.project_name}/sonar/password"
  description = "Sonar Password"
  type        = "SecureString"
  value       = var.sonar_password
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
} 

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "sonar_secret" {
  name        = "/pac/${var.project_name}/sonar/secret"
  description = "Sonar Secret"
  type        = "SecureString"
  value       = var.sonar_secret
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
} 

#
# http://www.terraform.io/docs/providers/aws/r/ssm_parameter.html
#
resource "aws_ssm_parameter" "sonar_username" {
  name        = "/pac/${var.project_name}/sonar/username"
  description = "Sonar Username"
  type        = "SecureString"
  value       = var.sonar_username
  key_id      = data.terraform_remote_state.bootstrap.outputs.kms_key.key.id
  overwrite   = true
  tags = {
    pac-project-name = var.project_name
  }
}

# Get the latest Amazon Linux 2 AMD ID for the current region
data "aws_ami" "amazon-linux" {
  filter {
    name   = "owner-alias"
    values = ["amazon"]
  }
  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
  most_recent = true
  owners      = ["amazon"]
}

# Declare the startup script for the EC2 worker instance
data "template_file" "user_data_script" {
  template = "${file("${path.module}/user-data.sh")}"
  vars = {
    aws_access_key_id         = var.aws_access_key_id
    aws_account_number        = var.aws_account_number
    aws_secret_access_key     = var.aws_secret_access_key
    github_auth               = var.github_auth
    github_password           = var.github_password
    github_username           = var.github_username
    iam_user_password         = var.iam_user_password
    jenkins_password          = var.jenkins_password
    jenkins_username          = var.jenkins_username
    jwt_issuer                = var.jwt_issuer
    jwt_secret                = var.jwt_secret
    postgres_password         = var.postgres_password
    project_name              = var.project_name
    psi_aws_access_key_id     = var.psi_aws_access_key_id
    psi_aws_account_number    = var.psi_aws_account_number
    psi_aws_secret_access_key = var.psi_aws_secret_access_key
    sonar_jdbc_password       = var.sonar_jdbc_password
    sonar_jdbc_username       = var.sonar_jdbc_username
    sonar_password            = var.sonar_password
    sonar_secret              = var.sonar_secret
    sonar_username            = var.sonar_username
  }
}

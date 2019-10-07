locals {
  repos = [
    "{{ .projectName }}-jenkins",
    "{{ .projectName }}-selenium-hub",
    "{{ .projectName }}-selenium-node-chrome",
    "{{ .projectName }}-sonar-db",
    "{{ .projectName }}-sonarqube",
    "{{ .projectName }}-dotnetml",
    "{{ .projectName }}-scikit-sagemaker"
  ]
}

resource "aws_ecr_repository" "repo" {
  count = length(local.repos)
  name  = element(local.repos, count.index)

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

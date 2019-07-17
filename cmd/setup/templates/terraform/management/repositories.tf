variable "enable_create_repos" {
  default = "false"
}

locals {
  repos = [
    "pac-jenkins",
    "pac-selenium-hub",
    "pac-selenium-node-chrome",
    "pac-sonar-db",
    "sonarqube"
  ]
}

resource "aws_ecr_repository" "repo" {
  count = "${var.enable_create_repos == "true" ? length(local.repos) : 0}"
  name  = "${element(local.repos, count.index)}"
}

output "enable_create_repos" {
  value = "${var.enable_create_repos}"
}

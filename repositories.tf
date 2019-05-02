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
    count = "${length(local.repos)}"
    name  = "${element(local.repos, count.index)}"
}
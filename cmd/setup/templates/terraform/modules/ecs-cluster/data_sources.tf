data "terraform_remote_state" "vpc" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/vpc"
        region = "us-east-2"
    }
}

data "terraform_remote_state" "jenkins" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/tasks/jenkins"
        region = "us-east-2"
    }
}
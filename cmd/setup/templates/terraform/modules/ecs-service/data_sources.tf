data "terraform_remote_state" "vpc" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/vpc"
        region = "us-east-2"
    }
}

data "terraform_remote_state" "ecs_cluster" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/ecs_cluster"
        region = "us-east-2"
    }
}


data "terraform_remote_state" "task_jenkins" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/tasks/jenkins"
        region = "us-east-2"
    }
}

data "terraform_remote_state" "alb" {
    backend = "s3"

    config {
        bucket = "pac-terraform-state-dev"
        key    = "state/tasks/alb"
        region = "us-east-2"
    }
}
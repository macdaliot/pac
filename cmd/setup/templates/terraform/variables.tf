variable "hosted_zone" {
    default = "pac.pyramidchallenges.com"
}
variable "project_name" {
    default="{{ .projectName }}"
    description = "project name"
}

# variable "resource_name" {
#     description = "API endpoint resource name to identify lambda function"
# }
variable "az_count" {
    description = "number of availability zones to deploy to"
}

variable "app_count" {
    description = "number of containers to deploy"
}

variable "app_port" {
    default = "8080"
}

variable "fargate_cpu" {
    default = 2048
}

variable "fargate_memory" {
    default = 4096
}

variable "app_image" {
    default = "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins:latest"
}

variable "task_role_arn" {
    default = "arn:aws:iam::118104210923:role/jenkins_instance"
}

variable "execution_role_arn" {
    default = "arn:aws:iam::118104210923:role/ecsTaskExecutionRole"
}

variable "services" {
    type = "list"
}

variable "lambda_function_payload" {
    type = "string"
    description = "Name of code file for lambda function"
    default = "function.zip"
}

variable "create_elasticsearch" {
  description = "Whether or not to create Elasticsearch service"
  default = "false"
}

variable "es_version" {
    default = "6.4"
}
variable "es_instance_type" {
    default = "r4.large.elasticsearch"
}

variable "es_automated_snapshot_start_hour" {
    default = 23
}
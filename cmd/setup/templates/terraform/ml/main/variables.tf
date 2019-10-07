variable "account_id" {}

variable "branch" {}

variable "compute_type" { default = "BUILD_GENERAL1_SMALL" }

variable "container_type" { default = "LINUX_CONTAINER" }

#variable "evaluate_model_project" {}

variable "image_url" {}

variable "enable_cloudwatch_metrics" {}

variable "github_token" {}

variable "github_user" {}

variable "hosted_zone" {}

variable "image" { default = "docker.io/python" }

variable image_repo {}

variable "image_name" {}

variable "image_tag" { default = "3.6.9" }

variable "log_level" {}

variable "project_name" {}

variable "project_fqdn" {}

variable "region" {}

variable "repo" {}

#variable "sagemaker_model" {}

#variable "sagemaker_program" {}
variable sagemaker_image_name {}

#variable "source_directory" {}

variable "stack_name" {}

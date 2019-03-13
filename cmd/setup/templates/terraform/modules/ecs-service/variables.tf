variable "service_name" {}
variable "cluster_id" {}
variable "task_definition_arn" {}
variable "app_count" {}
variable "launch_type" {}
variable "security_group_id" {}
variable "private_subnets" {}
variable "assign_public_ip" {}

#load_balancer
variable "alb_target_group_id" {}
variable "container_name" {}
variable "app_port" {}
variable "depends_on_listener" {}
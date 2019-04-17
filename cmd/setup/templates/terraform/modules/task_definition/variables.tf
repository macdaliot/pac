variable "execution_role_arn" {}
variable "task_role_arn" {}
variable "family" {}
variable "network_mode" {
    default = "awsvpc"
}
variable "requires_compatibilities" {
    type = "list"
    default = ["FARGATE"]
}
variable "cpu" {}
variable "memory" {}
variable "container_definitions" {}
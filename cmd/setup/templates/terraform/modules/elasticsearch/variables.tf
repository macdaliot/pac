variable "project_name" {}
variable "version" {
    default = "6.4"
}
variable "instance_type" {
    default = "r4.large.elasticsearch"
}

variable "automated_snapshot_start_hour" {
    default = 23
}
variable "vpc_id" {}

variable "subnet" {}
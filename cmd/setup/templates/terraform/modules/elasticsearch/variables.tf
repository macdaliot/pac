variable "project_name" {}
variable "es_version" {
    default = "6.4"
}
variable "es_instance_type" {
    default = "r4.large.elasticsearch"
}

variable "es_automated_snapshot_start_hour" {
    default = 23
}
variable "vpc_id" {}

variable "subnet" {}
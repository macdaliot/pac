variable "project_name" {
    description = "Project name"
    type        = "string"
}

variable "billing_mode" {
    default = "PROVISIONED"
}

variable "read_capacity" {
    default = 1000
}

variable "write_capacity" {
    default = 1000
}

variable "hash_key" {
    default = "id"
}

variable "stream_view_type" {
    default = "NEW_IMAGE"
}

variable "stream_enabled" {
    default = true
}

variable "attribute_name" {
    default = "id"
}

variable "attribute_type" {
    default = "S"
}

variable "table_name" {}
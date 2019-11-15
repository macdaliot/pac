variable "attribute_name" {
  default = "id"
}

variable "attribute_type" {
  default = "S"
}

variable "billing_mode" {
  default = "PROVISIONED"
}

variable "hash_key" {
  default = "id"
}

variable "read_capacity" {
  default = 1000
}

variable "stream_view_type" {
  default = "NEW_IMAGE"
}

variable "stream_enabled" {
  default = true
}

variable "table_name" {}

variable "write_capacity" {
  default = 1000
}


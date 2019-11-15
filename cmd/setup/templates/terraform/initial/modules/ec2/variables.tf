variable "ami_id" {}

variable "disk_space" {
  default = 8
  type    = number
}

variable "key_name" {
  default = null
}

variable "name" {}

variable "size" {}

variable "startup_script" {
  default = null
}

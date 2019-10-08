variable "acl" {
    default = "private"
}

variable "enable_versioning" {
    type    = bool
    default = true
}

variable "force_destroy" { default = true }

variable "key_id" {
    description = "AWS KMS key id"
}

variable "project_name" {}

variable "region" {}
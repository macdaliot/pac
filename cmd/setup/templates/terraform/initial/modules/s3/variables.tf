variable "acl" {
  default = "private"
}

variable "bucket_name" {}

variable "enable_versioning" {
  default = true
  type    = bool
}

variable "key_id" {
  description = "AWS KMS key id"
}

variable "region" {}

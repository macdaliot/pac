variable "alias_name" {
  description = "The alias of the key. Must begin with 'alias/'"
}

variable "deletion_window" {
  default     = 7
  description = "Duration in days after which the key is deleted after destruction of the resource. Must be between 7 and 30 days."
}

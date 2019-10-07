variable "deletion_window" {
    description = "Duration in days after which the key is deleted after destruction of the resource, must be between 7 and 30 days."
    default = 7
}

variable "project_name" {}
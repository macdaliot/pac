variable "domain_name" {}
variable "validation_method" {
    default = "DNS"
    description = "Which method to use for validation. DNS or EMAIL are valid, NONE can be used for certificates that were imported into ACM and then into Terraform"
}
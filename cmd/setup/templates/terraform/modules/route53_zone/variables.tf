variable "project_name" {
    description = "Name of project"
}
variable "hosted_zone" {
    default = "pac.pyramidchallenges.com"
    description = "Existing AWS hosted zone FQDN (i.e. pac.pyramidchallenges.com)"
}

variable "region" {
    default = "us-east-2"
    description = "AWS region"
}
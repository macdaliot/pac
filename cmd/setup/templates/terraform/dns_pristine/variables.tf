variable "project_name" {
    default = "{{ .projectName }}"
}

variable "hosted_zone" {
    default = "pac.pyramidchallenges.com"
}

variable "region" {
    default = "{{ .region }}"
}

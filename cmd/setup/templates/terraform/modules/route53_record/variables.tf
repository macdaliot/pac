variable "zone_id" {}
variable "names" {
    type = "list"
}
variable "type" {
    default = "A"
}
variable "ttl" {
    default = "300"
}
variable "records" {
    type = "list"
}
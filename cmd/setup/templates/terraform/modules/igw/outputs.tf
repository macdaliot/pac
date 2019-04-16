output "internet_gateway_id" {
    value = "${aws_internet_gateway.igw.id}"
}

output "elastic_ips" {
    value = "${aws_eip.gw.*.id}"
}
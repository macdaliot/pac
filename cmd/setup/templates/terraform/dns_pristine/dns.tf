# Creates the route53 hosted zone and NS records for the project#
# http://www.terraform.io/docs/providers/aws/r/route53_zone.html
#

resource "aws_route53_zone" "main" {
  name = "${var.project_name}.${var.hosted_zone}"
  force_destroy = true
}

output "main_zone_id" {
  value = "${aws_route53_zone.main.zone_id}"
}
resource "aws_route53_record" "ns" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "${var.project_name}.${var.hosted_zone}"
  type    = "NS"
  ttl     = "30" # default 30, why so long?

  records = [
    "${aws_route53_zone.main.name_servers.0}",
    "${aws_route53_zone.main.name_servers.1}",
    "${aws_route53_zone.main.name_servers.2}",
    "${aws_route53_zone.main.name_servers.3}"
  ]
}

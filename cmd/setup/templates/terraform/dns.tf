# Creates the route53 hosted zone and NS records for the project#
# http://www.terraform.io/docs/providers/aws/r/route53_zone.html
#
# data "aws_route53_zone" "primary" {
#   name = "${var.hosted_zone}"
#   private_zone = false
# }

resource "aws_route53_zone" "main" {
  name = "${var.project_name}.${var.hosted_zone}"
  force_destroy = true
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

resource "aws_route53_record" "integration" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "integration.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_lb.main.dns_name}"
    zone_id                = "${aws_lb.main.zone_id}"
    evaluate_target_health = true
  }
}

# output "integration_dns_record" {
#     value = "${aws_route53_record.integration.name}"
# }

# resource "aws_route53_record" "demo" {
#   zone_id = "${aws_route53_zone.main.zone_id}"
#   name    = "demo.${var.project_name}.${var.hosted_zone}"
#   type    = "A"

#   alias {
#     name                   = "${aws_lb.main.dns_name}"
#     zone_id                = "${aws_lb.main.zone_id}"
#     evaluate_target_health = true
#   }
# }

# output "demo_dns_record" {
#     value = "${aws_route53_record.demo.name}"
# }

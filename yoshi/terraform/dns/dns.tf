# Creates the route53 hosted zone and NS records for the project#
# http://www.terraform.io/docs/providers/aws/r/route53_zone.html
#
data "aws_route53_zone" "primary" {
  name = "${var.hosted_zone}"
  private_zone = false
}

resource "aws_route53_zone" "main" {
  name = "${var.project_name}.${var.hosted_zone}"
  force_destroy = true
}

output "main_zone_id" {
  value = "${aws_route53_zone.main.zone_id}"
}
resource "aws_route53_record" "ns" {
  zone_id = "${data.aws_route53_zone.primary.zone_id}"
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

#----------------------------------------------------------------------------------------------------------------------
# CREATE AND VALIDATE WILDCARD SSL CERT FOR HOSTED ZONE
#----------------------------------------------------------------------------------------------------------------------
provider "aws" {
  alias = "acm"
  region = "${var.region}"
}

provider "aws" {
  alias = "route53"
  region = "${var.region}"
}

variable "domain_names" {
  description = "List of domains to associate with the new certificate. ACM currently supports up to 10 domains, any or all of which can contain wildcards. The first domain should be the primary domain"
  type = "list"
  default = ["yoshi.pac.pyramidchallenges.com", "*.yoshi.pac.pyramidchallenges.com"]
}

data "aws_route53_zone" "zid" {
  #description = "The Route53 zone ID in which to create validation records"
  #default = "${data.aws_route53_zone.primary.zone_id}"
  
  zone_id = "${aws_route53_zone.main.zone_id}"
}

variable "zone_ids" {
  description = "Map of zone IDs indexed by domain name (when issuing a certificate spanning multiple zones)"
  type = "map"
  default = {
    "example.com" = "Z1234567890ABC"
  }
}

resource "aws_acm_certificate" "main" {
  provider = "aws.acm"
  domain_name = "${var.domain_names[0]}"
  subject_alternative_names = "${slice(var.domain_names, 1, length(var.domain_names))}"
  validation_method = "DNS"
  
  tags {
    Name = "${replace(var.domain_names[0], "*.", "star.")}"
    terraform = "true"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "validation" {
  provider = "aws.route53"
  count = "${length(var.domain_names)}"
  name = "${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_name")}"
  type = "${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_type")}"
  # default required for zone_ids lookup because https://github.com/hashicorp/terraform/issues/11574
  zone_id = "${data.aws_route53_zone.zid.zone_id != "" ? data.aws_route53_zone.zid.zone_id : lookup(var.zone_ids, element(var.domain_names, count.index), false)}"
  records = ["${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_value")}"]
  ttl = 60
}

resource "aws_acm_certificate_validation" "main" {
  provider = "aws.acm"
  certificate_arn = "${aws_acm_certificate.main.arn}"
  validation_record_fqdns = ["${aws_route53_record.validation.*.fqdn}"]
}

output "acm_cert_arn" {
  description = "The ARN of the issued certificate"
  value = "${aws_acm_certificate.main.arn}"
}

# resource "aws_route53_record" "integration" {
#   zone_id = "${aws_route53_zone.main.zone_id}"
#   name    = "integration.${var.project_name}.${var.hosted_zone}"
#   type    = "A"

#   alias {
#     name                   = "${aws_lb.main.dns_name}"
#     zone_id                = "${aws_lb.main.zone_id}"
#     evaluate_target_health = true
#   }
# }

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

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
  default = ["{{ .projectName }}.pac.pyramidchallenges.com", "*.{{ .projectName }}.pac.pyramidchallenges.com"]
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

output "arn" {
  description = "The ARN of the issued certificate"
  value = "${aws_acm_certificate.main.arn}"
}
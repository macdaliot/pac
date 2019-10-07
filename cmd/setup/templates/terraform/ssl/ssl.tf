#----------------------------------------------------------------------------------------------------------------------
# CREATE AND VALIDATE WILDCARD SSL CERT FOR HOSTED ZONE
#----------------------------------------------------------------------------------------------------------------------
provider "aws" {
  alias  = "acm"
  region = var.region
}

provider "aws" {
  alias  = "route53"
  region = var.region
}

data "aws_route53_zone" "zid" {
  zone_id = data.terraform_remote_state.dns.outputs.main_zone_id
}

resource "aws_acm_certificate" "main" {
  provider                  = aws.acm
  domain_name               = var.cert_domain_name
  validation_method         = "DNS"

  tags = {
    Name      = replace(var.cert_domain_name, "*.", "star.")
    terraform = "true"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "cert_validation" {
  name    = "${aws_acm_certificate.main.domain_validation_options.0.resource_record_name}"
  type    = "${aws_acm_certificate.main.domain_validation_options.0.resource_record_type}"
  zone_id = data.terraform_remote_state.dns.outputs.main_zone_id
  records = ["${aws_acm_certificate.main.domain_validation_options.0.resource_record_value}"]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "main" {
  certificate_arn         = "${aws_acm_certificate.main.arn}"
  validation_record_fqdns = ["${aws_route53_record.cert_validation.fqdn}"]
}

output "acm_cert" {
  description = "The issued ACM certificate"
  value       = aws_acm_certificate.main
}
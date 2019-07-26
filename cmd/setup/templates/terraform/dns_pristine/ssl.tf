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

variable "domain_names" {
  description = "List of domains to associate with the new certificate. ACM currently supports up to 10 domains, any or all of which can contain wildcards. The first domain should be the primary domain"
  type        = list(string)
  default     = ["{{ .projectName }}.pac.pyramidchallenges.com", "*.{{ .projectName }}.pac.pyramidchallenges.com"]
}

data "aws_route53_zone" "zid" {
  #description = "The Route53 zone ID in which to create validation records"
  #default = "${data.aws_route53_zone.primary.zone_id}"

  zone_id = aws_route53_zone.main.zone_id
}

variable "zone_ids" {
  description = "Map of zone IDs indexed by domain name (when issuing a certificate spanning multiple zones)"
  type        = map(string)
  default = {
    "example.com" = "Z1234567890ABC"
  }
}

resource "aws_acm_certificate" "main" {
  provider                  = aws.acm
  domain_name               = var.domain_names[0]
  subject_alternative_names = slice(var.domain_names, 1, length(var.domain_names))
  validation_method         = "DNS"

  tags = {
    Name      = replace(var.domain_names[0], "*.", "star.")
    terraform = "true"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "validation" {
  provider = aws.route53
  count    = length(var.domain_names)
  name     = aws_acm_certificate.main.domain_validation_options[count.index]["resource_record_name"]
  type     = aws_acm_certificate.main.domain_validation_options[count.index]["resource_record_type"]

  # default required for zone_ids lookup because https://github.com/hashicorp/terraform/issues/11574
  zone_id = data.aws_route53_zone.zid.zone_id != "" ? data.aws_route53_zone.zid.zone_id : lookup(var.zone_ids, element(var.domain_names, count.index), false)
  # TF-UPGRADE-TODO: In Terraform v0.10 and earlier, it was sometimes necessary to
  # force an interpolation expression to be interpreted as a list by wrapping it
  # in an extra set of list brackets. That form was supported for compatibilty in
  # v0.11, but is no longer supported in Terraform v0.12.
  #
  # If the expression in the following list itself returns a list, remove the
  # brackets to avoid interpretation as a list of lists. If the expression
  # returns a single list item then leave it as-is and remove this TODO comment.
  records = [aws_acm_certificate.main.domain_validation_options[count.index]["resource_record_value"]]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "main" {
  provider                = aws.acm
  certificate_arn         = aws_acm_certificate.main.arn
  validation_record_fqdns = aws_route53_record.validation.*.fqdn
}

output "acm_cert_arn" {
  description = "The ARN of the issued certificate"
  value       = aws_acm_certificate.main.arn
}


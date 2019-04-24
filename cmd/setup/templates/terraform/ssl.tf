# disabled until needed due to cost and need to learn how to integrate with other services
#
# https://ifritltd.com/2018/07/02/ssl-termination-with-alb-aws-certificate-manager-and-terraform/
# https://blog.valouille.fr/post/2018-03-22-how-to-use-terraform-to-deploy-an-alb-application-load-balancer-with-multiple-ssl-certificates/
#
# # Creates a security certificate for the domain
# module "acm_cert" {
#   source = "./modules/acm_certificate"

#   domain_name = "${var.project_name}.${var.hosted_zone}"
# }

# module "cert_validation_dns_record" {
#   source  = "./modules/route53_record"

#   name    = "${module.route53_zone.domain_validation_options.0.resource_record_name}"
#   type    = "CNAME"
#   zone_id = "${module.route53_zone.child_zone_id}"
#   # we are not setting subject aleternate names (SANs) so we can safely use element 0
#   records = ["${module.acm_cert.domain_validation_options.0.resource_record_value}"]
#   ttl    = 60
# }

# module "cert_validation" {
#   source = "./modules/acm_certificate/validation"

#   certificate_arn         = "${module.acm_cert.cert.arn}"
#   # we are not setting subject aleternate names (SANs) so we can safely use element 0
#   validation_record_fqdns = ["${module.acm_cert.cert.domain_validation_options.0.resource_record_name}"]
# }

#----------------------------------------------------------------------------------------------------------------------
# CREATE AND VALIDATE WILDCARD SSL CERT FOR HOSTED ZONE
#----------------------------------------------------------------------------------------------------------------------
# provider "aws" {
#   alias = "acm"
#   region = "${var.region}"
# }

# provider "aws" {
#   alias = "route53"
#   region = "${var.region}"
# }

# variable "domain_names" {
#   description = "List of domains to associate with the new certificate. ACM currently supports up to 10 domains, any or all of which can contain wildcards. The first domain should be the primary domain"
#   type = "list"
#   default = ["{{ .projectFQDN }}", "*.{{ .projectFQDN }}"]
# }

# data "aws_route53_zone" "zid" {  
#     zone_id = "${aws_route53_zone.main.zone_id}"
# }

# variable "zone_ids" {
#   description = "Map of zone IDs indexed by domain name (when issuing a certificate spanning multiple zones)"
#   type = "map"
#   default = {
#     "example.com" = "Z1234567890ABC"
#   }
# }

# resource "aws_acm_certificate" "main" {
#   provider = "aws.acm"
#   domain_name = "${var.domain_names[0]}"
#   subject_alternative_names = "${slice(var.domain_names, 1, length(var.domain_names))}"
#   validation_method = "DNS"
#   tags {
#     Name = "${replace(var.domain_names[0], "*.", "star.")}"
#     terraform = "true"
#   }
# }

# resource "aws_route53_record" "validation" {
#   provider = "aws.route53"
#   count = "${length(var.domain_names)}"
#   name = "${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_name")}"
#   type = "${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_type")}"
#   # default required for zone_ids lookup because https://github.com/hashicorp/terraform/issues/11574
#   zone_id = "${data.aws_route53_zone.zid.zone_id != "" ? data.aws_route53_zone.zid.zone_id : lookup(var.zone_ids, element(var.domain_names, count.index), false)}"
#   records = ["${lookup(aws_acm_certificate.main.domain_validation_options[count.index], "resource_record_value")}"]
#   ttl = 60
# }

# resource "aws_acm_certificate_validation" "main" {
#   provider = "aws.acm"
#   certificate_arn = "${aws_acm_certificate.main.arn}"
#   validation_record_fqdns = ["${aws_route53_record.validation.*.fqdn}"]
# }

# output "arn" {
#   description = "The ARN of the issued certificate"
#   value = "${aws_acm_certificate.main.arn}"
# }
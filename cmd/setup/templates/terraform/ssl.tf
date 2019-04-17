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
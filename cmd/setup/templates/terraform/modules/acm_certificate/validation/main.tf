#
# https://www.terraform.io/docs/providers/aws/r/acm_certificate_validation.html
#
resource "aws_acm_certificate_validation" "cert_validation" {
  certificate_arn         = "${aws_acm_certificate.cert.arn}"
  validation_record_fqdns = ["${aws_acm_certificate.cert.domain_validation_options.0.resource_record_name}"]
}
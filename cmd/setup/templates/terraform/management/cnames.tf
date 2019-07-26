variable "cnames" {
  default = ["jenkins", "selenium", "sonarqube"]
}

#
# http://www.terraform.io/docs/providers/aws/r/route53_record.html
#
#
# These records are created in application VPC instead of the management VPC because they requie a load balancer which
# is created in the application VPC per design requirements to speed up setup
#
resource "aws_route53_record" "record" {
  count   = length(var.cnames)
  zone_id = data.terraform_remote_state.dns.outputs.main_zone_id
  name    = element(var.cnames, count.index)
  type    = "CNAME"
  ttl     = "60"
  records = [aws_lb.management.dns_name]
}

output "cnames" {
  value = aws_route53_record.record.*.name
}


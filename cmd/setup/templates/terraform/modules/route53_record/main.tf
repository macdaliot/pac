#
# https://www.terraform.io/docs/providers/aws/r/route53_record.html
#
resource "aws_route53_record" "record" {
  count   = "${length(var.names)}"
  zone_id = "${var.zone_id}"
  name    = "${element(var.names,count.index)}"
  type    = "${var.type}"
  ttl     = "${var.ttl}"
  records = ["${var.records}"]
}
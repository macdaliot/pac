#
# https://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group" "target_group" {
  name        = "pac-${var.project_name}-i-${var.resource_name}"
  # port        = "${var.app_port}"
  # protocol    = "${var.app_protocol}"
  # vpc_id      = "${var.vpc_id}"
  target_type = "${var.target_type}"
}
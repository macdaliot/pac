#
# https://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html
# alb listener rule
#
resource "aws_lb_listener_rule" "rule" {
  listener_arn = "${var.alb_listener_arn}"

  priority = "${var.priority}"

  action {
    type = "forward"
    target_group_arn = "${var.alb_target_group_arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/api/${var.resource_name}"]
  }
}
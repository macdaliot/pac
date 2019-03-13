#
# https://www.terraform.io/docs/providers/aws/r/lb_listener.html
# alb listener
#
resource "aws_alb_listener" "alb_listener" {
  load_balancer_arn = "${var.alb_arn}"
  port              = "${var.alb_listener_port}"
  protocol          = "${var.alb_lister_protocol}"

  default_action {
    target_group_arn = "${var.target_group_arn}"
    type             = "${var.listener_type}"
  }
}
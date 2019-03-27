#
# https://www.terraform.io/docs/providers/aws/r/lb_listener.html
# alb listener
#
resource "aws_lb_listener" "api" {
  load_balancer_arn = "${aws_alb.main.arn}"
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "redirect"

    redirect {
      port        = "80"
      protocol    = "HTTP"
      status_code = "HTTP_301"
      path        = "/api"
      query       = "#{query}"
    }
  }
}
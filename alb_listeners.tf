#
# http://www.terraform.io/docs/providers/aws/r/lb_listener.html
#
resource "aws_lb_listener" "api" {
  load_balancer_arn = "${aws_lb.management.arn}"
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = "${data.terraform_remote_state.dns.acm_cert_arn}"



  default_action {
    type             = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
      path        = "/"
      query       = "#{query}"
    }
  }
}

output "aws_lb_listener_api_arn" {
  value = "${aws_lb_listener.api.arn}"
}

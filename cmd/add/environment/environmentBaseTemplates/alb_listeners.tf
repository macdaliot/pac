#
# http://www.terraform.io/docs/providers/aws/r/lb_listener.html
#
resource "aws_lb_listener" "{{ .environmentName }}" {
  load_balancer_arn = "${aws_lb.application.arn}"
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = "${data.terraform_remote_state.dns.acm_cert_arn}"
  default_action {
    type             = "redirect"
    redirect {
      status_code = "HTTP_301"
      protocol    = "HTTPS"
      host        = "${var.environment_name}.${var.project_fqdn}"
      port        = "443"
      path        = "/#{path}"
      query       = "#{query}"
    }
  }
}

output "aws_lb_listener_{{ .environmentName }}_arn" {
  value = "${aws_lb_listener.{{ .environmentName }}.arn}"
}

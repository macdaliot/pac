#
# http://www.terraform.io/docs/providers/aws/r/lb_listener.html
#
# Tags not supported.
#
resource "aws_lb_listener" "{{.environmentName}}" {
  load_balancer_arn = aws_lb.application.arn
  port              = "443"
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = data.terraform_remote_state.ssl.outputs.acm_cert.arn
  default_action {
    type = "redirect"
    redirect {
      status_code = "HTTP_301"
      protocol    = "HTTPS"
      host        = aws_s3_bucket.{{.environmentName}}.id
      port        = "443"
      path        = "/#{path}"
      query       = "#{query}"
    }
  }
}

output "aws_lb_listener_{{.environmentName}}_arn" {
  value = aws_lb_listener.{{.environmentName}}.arn
}


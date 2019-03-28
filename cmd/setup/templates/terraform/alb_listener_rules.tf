#
# https://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html#condition
#

resource "aws_lb_listener_rule" "jenkins_host" {
  listener_arn = "${aws_lb_listener.api.arn}"
  priority     = 99

  action {
    type             = "forward"
    target_group_arn = "${aws_lb_target_group.jenkins.arn}"
  }

  condition {
    field  = "host-header"
    values = ["jenkins.${var.project_name}.${var.hosted_zone}"]
  }
}

resource "aws_lb_listener_rule" "sonarqube_host" {
  listener_arn = "${aws_lb_listener.api.arn}"
  priority     = 98

  action {
    type             = "forward"
    target_group_arn = "${aws_lb_target_group.sonarqube.arn}"
  }

  condition {
    field  = "host-header"
    values = ["sonarqube.${var.project_name}.${var.hosted_zone}"]
  }
}

resource "aws_lb_listener_rule" "selenium_host" {
  listener_arn = "${aws_lb_listener.api.arn}"
  priority     = 97

  action {
    type             = "forward"
    target_group_arn = "${aws_lb_target_group.selenium.arn}"
  }

  condition {
    field  = "host-header"
    values = ["selenium.${var.project_name}.${var.hosted_zone}"]
  }
}
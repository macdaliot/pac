
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
#http://#{host}:80/api?#{query}
#Redirect all traffic from the ALB to the target group
resource "aws_alb_listener" "jenkins" {
  load_balancer_arn = "${aws_alb.main.id}"
  port              = "8080"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group.jenkins.id}"
    type             = "forward"
  }
}

resource "aws_alb_listener" "sonarqube" {
  load_balancer_arn = "${aws_alb.main.id}"
  port              = "9000"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group.sonarqube.id}"
    type             = "forward"
  }
}

resource "aws_alb_listener" "selenium" {
  load_balancer_arn = "${aws_alb.main.id}"
  port              = "4448"
  protocol          = "HTTP"

  default_action {
    target_group_arn = "${aws_alb_target_group.selenium.id}"
    type             = "forward"
  }
}
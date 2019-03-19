resource "aws_alb_target_group" "jenkins" {
  name        = "${var.project_name}-ecs-jenkins"
  port        = "8080"
  protocol    = "HTTP"
  vpc_id      = "${aws_vpc.application.id}"
  target_type = "ip"

  health_check {
    interval = 300
    path = "/login"
    port = 8080
    timeout = 60
    healthy_threshold = 3
    unhealthy_threshold = 3
    matcher = 200
  }
}

resource "aws_alb_target_group" "sonarqube" {
  name        = "${var.project_name}-ecs-sonarqube"
  port        = "9000"
  protocol    = "HTTP"
  vpc_id      = "${aws_vpc.application.id}"
  target_type = "ip"
}

resource "aws_alb_target_group" "selenium" {
  name        = "${var.project_name}-ecs-selenium"
  port        = "4444"
  protocol    = "HTTP"
  vpc_id      = "${aws_vpc.application.id}"
  target_type = "ip"
}

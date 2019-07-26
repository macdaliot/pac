resource "aws_lb_target_group" "sonarqube" {
  name        = "${var.project_name}-ecs-sonarqube"
  port        = "9000"
  protocol    = "HTTP"
  vpc_id      = aws_vpc.management_vpc.id
  target_type = "ip"
}

output "target_group_sonarqube_arn" {
  value = aws_lb_target_group.sonarqube.arn
}

resource "aws_lb_target_group" "selenium" {
  name        = "${var.project_name}-ecs-selenium"
  port        = "4448"
  protocol    = "HTTP"
  vpc_id      = aws_vpc.management_vpc.id
  target_type = "ip"
}

output "target_group_selenium_arn" {
  value = aws_lb_target_group.selenium.arn
}


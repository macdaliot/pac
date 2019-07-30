resource "aws_lb" "management" {
  name            = "${var.project_name}-management-lb"
  subnets         = aws_subnet.private.*.id
  security_groups = [aws_security_group.management_lb.id]

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

output "alb_management_name" {
  value = aws_lb.management.name
}

output "alb_management_arn" {
  value = aws_lb.management.arn
}
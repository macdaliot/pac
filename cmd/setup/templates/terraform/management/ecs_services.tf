resource "aws_ecs_service" "sonarqube" {
  name                              = "sonarqube-ecs-service"
  cluster                           = aws_ecs_cluster.main.arn
  task_definition                   = aws_ecs_task_definition.sonarqube.arn
  desired_count                     = var.app_count
  launch_type                       = "FARGATE"
  health_check_grace_period_seconds = 30

  network_configuration {
    security_groups  = [aws_security_group.ecs_tasks.id]
    subnets          = aws_subnet.private.*.id
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.sonarqube.id
    container_name   = "sonarqube"
    container_port   = "9000"
  }

  depends_on = [aws_lb_listener.https]
}

resource "aws_ecs_service" "selenium" {
  name                              = "selenium-ecs-service"
  cluster                           = aws_ecs_cluster.main.arn
  task_definition                   = aws_ecs_task_definition.selenium.arn
  desired_count                     = var.app_count
  launch_type                       = "FARGATE"
  health_check_grace_period_seconds = 30

  network_configuration {
    security_groups  = [aws_security_group.ecs_tasks.id]
    subnets          = aws_subnet.private.*.id
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.selenium.id
    container_name   = "pac-selenium-hub-${var.project_name}"
    container_port   = "4448"
  }

  depends_on = [aws_lb_listener.https]
}


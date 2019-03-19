resource "aws_ecs_service" "jenkins" {
  name            = "jenkins-ecs-service"
  cluster         = "${aws_ecs_cluster.main.id}"
  task_definition = "${aws_ecs_task_definition.jenkins.arn}"
  desired_count   = "${var.app_count}"
  launch_type     = "FARGATE"
  health_check_grace_period_seconds = 300

  network_configuration {
    security_groups  = ["${aws_security_group.ecs_tasks.id}"]
    subnets          = ["${aws_subnet.public.*.id}"]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = "${aws_alb_target_group.jenkins.id}"
    container_name   = "pac-jenkins"
    container_port   = "${var.app_port}"
  }

  depends_on = [
    "aws_alb_listener.jenkins"
  ]
}

resource "aws_ecs_service" "sonarqube" {
  name            = "sonarqube-ecs-service"
  cluster         = "${aws_ecs_cluster.main.id}"
  task_definition = "${aws_ecs_task_definition.sonarqube.arn}"
  desired_count   = "${var.app_count}"
  launch_type     = "FARGATE"

  network_configuration {
    security_groups  = ["${aws_security_group.ecs_tasks.id}"]
    subnets          = ["${aws_subnet.public.*.id}"]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = "${aws_alb_target_group.sonarqube.id}"
    container_name   = "sonarqube"
    container_port   = "9000"
  }

  depends_on = [
    "aws_alb_listener.sonarqube"
  ]
}

resource "aws_ecs_service" "selenium" {
  name            = "selenium-ecs-service"
  cluster         = "${aws_ecs_cluster.main.id}"
  task_definition = "${aws_ecs_task_definition.selenium.arn}"
  desired_count   = "${var.app_count}"
  launch_type     = "FARGATE"

  network_configuration {
    security_groups  = ["${aws_security_group.ecs_tasks.id}"]
    subnets          = ["${aws_subnet.public.*.id}"]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = "${aws_alb_target_group.selenium.id}"
    container_name   = "selenium"
    container_port   = "4444"
  }

  depends_on = [
    "aws_alb_listener.selenium"
  ]
}
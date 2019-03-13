resource "aws_ecs_service" "service" {
  name            = "${var.service_name}-service"
  cluster         = "${var.cluster_id}"
  task_definition = "${var.task_definition_arn}"
  desired_count   = "${var.app_count}"
  launch_type     = "${var.launch_type}"

  network_configuration {
    security_groups  = ["${var.security_group_id}"]
    subnets          = ["${var.private_subnets}"]
    assign_public_ip = "${var.assign_public_ip}"
  }

  load_balancer {
    target_group_arn = "${var.alb_target_group_id}"
    container_name   = "{${var.container_name}}"
    container_port   = "${var.app_port}"
  }

  depends_on = [
    "${var.depends_on_listener}"
  ]
}
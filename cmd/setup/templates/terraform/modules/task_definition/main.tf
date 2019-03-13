resource "aws_ecs_task_definition" "task_definition" {
  execution_role_arn       = "${var.execution_role_arn}"
  task_role_arn            = "${var.task_role_arn}"
  family                   = "${var.service_name}"
  network_mode             = "${var.network_mode}"
  requires_compatibilities = "${var.requires_compatibilities}"
  cpu                      = "${var.fargate_cpu}"
  memory                   = "${var.fargate_memory}"

  container_definitions    = "${var.container_definitions}"
}
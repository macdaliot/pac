resource "aws_alb_target_group" "target_group" {
  name        = "${var.project_name}-target-group"
  port        = "${var.app_port}"
  protocol    = "${var.app_protocol}"
  vpc_id      = "${var.vpc_id}"
  target_type = "${var.target_type}"
}

resource "aws_alb_listener" "alb_listener" {
  load_balancer_arn = "${var.alb_arn}"
  port              = "${var.alb_listener_port}"
  protocol          = "${var.alb_lister_protocol}"

  default_action {
    target_group_arn = "${var.target_group_arn}"
    type             = "${var.listener_type}"
  }
}

resource "aws_ecs_task_definition" "task_definition" {
  execution_role_arn       = "${var.execution_role_arn}"
  task_role_arn            = "${var.task_role_arn}"
  family                   = "${var.service_name}"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "${var.fargate_cpu}"
  memory                   = "${var.fargate_memory}"

  container_definitions = <<DEFINITION
[
  {
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins",
    "memory": 4096,
    "name": "pac-jenkins",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 8080
      }
    ]
  }
]
DEFINITION
}

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
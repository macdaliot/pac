# resource "aws_ecs_service" "vault_server" {
#   name            = "vault-ecs-service"
#   cluster         = "${aws_ecs_cluster.main.id}"
#   task_definition = "${aws_ecs_task_definition.vault.arn}"
#   desired_count   = 3
#   launch_type     = "FARGATE"
#   health_check_grace_period_seconds = 300

#   network_configuration {
#     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
#     subnets          = ["${aws_subnet.public.*.id}"]
#     assign_public_ip = true
#   }

#   load_balancer {
#     target_group_arn = "${aws_lb_target_group.vault.id}"
#     container_name   = "pac-vault-${var.project_name}"
#     container_port   = "8500"
#   }

#   depends_on = [
#     "aws_lb_listener.vault"
#   ]
# }

# resource "aws_autoscaling_group" "cluster" {
#   name                 = "terratest"
#   vpc_zone_identifier  = ["${aws_vpc.application_vpc.id}"]
#   launch_configuration = "${aws_launch_configuration.cluster.name}"

#   desired_capacity = 3
#   min_size         = 3
#   max_size         = 3
# }

# resource "aws_launch_configuration" "cluster" {
#   name                 = "terratest"
#   image_id             = "${data.aws_ami.ecs_optimized.id}"
#   iam_instance_profile = "${aws_iam_instance_profile.ecs_agent.name}"
#   user_data            = "${data.template_file.user_data.rendered}"
#   instance_type        = "t2.small"
# }

# data "aws_ami" "ecs_optimized" {
#   owners = ["amazon"]
#   filter {
#     name   = "name"
#     values = ["amzn2-ami-ecs-hvm-2.0.20190301-x86_64-ebs"]
#   }
# }

# data "template_file" "user_data" {
#   template = "${file("user_data.yml")}"

#   vars {
#     ecs_cluster = "terratest"
#   }
# }

# # Define the role.
# resource "aws_iam_role" "ecs_agent" {
#   name               = "ecs-agent"
#   assume_role_policy = "${data.aws_iam_policy_document.ecs_agent.json}"
# }

# # Allow EC2 service to assume this role.
# data "aws_iam_policy_document" "ecs_agent" {
#   statement {
#     actions = ["sts:AssumeRole"]

#     principals {
#       type        = "Service"
#       identifiers = ["ec2.amazonaws.com"]
#     }
#   }
# }

# # Give this role the permission to do ECS Agent things.
# resource "aws_iam_role_policy_attachment" "ecs_agent" {
#   role       = "${aws_iam_role.ecs_agent.name}"
#   policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
# }

# resource "aws_iam_instance_profile" "ecs_agent" {
#   name = "ecs-agent"
#   role = "${aws_iam_role.ecs_agent.name}"
# }

# resource "aws_lb_target_group" "vault" {
#   name        = "${var.project_name}-ecs-vault"
#   port        = "8500"
#   protocol    = "HTTP"
#   vpc_id      = "${aws_vpc.application_vpc.id}"
#   target_type = "ip"


#   health_check {
#     interval = 300
#     path = "/ui"
#     port = 8500
#     timeout = 60
#     healthy_threshold = 3
#     unhealthy_threshold = 3
#     matcher = 200
#   }
# }

#
# https://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html#condition
#
# resource "aws_lb_listener_rule" "vault_host" {
#   listener_arn = "${aws_lb_listener.vault.arn}"
#   priority     = 99

#   action {
#     type             = "forward"
#     target_group_arn = "${aws_lb_target_group.vault.arn}"
#   }

#   condition {
#     field  = "host-header"
#     values = ["vault.${var.project_name}.${var.hosted_zone}"]
#   }
# }

# resource "aws_lb_listener" "vault" {
#   load_balancer_arn = "${aws_lb.main.id}"
#   port              = "8500"
#   protocol          = "HTTP"

#   default_action {
#     target_group_arn = "${aws_lb_target_group.vault.id}"
#     type             = "forward"
#   }
# }

# resource "aws_lb_listener" "vault" {
#   load_balancer_arn = "${aws_lb.main.id}"
#   port              = "8500"
#   protocol          = "HTTP"

#   default_action {
#     target_group_arn = "${aws_lb_target_group.vault.id}"
#     type             = "forward"
#   }
# }

# resource "aws_ecs_task_definition" "vault" {
#   execution_role_arn       = "${var.execution_role_arn}"
#   task_role_arn            = "${var.task_role_arn}"
#   family                   = "pac-vault"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["FARGATE"]
#   cpu                      = "4096"
#   memory                   = "8192"
  

#   container_definitions = <<DEFINITION
# [
#   {
#     "cpu": 1024,
#     "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/vault:srv",
#     "memory": 2048,
#     "name": "pac-vault-${var.project_name}",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 8500
#       }
#     ],
#     "logConfiguration": { 
#       "logDriver": "awslogs",
#       "options": { 
#           "awslogs-group" : "${var.project_name}-log-group",
#           "awslogs-region": "us-east-2",
#           "awslogs-stream-prefix": "ecs"
#       }
#     }
#   }
# ]
# DEFINITION

#     tags {
#         service = "vault-cluster"
#     }
# }
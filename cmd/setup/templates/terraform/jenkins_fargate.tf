# resource "aws_ecs_task_definition" "jenkins" {
#   execution_role_arn       = "${var.execution_role_arn}"
#   task_role_arn            = "${var.task_role_arn}"
#   family                   = "pac-jenkins"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["EC2", "FARGATE"]
#   cpu                      = 2048
#   memory                   = 16384

#   container_definitions = <<DEFINITION
# [
#   {
#     "secrets": [
#       {
#         "name": "JWT_ISSUER",
#         "valueFrom": "/pac/${var.project_name}/jwt/issuer"
#       },
#       {
#         "name": "JWT_SECRET",
#         "valueFrom": "/pac/${var.project_name}/jwt/secret"
#       },
#       {
#         "name": "USERNAME",
#         "valueFrom": "/pac/${var.project_name}/jenkins/username"
#       },
#       {
#         "name": "PASSWORD",
#         "valueFrom": "/pac/${var.project_name}/jenkins/password"
#       },
#       {
#         "name": "SONAR_SECRET",
#         "valueFrom": "/pac/sonar/secret"
#       },
#       {
#         "name": "GITHUB_USERNAME",
#         "valueFrom": "/pac/github/username"
#       },
#       {
#         "name": "GITHUB_PASSWORD",
#         "valueFrom": "/pac/github/password"
#       },
#       {
#         "name": "AWS_ACCESS_KEY_ID",
#         "valueFrom": "/pac/aws/access_id_key"
#       },
#       {
#         "name": "AWS_SECRET_ACCESS_KEY",
#         "valueFrom": "/pac/aws/secret_access_key"
#       }
#     ],
#     "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins:env",
#     "name": "pac-jenkins",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 8080
#       }
#     ],
#     "logConfiguration": { 
#       "logDriver": "awslogs",
#       "options": { 
#           "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
#           "awslogs-region": "${var.region}",
#           "awslogs-stream-prefix": "ecs"
#       }
#     },
#     "entryPoint": ["/runJenkins.sh"]
#   }
# ]
# DEFINITION
# }

# resource "aws_lb_target_group" "jenkins" {
#   name        = "${var.project_name}-ecs-ec2-jenkins"
#   port        = "8080"
#   protocol    = "HTTP"
#   vpc_id      = "${aws_vpc.management_vpc.id}"
#   target_type = "ip"

#   health_check {
#     interval = 300
#     path = "/login"
#     port = 8080
#     timeout = 60
#     healthy_threshold = 3
#     unhealthy_threshold = 3
#     matcher = 200
#   }
# }

# output "target_group_jenkins_arn" {
#   value = "${aws_lb_target_group.jenkins.arn}"
# }

# resource "aws_ecs_service" "jenkins" {
#   name            = "jenkins-ecs-service"
#   cluster         = "${aws_ecs_cluster.main.arn}"
#   task_definition = "${aws_ecs_task_definition.jenkins.arn}"
#   desired_count   = "${var.app_count}"
#   launch_type     = "FARGATE"
#   health_check_grace_period_seconds = 30

#   network_configuration {
#     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
#     subnets          = ["${aws_subnet.private.*.id}"]
#     assign_public_ip = true # need for pulling from ECR registry
#   }

#   load_balancer {
#     target_group_arn = "${aws_lb_target_group.jenkins.id}"
#     container_name   = "pac-jenkins"
#     container_port   = "8080"
#   }

#   depends_on = [
#     "aws_lb_listener.https"
#   ]
# }
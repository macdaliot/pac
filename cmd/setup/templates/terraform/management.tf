# # Creates the route53 hosted zone and NS records for the project
# module "route53_zone" {
#   source = "./modules/route53_zone"

#   project_name = "${var.project_name}"
#   hosted_zone = "${var.hosted_zone}"
# }

# # DO NOT DELETE
# # disabled until needed due to cost and need to learn how to integrate with other services
# #
# # https://ifritltd.com/2018/07/02/ssl-termination-with-alb-aws-certificate-manager-and-terraform/
# # https://blog.valouille.fr/post/2018-03-22-how-to-use-terraform-to-deploy-an-alb-application-load-balancer-with-multiple-ssl-certificates/
# #
# # # Creates a security certificate for the domain
# # module "acm_cert" {
# #   source = "./modules/acm_certificate"

# #   domain_name = "${var.project_name}.${var.hosted_zone}"
# # }

# # module "cert_validation_dns_record" {
# #   source  = "./modules/route53_record"

# #   name    = "${module.route53_zone.domain_validation_options.0.resource_record_name}"
# #   type    = "CNAME"
# #   zone_id = "${module.route53_zone.child_zone_id}"
# #   # we are not setting subject aleternate names (SANs) so we can safely use element 0
# #   records = ["${module.acm_cert.domain_validation_options.0.resource_record_value}"]
# #   ttl    = 60
# # }

# # module "cert_validation" {
# #   source = "./modules/acm_certificate/validation"

# #   certificate_arn         = "${module.acm_cert.cert.arn}"
# #   # we are not setting subject aleternate names (SANs) so we can safely use element 0
# #   validation_record_fqdns = ["${module.acm_cert.cert.domain_validation_options.0.resource_record_name}"]
# # }

# #Fetch AZs in the current region
# data "aws_availability_zones" "available" {}

# resource "aws_vpc" "management" {
#   cidr_block = "10.0.0.0/16"

#   tags {
#     name = "${var.project_name}-management-vpc"
#   }
# }

# # Create var.az_count private subnets, each in a different AZ
# resource "aws_subnet" "private" {
#   count             = "${var.az_count}"
#   cidr_block        = "${cidrsubnet(aws_vpc.management.cidr_block, 8, count.index)}"
#   availability_zone = "${data.aws_availability_zones.available.names[count.index]}"
#   vpc_id            = "${aws_vpc.management.id}"

#   tags {
#     name = "${var.project_name}-private-${count.index}"
#   }
# }

# # ALB Security group
# # This is the group you need to edit if you want to restrict access to your application
# resource "aws_security_group" "lb" {
#   name        = "${var.project_name}-alb"
#   description = "controls access to the ALB"
#   vpc_id      = "${aws_vpc.application.id}"
#   revoke_rules_on_delete = true

#   ingress {
#     protocol    = "tcp"
#     from_port   = "80"
#     to_port     = "80"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   ingress {
#     protocol    = "tcp"
#     from_port   = "8080"
#     to_port     = "8080"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   ingress {
#     protocol    = "tcp"
#     from_port   = "9000"
#     to_port     = "9000"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   ingress {
#     protocol    = "tcp"
#     from_port   = "4445"
#     to_port     = "4445"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }

# # Traffic to the ECS Cluster should only come from the ALB
# resource "aws_security_group" "ecs_tasks" {
#   name        = "${var.project_name}-ecs-tasks"
#   description = "allow inbound access from the ALB only"
#   vpc_id      = "${aws_vpc.application.id}"

#   ingress {
#     protocol    = "tcp"
#     from_port   = "8080"
#     to_port     = "8080"
#     security_groups = ["${aws_security_group.lb.id}"]
#   }

#   ingress {
#     protocol    = "tcp"
#     from_port   = "9000"
#     to_port     = "9000"
#     security_groups = ["${aws_security_group.lb.id}"]
#   }

#   ingress {
#     protocol    = "tcp"
#     from_port   = "4444"
#     to_port     = "4444"
#     security_groups = ["${aws_security_group.lb.id}"]
#   }

#   egress {
#     protocol    = "-1"
#     from_port   = 0
#     to_port     = 0
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }

# resource "aws_alb_target_group" "jenkins" {
#   name        = "${var.project_name}-ecs-jenkins"
#   port        = "8080"
#   protocol    = "HTTP"
#   vpc_id      = "${aws_vpc.application.id}"
#   target_type = "ip"
# }

# resource "aws_alb_target_group" "sonarqube" {
#   name        = "${var.project_name}-ecs-sonarqube"
#   port        = "9000"
#   protocol    = "HTTP"
#   vpc_id      = "${aws_vpc.application.id}"
#   target_type = "ip"
# }

# resource "aws_alb_target_group" "selenium" {
#   name        = "${var.project_name}-ecs-selenium"
#   port        = "4444"
#   protocol    = "HTTP"
#   vpc_id      = "${aws_vpc.application.id}"
#   target_type = "ip"
# }

# #Redirect all traffic from the ALB to the target group
# resource "aws_alb_listener" "jenkins" {
#   load_balancer_arn = "${aws_alb.main.id}"
#   port              = "8080"
#   protocol          = "HTTP"

#   default_action {
#     target_group_arn = "${aws_alb_target_group.jenkins.id}"
#     type             = "forward"
#   }
# }

# resource "aws_alb_listener" "sonarqube" {
#   load_balancer_arn = "${aws_alb.main.id}"
#   port              = "9000"
#   protocol          = "HTTP"

#   default_action {
#     target_group_arn = "${aws_alb_target_group.sonarqube.id}"
#     type             = "forward"
#   }
# }

# resource "aws_alb_listener" "selenium" {
#   load_balancer_arn = "${aws_alb.main.id}"
#   port              = "4445"
#   protocol          = "HTTP"

#   default_action {
#     target_group_arn = "${aws_alb_target_group.selenium.id}"
#     type             = "forward"
#   }
# }

# resource "aws_ecs_cluster" "main" {
#   name = "${var.project_name}"
# }

# resource "aws_ecs_task_definition" "jenkins" {
#   execution_role_arn       = "${var.execution_role_arn}"
#   task_role_arn            = "${var.task_role_arn}"
#   family                   = "pac-jenkins"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["FARGATE"]
#   cpu                      = "${var.fargate_cpu}"
#   memory                   = "${var.fargate_memory}"

#   container_definitions = <<DEFINITION
# [
#   {
#     "cpu": 2048,
#     "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins",
#     "memory": 4096,
#     "name": "pac-jenkins",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 8080
#       }
#     ]
#   }
# ]
# DEFINITION
# }

# resource "aws_ecs_task_definition" "sonarqube" {
#   execution_role_arn       = "${var.execution_role_arn}"
#   task_role_arn            = "${var.task_role_arn}"
#   family                   = "sonarqube"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["FARGATE"]
#   cpu                      = "4096"
#   memory                   = 8192

#   container_definitions = <<DEFINITION
# [
#   {
#     "cpu": 2048,
#     "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/sonarqube",
#     "memory": 4096,
#     "name": "sonarqube",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 9000
#       }
#     ],
#     "environment": [
#       {
#         "name": "sonar.jdbc.password",
#         "value": "sonar"
#       },
#       {
#         "name": "sonar.jdbc.url",
#         "value": "jdbc:postgresql://localhost/sonar"
#       },
#       {
#         "name": "sonar.jdbc.username",
#         "value": "sonar"
#       }
#     ]
#   },
#   {
#     "cpu": 2048,
#     "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-sonar-db",
#     "memory": 4096,
#     "name": "pac-sonar-db",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 5432
#       }
#     ],
#     "environment": [
#       {
#         "name": "POSTGRES_PASSWORD",
#         "value": "pyramid"
#       }
#     ]
#   }
# ]
# DEFINITION
# }

# resource "aws_ecs_task_definition" "selenium" {
#   execution_role_arn       = "${var.execution_role_arn}"
#   task_role_arn            = "${var.task_role_arn}"
#   family                   = "pac-selenium"
#   network_mode             = "awsvpc"
#   requires_compatibilities = ["FARGATE"]
#   cpu                      = "${var.fargate_cpu}"
#   memory                   = "${var.fargate_memory}"

#   container_definitions = <<DEFINITION
# [
#   {
#     "cpu": 2048,
#     "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/selenium",
#     "memory": 4096,
#     "name": "selenium",
#     "networkMode": "awsvpc",
#     "portMappings": [
#       {
#         "containerPort": 4444
#       }
#     ]
#   }
# ]
# DEFINITION
# }

# resource "aws_ecs_service" "jenkins" {
#   name            = "jenkins-ecs-service"
#   cluster         = "${aws_ecs_cluster.main.id}"
#   task_definition = "${aws_ecs_task_definition.jenkins.arn}"
#   desired_count   = "${var.app_count}"
#   launch_type     = "FARGATE"

#   network_configuration {
#     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
#     subnets          = ["${aws_subnet.public.*.id}"]
#     assign_public_ip = true
#   }

#   load_balancer {
#     target_group_arn = "${aws_alb_target_group.jenkins.id}"
#     container_name   = "pac-jenkins"
#     container_port   = "${var.app_port}"
#   }

#   depends_on = [
#     "aws_alb_listener.jenkins"
#   ]
# }

# resource "aws_ecs_service" "sonarqube" {
#   name            = "sonarqube-ecs-service"
#   cluster         = "${aws_ecs_cluster.main.id}"
#   task_definition = "${aws_ecs_task_definition.sonarqube.arn}"
#   desired_count   = "${var.app_count}"
#   launch_type     = "FARGATE"

#   network_configuration {
#     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
#     subnets          = ["${aws_subnet.public.*.id}"]
#     assign_public_ip = true
#   }

#   load_balancer {
#     target_group_arn = "${aws_alb_target_group.sonarqube.id}"
#     container_name   = "sonarqube"
#     container_port   = "9000"
#   }

#   depends_on = [
#     "aws_alb_listener.sonarqube"
#   ]
# }

# resource "aws_ecs_service" "selenium" {
#   name            = "selenium-ecs-service"
#   cluster         = "${aws_ecs_cluster.main.id}"
#   task_definition = "${aws_ecs_task_definition.selenium.arn}"
#   desired_count   = "${var.app_count}"
#   launch_type     = "FARGATE"

#   network_configuration {
#     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
#     subnets          = ["${aws_subnet.public.*.id}"]
#     assign_public_ip = true
#   }

#   load_balancer {
#     target_group_arn = "${aws_alb_target_group.selenium.id}"
#     container_name   = "selenium"
#     container_port   = "4444"
#   }

#   depends_on = [
#     "aws_alb_listener.selenium"
#   ]
# }

# module "roles" {
#   source = "./modules/roles"

#   project_name = "${var.project_name}"
# }

# # DO NOT DELETE
# # disabled until needed due to cost and need to learn how to integrate with other services
# #
# # module "elasticsearch" {
# #   source = "./modules/elasticsearch"

# #   project_name = "${var.project_name}"
# #   vpc_id = "${aws_vpc.management.id}"
# #   subnet = "${aws_subnet.private.0.id}"
# # }

# module "peer_vpcs" {
#   source = "./modules/vpc/peering_connection"

#   peer_vpc_id = "${aws_vpc.management.id}"
#   vpc_id      = "${aws_vpc.application.id}"
# }

# resource "aws_route_table" "management_vpc" {
#   vpc_id = "${aws_vpc.management.id}"

#   route {
#     cidr_block = "${aws_vpc.application.cidr_block}"
#     vpc_peering_connection_id = "${module.peer_vpcs.id}"
#   }

#   tags = {
#     Name = "${var.project_name} vpc peering route table"
#   }
# }
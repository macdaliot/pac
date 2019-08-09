# Determine most recent ECS optimized AMI

data "aws_ami" "ecs_ami" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-*-amazon-ecs-optimized"]
  }
}

resource "aws_iam_instance_profile" "ecsInstanceProfile" {
  name = "ecsInstanceProfile-{{ .env }}-${var.project_name}"
  role = aws_iam_role.ecsInstanceRole_{{ .env }}.name

  # Tags not supported
}

# Generate user_data from template file
data "template_file" "user_data" {
  template = file("${path.module}/user-data.sh")

  vars = {
    ecs_cluster_name = aws_ecs_cluster.main.name
  }
}

# Create Launch Configuration
resource "aws_launch_configuration" "as_conf" {
  image_id                    = data.aws_ami.ecs_ami.id
  # changed from t3.large to t3.xlarge so there was enough memory to autoscale
  instance_type               = "t3.xlarge"
  security_groups             = [aws_security_group.ecs_tasks.id]
  iam_instance_profile        = aws_iam_instance_profile.ecsInstanceProfile.id
  associate_public_ip_address = true # need for pulling from ECR registry

  root_block_device {
    volume_size = "60"
  }

  user_data = data.template_file.user_data.rendered

  lifecycle {
    create_before_destroy = true
  }

  # Tags not supported
}

# Create Auto Scaling Group
resource "aws_autoscaling_group" "asg" {
  name = "asg-${aws_launch_configuration.as_conf.name}"

  //availability_zones        = "${var.aws_zones}"
  vpc_zone_identifier       = aws_subnet.private.*.id
  min_size                  = var.app_count
  max_size                  = var.app_count
  desired_capacity          = var.app_count
  launch_configuration      = aws_launch_configuration.as_conf.id
  health_check_type         = "EC2"
  health_check_grace_period = "120"
  default_cooldown          = "30"

  lifecycle {
    create_before_destroy = true
  }

  # Tags not supported
}

resource "aws_lb_target_group" "jenkins" {
  name        = "${var.project_name}-ecs-ec2-jenkins"
  port        = "8080"
  protocol    = "HTTP"
  vpc_id      = aws_vpc.management_vpc.id
  target_type = "instance"

  health_check {
    interval            = 300
    path                = "/login"
    port                = 8080
    timeout             = 60
    healthy_threshold   = 3
    unhealthy_threshold = 3
    matcher             = 200
  }

  # The new ARN and resource ID format must be enabled to add tags to the service
  # 
  # tags = {
  #   pac-project-name = var.project_name
  #   environment      = "management"
  # }
}

output "target_group_jenkins_arn" {
  value = aws_lb_target_group.jenkins.arn
}

resource "aws_ecs_service" "jenkins" {
  name                              = "jenkins-ecs-service"
  cluster                           = aws_ecs_cluster.main.arn
  task_definition                   = aws_ecs_task_definition.jenkins.arn
  desired_count                     = var.app_count
  launch_type                       = "EC2"
  health_check_grace_period_seconds = 30

  #   network_configuration {
  #     security_groups  = ["${aws_security_group.ecs_tasks.id}"]
  #     subnets          = ["${aws_subnet.private.*.id}"]
  #     assign_public_ip = true # need for pulling from ECR registry
  #   }

  load_balancer {
    target_group_arn = aws_lb_target_group.jenkins.id
    container_name   = "pac-jenkins"
    container_port   = "8080"
  }

  depends_on = [aws_lb_listener.https]

  # tags = {
  #   pac-project-name = var.project_name
  #   environment      = "management"
  # }
}

resource "aws_ecs_task_definition" "jenkins" {
  execution_role_arn       = var.execution_role_arn
  task_role_arn            = var.task_role_arn
  family                   = "pac-jenkins"
  network_mode             = "bridge"
  requires_compatibilities = ["EC2"]

  # Don't set memory to use all of the RAM, otherwise the container won't start.
  # For example, instead of 8G which is the maximum amount of RAM on the instance,
  # I set the container to use 7.5G. The operating needs some to do its work.
  memory = 7680

  container_definitions = <<DEFINITION
[
  {
    "secrets": [
      {
        "name": "JWT_ISSUER",
        "valueFrom": "/pac/${var.project_name}/jwt/issuer"
      },
      {
        "name": "JWT_SECRET",
        "valueFrom": "/pac/${var.project_name}/jwt/secret"
      },
      {
        "name": "USERNAME",
        "valueFrom": "/pac/${var.project_name}/jenkins/username"
      },
      {
        "name": "PASSWORD",
        "valueFrom": "/pac/${var.project_name}/jenkins/password"
      },
      {
        "name": "SONAR_SECRET",
        "valueFrom": "/pac/sonar/secret"
      },
      {
        "name": "GITHUB_USERNAME",
        "valueFrom": "/pac/github/username"
      },
      {
        "name": "GITHUB_PASSWORD",
        "valueFrom": "/pac/github/password"
      },
      {
        "name": "AWS_ACCESS_KEY_ID",
        "valueFrom": "/pac/aws/access_id_key"
      },
      {
        "name": "AWS_SECRET_ACCESS_KEY",
        "valueFrom": "/pac/aws/secret_access_key"
      }
    ],
    "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins:tf-0.12.4",
    "name": "pac-jenkins",
    "portMappings": [
      {
        "containerPort": 8080,
        "hostPort": 8080 
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "ecs"
      }
    },
    "entryPoint": ["/runJenkins.sh"]
  }
]
DEFINITION

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

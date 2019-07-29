resource "aws_ecs_task_definition" "sonarqube" {
  execution_role_arn       = var.execution_role_arn
  task_role_arn            = var.task_role_arn
  family                   = "sonarqube"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "4096"
  memory                   = 8192

  container_definitions = <<DEFINITION
[
  {
    "secrets" : [
      {
        "name" : "sonar.jdbc.username",
        "valueFrom" : "/pac/${var.project_name}/sonar/sonar_jdbc_username"
      },
      {
        "name" : "sonar.jdbc.password",
        "valueFrom" : "/pac/${var.project_name}/sonar/sonar_jdbc_password"
      },
      {
        "name": "sonar.jdbc.url",
        "valueFrom": "/pac/${var.project_name}/sonar/sonar_jdbc_url"
      }
    ],
    "cpu": 2048,
    "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/sonarqube",
    "memory": 4096,
    "name": "sonarqube",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 9000
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "ecs"
      }
    }
  },
  {
    "secrets": [
      {
        "name": "POSTGRES_PASSWORD",
        "valueFrom": "/pac/${var.project_name}/postgres/password" 
      }
    ],
    "cpu": 2048,
    "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/pac-sonar-db",
    "memory": 4096,
    "name": "pac-sonar-db",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 5432
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "ecs"
      }
    }
  }
]
DEFINITION

}

resource "aws_ecs_task_definition" "selenium" {
  execution_role_arn = var.execution_role_arn
  task_role_arn = var.task_role_arn
  family = "pac-selenium"
  network_mode = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu = "4096"
  memory = "8192"

  container_definitions = <<DEFINITION
[
  {
    "cpu": 2048,
    "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/pac-selenium-hub",
    "memory": 4096,
    "name": "pac-selenium-hub-${var.project_name}",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 4448
      }
    ],
    "environment": [
      {
        "name": "SE_OPTS",
        "value": "-port 4448"
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "ecs"
      }
    }
  },
  {
    "cpu": 2048,
    "image": "{{ .awsID }}.dkr.ecr.us-east-2.amazonaws.com/pac-selenium-node-chrome",
    "memory": 4096,
    "name": "pac-selenium-node-chrome-${var.project_name}",
    "networkMode": "awsvpc",
    "environment": [
      {
        "name" : "HUB_HOST",
        "value": "localhost"
      },
      {
        "name" : "HUB_PORT",
        "value": "4448"
      },
      {
        "name" : "NODE_MAX_INSTANCES",
        "value": "5"
      },
      {
        "name" : "NODE_MAX_SESSION",
        "value": "5"
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "${aws_cloudwatch_log_group.{{ .projectName }}_log_group.name}",
          "awslogs-region": "${var.region}",
          "awslogs-stream-prefix": "ecs"
      }
    }
  }
]
DEFINITION

}


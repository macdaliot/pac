resource "aws_ecs_task_definition" "jenkins" {
  execution_role_arn       = "${var.execution_role_arn}"
  task_role_arn            = "${var.task_role_arn}"
  family                   = "pac-jenkins"
  network_mode             = "awsvpc"
  requires_compatibilities = ["EC2", "FARGATE"]
  cpu                      = 2048
  memory                   = 16384

  container_definitions = <<DEFINITION
[
  {
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins",
    "name": "pac-jenkins",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 8080
      }
    ],
    "environment": [
      {
        "name": "jwt_issuer",
        "value": "urn:pacAuth"
      },
      {
        "name": "jwt_secret",
        "value": "${random_string.password.0.result}"
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
          "awslogs-stream-prefix": "ecs"
      }
    }
  }
]
DEFINITION
}

resource "aws_ecs_task_definition" "sonarqube" {
  execution_role_arn       = "${var.execution_role_arn}"
  task_role_arn            = "${var.task_role_arn}"
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
        "name" : "sonar.jdbc.password",
        "valueFrom" : "SONAR_JDBC_PASSWORD"
      }
    ],
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/sonarqube",
    "memory": 4096,
    "name": "sonarqube",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 9000
      }
    ],
    "environment": [
      {
        "name": "sonar.jdbc.url",
        "value": "jdbc:postgresql://localhost/sonar"
      },
      {
        "name": "sonar.jdbc.username",
        "value": "sonar"
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
          "awslogs-stream-prefix": "ecs"
      }
    }
  },
  {
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-sonar-db",
    "memory": 4096,
    "name": "pac-sonar-db",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 5432
      }
    ],
    "environment": [
      {
        "name": "POSTGRES_PASSWORD",
        "value": "pyramid"
      }
    ],
    "logConfiguration": { 
      "logDriver": "awslogs",
      "options": { 
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
          "awslogs-stream-prefix": "ecs"
      }
    }
  }
]
DEFINITION
}

resource "aws_ecs_task_definition" "selenium" {
  execution_role_arn       = "${var.execution_role_arn}"
  task_role_arn            = "${var.task_role_arn}"
  family                   = "pac-selenium"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "4096"
  memory                   = "8192"

  container_definitions = <<DEFINITION
[
  {
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-selenium-hub",
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
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
          "awslogs-stream-prefix": "ecs"
      }
    }
  },
  {
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-selenium-node-chrome",
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
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
          "awslogs-stream-prefix": "ecs"
      }
    }
  }
]
DEFINITION
}
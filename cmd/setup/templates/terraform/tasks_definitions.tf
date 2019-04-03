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
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-jenkins:env",
    "name": "pac-jenkins",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 8080
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
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/sonarqube",
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
          "awslogs-group" : "/ecs/${var.project_name}-log-group",
          "awslogs-region": "us-east-2",
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
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/pac-sonar-db",
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
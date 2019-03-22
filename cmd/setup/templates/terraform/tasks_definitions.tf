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
    "name": "pac-${var.project_name}-jenkins",
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
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/sonarqube",
    "memory": 4096,
    "name": "pac-${var.project_name}-sonarqube",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 9000
      }
    ],
    "environment": [
      {
        "name": "sonar.jdbc.password",
        "value": "sonar"
      },
      {
        "name": "sonar.jdbc.url",
        "value": "jdbc:postgresql://localhost/sonar"
      },
      {
        "name": "sonar.jdbc.username",
        "value": "sonar"
      }
    ]
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
    ]
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
  cpu                      = "${var.fargate_cpu}"
  memory                   = "${var.fargate_memory}"

  container_definitions = <<DEFINITION
[
  {
    "cpu": 2048,
    "image": "118104210923.dkr.ecr.us-east-2.amazonaws.com/selenium",
    "memory": 4096,
    "name": "pac-${var.project_name}-selenium",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 4444
      }
    ]
  }
]
DEFINITION
}
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
    "secrets": [
      {
        "name": "sonar.jdbc.password",
        "valueFrom": "arn:aws:ssm:us-east-2:118104210923:parameter/sonar_jdbc_password"
      },
      {
        "name": "sonar.jdbc.url",
        "valueFrom": "arn:aws:ssm:us-east-2:118104210923:parameter/sonar_jdbc_url"
      },
      {
        "name": "sonar.jdbc.username",
        "valueFrom": "arn:aws:ssm:us-east-2:118104210923:parameter/sonar_jdbc_username"
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
    ]
  },
  {
    "secrets": [
      {
        "name": "POSTGRES_PASSWORD",
        "valueFrom": "arn:aws:ssm:us-east-2:118104210923:parameter/postgres_password"
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
    "name": "selenium",
    "networkMode": "awsvpc",
    "portMappings": [
      {
        "containerPort": 4448
      }
    ]
  }
]
DEFINITION
}
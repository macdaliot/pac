# ALB Security group
# This is the group you need to edit if you want to restrict access to your application
resource "aws_security_group" "lb" {
  name        = "${var.project_name}-alb"
  description = "controls access to the ALB"
  vpc_id      = "${aws_vpc.application.id}"
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "80"
    to_port     = "80"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "8080"
    to_port     = "8080"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "9000"
    to_port     = "9000"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "4445"
    to_port     = "4445"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Traffic to the ECS Cluster should only come from the ALB
resource "aws_security_group" "ecs_tasks" {
  name        = "${var.project_name}-ecs-tasks"
  description = "allow inbound access from the ALB only"
  vpc_id      = "${aws_vpc.application.id}"

  ingress {
    protocol    = "tcp"
    from_port   = "8080"
    to_port     = "8080"
    security_groups = ["${aws_security_group.lb.id}"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "9000"
    to_port     = "9000"
    security_groups = ["${aws_security_group.lb.id}"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "4444"
    to_port     = "4444"
    security_groups = ["${aws_security_group.lb.id}"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }
}
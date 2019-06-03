# Traffic to the ECS Cluster should only come from the ALB
resource "aws_security_group" "ecs_tasks" {
  name        = "${var.project_name}-ecs-tasks"
  description = "allow inbound access from the ALB only"
  vpc_id      = "${aws_vpc.management_vpc.id}"

  ingress {
    protocol    = "tcp"
    from_port   = "8080"
    to_port     = "8080"
    security_groups = ["${aws_security_group.management_lb.id}"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "9000"
    to_port     = "9000"
    security_groups = ["${aws_security_group.management_lb.id}"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "4448"
    to_port     = "4448"
    security_groups = ["${aws_security_group.management_lb.id}"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "8500"
    to_port     = "8500"
    security_groups = ["${aws_security_group.management_lb.id}"]
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.project_name}"
  }
}


output "secgroup_ecs_tasks_name" {
    value = "${aws_security_group.ecs_tasks.name}"
}
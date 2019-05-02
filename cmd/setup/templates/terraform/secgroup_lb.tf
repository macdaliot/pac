# ALB Security group
# This is the group you need to edit if you want to restrict access to your application
resource "aws_security_group" "lb" {
  name        = "${var.project_name}-alb"
  description = "controls access to the ALB"
  vpc_id      = "${aws_vpc.application_vpc.id}"
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "80"
    to_port     = "80"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "${var.project_name}"
  }
}

# resource "aws_security_group" "lb" {
#   name        = "${var.project_name}-alb"
#   description = "controls access to the ALB"
#   vpc_id      = "${aws_vpc.application_vpc.id}"
#   revoke_rules_on_delete = true

#   ingress {
#     protocol    = "tcp"
#     from_port   = "443"
#     to_port     = "443"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }

#   tags {
#     Name = "${var.project_name}"
#   }
# }

output "secgroup_lb_name" {
    value = "${aws_security_group.lb.name}"
}
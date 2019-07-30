# ALB Security group
# This is the group you need to edit if you want to restrict access to your application
resource "aws_security_group" "management_lb" {
  name                   = "${var.project_name}-management-alb"
  description            = "controls access to the ALB"
  vpc_id                 = aws_vpc.management_vpc.id
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "443"
    to_port     = "443"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

output "secgroup_management_lb_name" {
  value = aws_security_group.management_lb.name
}


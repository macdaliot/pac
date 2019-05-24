resource "aws_lb" "management" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.private.*.id}"]
  security_groups = ["${aws_security_group.management_lb.id}"]
}

output "alb_management_name" {
    value = "${aws_lb.management.name}"
}

output "alb_management_arn" {
    value = "${aws_lb.management.arn}"
}
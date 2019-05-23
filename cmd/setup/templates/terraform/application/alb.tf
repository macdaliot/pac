resource "aws_lb" "application" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.lb.id}"]
}

output "alb_application_name" {
  value = "${aws_lb.application.name}"
}

output "alb_application_arn" {
  value = "${aws_lb.application.arn}"
}
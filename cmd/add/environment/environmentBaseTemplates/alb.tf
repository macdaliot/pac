resource "aws_lb" "application" {
  name            = "${var.project_name}-${var.environment_name}-lb"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.application_lb.id}"]
}

output "alb_application_name" {
  value = "${aws_lb.application.name}"
}

output "alb_application_arn" {
  value = "${aws_lb.application.arn}"
}

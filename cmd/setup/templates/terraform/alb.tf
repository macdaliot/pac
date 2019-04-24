resource "aws_lb" "main" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.lb.id}"]
}

output "alb_name" {
    value = "${aws_lb.main.name}"
}

output "alb_arn" {
    value = "${aws_lb.main.arn}"
}
#
# https://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group_attachment" "target_group_attachment" {
  target_group_arn = "${var.alb_target_group_arn}"
  target_id        = "${var.lambda_arn}"
  #depends_on       = ["aws_lambda_permission.with_lb"]
}
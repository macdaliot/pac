#
# https://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html
#
# Provides the ability to register instances and containers with an Application Load Balancer (ALB) 
# or Network Load Balancer (NLB) target group
#
resource "aws_lambda_permission" "with_lb" {
  statement_id  = "AllowExecutionFromlb"
  action        = "lambda:InvokeFunction"
  function_name = "pac-${var.project_name}-i-${var.resource_name}"
  principal     = "elasticloadbalancing.amazonaws.com"
  #source_arn    = "${aws_alb_target_group.pac_lambda_target_group.arn}"
  source_arn    = "${var.source_arn}"
}

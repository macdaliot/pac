resource "aws_flow_log" "vpc_log" {
  iam_role_arn    = var.iam_role_arn
  log_destination = var.log_group_arn
  traffic_type    = var.traffic_type
  vpc_id          = var.vpc_id
}
resource "aws_flow_log" "fl" {
  iam_role_arn    = var.iam_role_arn
  log_destination = var.log_destination
  traffic_type    = var.traffic_type
  vpc_id          = var.vpc_id
}
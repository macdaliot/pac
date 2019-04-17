# Set up cloudwatch group and log stream and retain logs for 1 day
resource "aws_cloudwatch_log_group" "pac_log_group" {
  name              = "/ecs/${var.project_name}-log-group"
  retention_in_days = 1

  tags {
    Name = "${var.project_name}-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "pac_log_stream" {
  name           = "${var.project_name}-log-stream"
  log_group_name = "${aws_cloudwatch_log_group.pac_log_group.name}"
}
# Set up cloudwatch group and log stream and retain logs for 1 day

# A log group is a group of log streams that share the same retention, monitoring, and access control settings. 
# You can define log groups and specify which streams to put into each group. 
resource "aws_cloudwatch_log_group" "yoshi_log_group" {
  name              = "/ecs/${var.project_name}-log-group"
  retention_in_days = "${var.cwl_retention}"

  tags {
    Name = "${var.project_name}-log-group"
  }
}

# A log stream is a sequence of log events that share the same source.
# Each separate source of logs into CloudWatch Logs makes up a separate log stream.
resource "aws_cloudwatch_log_stream" "yoshi_log_stream" {
  name           = "${var.project_name}-log-stream"
  log_group_name = "${aws_cloudwatch_log_group.yoshi_log_group.name}"
}

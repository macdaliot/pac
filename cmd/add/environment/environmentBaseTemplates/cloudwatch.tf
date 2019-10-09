# Set up cloudwatch group and log stream and retain logs for 1 day

# A log group is a group of log streams that share the same retention, monitoring, and access control settings. 
# You can define log groups and specify which streams to put into each group. 
resource "aws_cloudwatch_log_group" "practiceone_log_group" {
  name              = "/ecs/${var.project_name}-log-group/${var.environment_name}"
  retention_in_days = var.cwl_retention
  tags = {
    Name = "${var.project_name}-${var.environment_abbr}-log-group"
    pac-project-name = var.project_name
    environment      = "${var.environment_name}"
  }
}

output "practiceone_log_group_name" {
  value = aws_cloudwatch_log_group.practiceone_log_group.name
}

output "practiceone_log_group_arn" {
  value = aws_cloudwatch_log_group.practiceone_log_group.arn
}

# A log stream is a sequence of log events that share the same source.
# Each separate source of logs into CloudWatch Logs makes up a separate log stream.
resource "aws_cloudwatch_log_stream" "practiceone_log_stream" {
  name           = "${var.project_name}-${var.environment_abbr}-log-stream"
  log_group_name = aws_cloudwatch_log_group.practiceone_log_group.name

  # Doesn't support tags
}

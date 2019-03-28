# output "aws_lb_listener_api_arn" {
#     value = "${aws_lb_listener.api.id}"
# }

output "aws_lb_target_group_jenkins" {
    value = "${aws_lb_target_group.jenkins.id}"
}

output "elasticsearch_created" {
  value = "${var.create_elasticsearch == "true" ? "Elasticsearch created" : "Elasticsearch not created" }"
}

# output "pac_lambda_execution_role_arn" {
#   value = "${aws_iam_role.pac_lambda_execution_role.arn}"
# }
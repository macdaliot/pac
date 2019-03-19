output "aws_lb_listener_api_arn" {
    value = "aws_lb_listener.api.arn"
}

output "elasticsearch_created" {
  value = "${var.create_elasticsearch == "true" ? "Elasticsearch created" : "Elasticsearch not created" }"
}
# https://metallapan.se/post/Provisioning_Elasticsearch_and_Kibana_with_Terraform/

#
# https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html
#
resource "aws_elasticsearch_domain" "es" {
  domain_name           = "${var.project_name}"
  elasticsearch_version = "${var.version}"

  cluster_config {
    instance_type = "${var.instance_type}"
  }

  vpc_options {
    subnet_ids = [
      "${var.subnet}"
    ]

    security_group_ids = ["${aws_security_group.es.id}"]
  }

  ebs_options {
      ebs_enabled = true
      volume_size = 10
  }

  snapshot_options {
    automated_snapshot_start_hour = "${var.automated_snapshot_start_hour}"
  }

  tags = {
    Domain = "${var.project_name}"
  }
}

#
# https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html
#
# resource "aws_elasticsearch_domain_policy" "main" {
#   domain_name = "${aws_elasticsearch_domain.es.domain_name}"

#   access_policies = <<POLICIES
# {
#     "Version": "2012-10-17",
#     "Statement": [
#         {
#             "Action": "es:*",
#             "Principal": "*",
#             "Effect": "Allow",
#             "Condition": {
#                 "IpAddress": {"aws:SourceIp": "127.0.0.1/32"}
#             },
#             "Resource": "${aws_elasticsearch_domain.es.arn}/*"
#         }
#     ]
# }
# POLICIES
# }

data "aws_region" "current" {}

data "aws_caller_identity" "current" {}

resource "aws_security_group" "es" {
  name        = "elasticsearch-${var.project_name}"
  description = "Managed by Terraform"
  vpc_id      = "${var.vpc_id}"

  ingress {
    from_port = 443
    to_port   = 443
    protocol  = "tcp"

    cidr_blocks = [
      "10.0.0.0/16"
    ]
  }
}

# resource "aws_iam_service_linked_role" "es" {
#   aws_service_name = "es.amazonaws.com"
# }
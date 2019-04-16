# resource "aws_elasticsearch_domain" "es" {
#   count                 = "${var.create_elasticsearch == "true" ? 1 : 0}"
#   domain_name           = "${var.project_name}"
#   elasticsearch_version = "${var.es_version}"

#   cluster_config {
#     instance_type = "${var.es_instance_type}"
#   }

#   vpc_options {
#     subnet_ids = [
#       "${aws_subnet.private.0.id}"
#     ]

#     security_group_ids = ["${aws_security_group.es.id}"]
#   }

#   ebs_options {
#       ebs_enabled = true
#       volume_size = 10
#   }

#   snapshot_options {
#     automated_snapshot_start_hour = "${var.es_automated_snapshot_start_hour}"
#   }

#   tags = {
#     Domain = "${var.project_name}"
#   }
# }

# #
# # https://www.terraform.io/docs/providers/aws/r/elasticsearch_domain_policy.html
# #
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

# resource "aws_security_group" "es" {
#   name        = "elasticsearch-${var.project_name}"
#   description = "Managed by Terraform"
#   vpc_id      = "${aws_vpc.management.id}"

#   ingress {
#     from_port = 443
#     to_port   = 443
#     protocol  = "tcp"

#     cidr_blocks = [
#       "10.0.0.0/16"
#     ]
#   }
# }

# resource "aws_iam_service_linked_role" "es" {
#   aws_service_name = "es.amazonaws.com"
# }
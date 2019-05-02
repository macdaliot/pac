
resource "aws_s3_bucket" "integration" {
  count         = "${var.enable_integration_bucket == "true" ? 1:0}"
  bucket        = "integration.${var.project_name}.${var.hosted_zone}"
  acl           = "public-read"
  force_destroy = true

  versioning {
    enabled = true
  }

  website {
    index_document = "index.html"
    error_document = "index.html"
  }

  tags = {
    Name = "${var.project_name} integration bucket"
  }

  # server_side_encryption_configuration {
  #   rule {
  #     apply_server_side_encryption_by_default {
  #       kms_master_key_id = "${aws_kms_key.project_key.key_id}"
  #       sse_algorithm     = "aws:kms"
  #     }
  #   }
  # }
}

# output "s3_integration_bucket_name" {
#     value = "${aws_s3_bucket.integration.bucket}"
# }

# output "integration_brdn" {
#     value = "${aws_s3_bucket.integration.*.bucket_regional_domain_name}"
# }
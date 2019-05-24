
resource "aws_s3_bucket" "project" {
  bucket        = "project.${var.project_name}.${var.hosted_zone}"
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
    Name = "${var.project_name} project bucket"
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

# output "s3_project_bucket_name" {
#     value = "${aws_s3_bucket.project.bucket}"
# }

# output "project_brdn" {
#     value = "${aws_s3_bucket.project.*.bucket_regional_domain_name}"
# }
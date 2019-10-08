#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket.html
#
resource "aws_s3_bucket" "b" {
  acl           = var.acl
  bucket        = var.project_name
  force_destroy = var.force_destroy
  region        = var.region

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = var.key_id
        sse_algorithm     = "aws:kms"
      }
    }
  }

  tags = {
    Name = "${var.project_name} bucket"
  }

  versioning {
    enabled = var.enable_versioning
  }
}
#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket.html
#
resource "aws_s3_bucket" "bucket" {
  acl           = var.acl
  bucket        = var.bucket_name
  force_destroy = "true"
  region        = var.region
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = var.key_id
        sse_algorithm     = "aws:kms"
      }
    }
  }
}

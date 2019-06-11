resource "aws_s3_bucket" "integration" {
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
}
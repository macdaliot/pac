resource "aws_s3_bucket" "{{ .environmentName }}" {
  bucket        = "${var.environment_name}.${var.project_fqdn}"
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
    Name = "${var.project_name} ${var.environment_name} bucket"
  }
}

output "s3_bucket_regional_domain_name" {
  value = aws_s3_bucket.{{ .environmentName }}.bucket_regional_domain_name
}


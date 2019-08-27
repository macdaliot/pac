resource "aws_s3_bucket" "builds" {
  bucket        = "builds.${var.project_name}.${var.hosted_zone}"
  acl           = "private"
  force_destroy = true

  tags = {
    Name = "artifact bucket"
    pac-project-name = var.project_name
    environment = "management"
  }
}

#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket.html
# S3 bucket to hold Lambda code (zip files)
#
resource "aws_s3_bucket" "lambda" {
  count         = "1"
  bucket        = "lambda.${var.project_name}.${var.hosted_zone}"
  force_destroy = true

  tags = {
    Name             = "${var.project_name} Lambda code bucket"
    pac-project-name = "${var.project_name}"
  }
}

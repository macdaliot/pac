#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket.html
#
resource "aws_s3_bucket" "b" {
  bucket = "${var.project_name}"
  acl    = "${var.acl}"

  tags = {
    Name = "${var.project_name} bucket"
  }
}
#
# https://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html
#
resource "aws_cloudfront_origin_access_identity" "oai" {
  comment = "${var.project_name}"
}
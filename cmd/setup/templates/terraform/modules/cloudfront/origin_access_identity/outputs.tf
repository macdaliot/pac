output "oai_path" {
    value = "${aws_cloudfront_origin_access_identity.oai.cloudfront_access_identity_path}"
}
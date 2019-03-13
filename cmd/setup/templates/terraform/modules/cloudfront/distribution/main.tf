#
# https://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html
#
resource "aws_cloudfront_distribution" "s3_distribution" {
  origin {
    domain_name = "${var.s3_brdn}"
    origin_id   = "${var.env}.${var.project_name}.${var.host_name}"

    s3_origin_config {
      origin_access_identity = "${var.oia_path}"
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  http_version        = "http2"
  comment             = "Distribution ${var.env}.${var.project_name}.${var.host_name}"
  default_root_object = "index.html"

  # logging_config {
  #   include_cookies = false
  #   bucket          = "mylogs.s3.amazonaws.com"
  #   prefix          = "myprefix"
  # }

  aliases = ["${var.env}.${var.project_name}.${var.host_name}"]

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "${var.env}.${var.project_name}.${var.host_name}"

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 300
    max_ttl                = 86400 # ask developers about this
  }
  
  custom_error_response {
    error_caching_min_ttl = 86400
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
  }

  custom_error_response {
    error_caching_min_ttl = 86400
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
  }

  price_class = "${var.price_class}"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  tags = {
    pac-project-name = "${var.project_name}"
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}
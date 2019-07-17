variable "price_class" {
  default = "PriceClass_100"
}

#----------------------------------------------------------------------------------------------------------------------
# CREATE AND VALIDATE SSL CERT FOR CLOUDFRONT VIA AWS CERTIFICATE MANAGER
#----------------------------------------------------------------------------------------------------------------------

#
# http://www.terraform.io/docs/providers/aws/r/acm_certificate.html
#
resource "aws_acm_certificate" "{{ .environmentName }}_cert" {
  domain_name       = "{{ .environmentName }}.${var.project_name}.${var.hosted_zone}"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }

}

resource "aws_route53_record" "{{ .environmentName }}_cert_validation" {
  name    = "${aws_acm_certificate.{{ .environmentName }}_cert.domain_validation_options.0.resource_record_name}"
  type    = "${aws_acm_certificate.{{ .environmentName }}_cert.domain_validation_options.0.resource_record_type}"
  zone_id = "${data.terraform_remote_state.dns.main_zone_id}"
  records = ["${aws_acm_certificate.{{ .environmentName }}_cert.domain_validation_options.0.resource_record_value}"]
  ttl     = 60
}

resource "aws_route53_record" "{{ .environmentName }}_cloudfront" {
  zone_id = "${data.terraform_remote_state.dns.main_zone_id}"
  name    = "{{ .environmentName }}.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_cloudfront_distribution.s3_distribution.domain_name}"
    zone_id                = "${aws_cloudfront_distribution.s3_distribution.hosted_zone_id}"
    evaluate_target_health = false
  }
}

resource "aws_acm_certificate_validation" "{{ .environmentName }}_cert" {
  certificate_arn         = "${aws_acm_certificate.{{ .environmentName }}_cert.arn}"
  validation_record_fqdns = ["${aws_route53_record.{{ .environmentName }}_cert_validation.fqdn}"]
}

resource "aws_lb_listener_certificate" "front_end" {
  listener_arn    = "${data.terraform_remote_state.management.aws_lb_listener_https_arn}"
  certificate_arn = "${aws_acm_certificate_validation.{{ .environmentName }}_cert.certificate_arn}"
}

#----------------------------------------------------------------------------------------------------------------------
# CREATE CLOUDFRONT DISTRUBUTION
#----------------------------------------------------------------------------------------------------------------------

# 
# http://www.terraform.io/docs/providers/aws/r/cloudfront_origin_access_identity.html
#
resource "aws_cloudfront_origin_access_identity" "{{ .environmentName }}_oai" {
  comment = "${var.project_name}"
}

output "{{ .environmentName }}_oai_path" {
  value = "${aws_cloudfront_origin_access_identity.{{ .environmentName }}_oai.cloudfront_access_identity_path}"
}

#
# http://www.terraform.io/docs/providers/aws/r/cloudfront_distribution.html
#
resource "aws_cloudfront_distribution" "s3_distribution" {
  origin {
    domain_name = "${aws_lb.application.dns_name}"
    origin_id   = "${aws_lb.application.dns_name}"

    custom_origin_config {
      http_port  = 80
      https_port = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols = ["TLSv1"]
    }
  }

  origin {
    domain_name = "{{ .environmentName }}.${var.project_name}.${var.hosted_zone}.s3.amazonaws.com"
    origin_id   = "{{ .environmentName }}.${var.project_name}.${var.hosted_zone}"

    s3_origin_config {
      origin_access_identity = "${aws_cloudfront_origin_access_identity.{{ .environmentName }}_oai.cloudfront_access_identity_path}"
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  http_version        = "http2"
  comment             = "Distribution {{ .environmentName }}.${var.project_name}.${var.hosted_zone}"
  default_root_object = "index.html"

  # logging_config {
  #   include_cookies = false
  #   bucket          = "mylogs.s3.amazonaws.com"
  #   prefix          = "myprefix"
  # }

  aliases = ["{{ .environmentName }}.${var.project_name}.${var.hosted_zone}"]

  ordered_cache_behavior {
    path_pattern           = "/api/*"
    allowed_methods        = ["GET", "HEAD", "OPTIONS", "PUT", "POST", "PATCH", "DELETE"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "${aws_lb.application.dns_name}"
    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 300
    max_ttl                = 86400 # ask developers about this

    forwarded_values {
      headers      = ["Host"]
      query_string = false

      cookies {
        forward = "none"
      }
    }
  }

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "{{ .environmentName }}.${var.project_name}.${var.hosted_zone}"

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
    acm_certificate_arn = "${aws_acm_certificate.{{ .environmentName }}_cert.arn}"
    ssl_support_method = "sni-only"
  }

  depends_on = ["aws_acm_certificate.{{ .environmentName }}_cert"]
}

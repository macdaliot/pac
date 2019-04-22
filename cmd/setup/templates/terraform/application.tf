# IGW for the public subnet
resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.application.id}"
}

# Route the public subnet traffic through the IGW
resource "aws_route" "internet_access" {
  route_table_id         = "${aws_vpc.application.main_route_table_id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.gw.id}"
}

resource "aws_lb" "main" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.lb.id}"]
}

variable "cnames" {
  default = ["api", "jenkins", "selenium", "sonarqube"]
}

#
# https://www.terraform.io/docs/providers/aws/r/route53_record.html
#
#
# These records are created in application VPC instead of the management VPC because they requie a load balancer which
# is created in the application VPC per design requirements to speed up setup
#
resource "aws_route53_record" "record" {
  

  count   = "${length(var.cnames)}"
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "${element(var.cnames,count.index)}"
  type    = "CNAME"
  ttl     = "60"
  records = ["${aws_lb.main.dns_name}"]
}

resource "aws_s3_bucket" "integration" {
  bucket        = "integration.${var.project_name}.${var.hosted_zone}"
  acl           = "public-read"
  force_destroy = true

  versioning {
    enabled = false
  }

  website {
    index_document = "index.html"
    error_document = "index.html"
  }

  tags = {
    Name = "${var.project_name} integration bucket"
  }

  # server_side_encryption_configuration {
  #   rule {
  #     apply_server_side_encryption_by_default {
  #       kms_master_key_id = "${data.aws_kms_key.testaciousness_kms_key.arn}"
  #       sse_algorithm     = "aws:kms"
  #     }
  #   }
  # }
}

resource "aws_s3_bucket" "demo" {
  bucket        = "demo.${var.project_name}.${var.hosted_zone}"
  acl           = "public-read"
  force_destroy = true

  versioning {
    enabled = false
  }

  website {
    index_document = "index.html"
    error_document = "index.html"
  }

  tags = {
    Name = "${var.project_name} demo bucket"
  }

  # server_side_encryption_configuration {
  #   rule {
  #     apply_server_side_encryption_by_default {
  #       kms_master_key_id = "${data.aws_kms_key.testaciousness_kms_key.arn}"
  #       sse_algorithm     = "aws:kms"
  #     }
  #   }
  # }
}

resource "aws_route53_record" "demo" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "demo.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_lb.main.dns_name}"
    zone_id                = "${aws_lb.main.zone_id}"
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "integration" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "integration.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_lb.main.dns_name}"
    zone_id                = "${aws_lb.main.zone_id}"
    evaluate_target_health = true
  }
}


# https://www.terraform.io/docs/providers/aws/r/iam_role.html
# IAM role
#
resource "aws_iam_role" "{{ .projectName }}_lambda_execution_role" {
  name = "${var.project_name}-lambda-execution-role"
  force_detach_policies = true

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "policy_attach" {
  role       = "${var.project_name}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaRole"

  depends_on = ["aws_iam_role.{{ .projectName }}_lambda_execution_role"]
}


resource "aws_iam_policy" "{{ .projectName }}_lambda_dynamodb" {
  name = "pac-${var.project_name}-lambda-dynamodb-policy"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Sid": "",
          "Effect": "Allow",
          "Action": [
              "dynamodb:BatchGetItem",
              "logs:CreateLogStream",
              "dynamodb:BatchWriteItem",
              "dynamodb:PutItem",
              "dynamodb:GetItem",
              "dynamodb:Scan",
              "dynamodb:Query",
              "dynamodb:UpdateItem",
              "logs:PutLogEvents",
              "logs:CreateLogGroup"
          ],
          "Resource": "*"
      }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_policy_attach" {
  role       = "${var.project_name}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "${aws_iam_policy.{{ .projectName }}_lambda_dynamodb.arn}"

  depends_on = ["aws_iam_role.{{ .projectName }}_lambda_execution_role"]
}


# DO NOT DELETE
# disabled until needed due to long create and destroy cycles
#
# module "pac_cloudfront_oai" {
#   source = "./modules/cloudfront/origin_access_identity"

#   project_name = "${var.project_name}"
# }

# module "pac_cloudfront" {
#   source = "./modules/cloudfront/distribution"

#   s3_brdn = "${module.pac_bucket.brdn}"
#   project_name = "${var.project_name}"
#   oia_path = "${module.pac_cloudfront_oai.oai_path}"
#   env = "integration"
#   price_class = "PriceClass_100"
#   host_name = "${var.hosted_zone}"
# }

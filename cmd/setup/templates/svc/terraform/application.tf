
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
}

resource "aws_route53_record" "demo" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "demo.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_alb.main.dns_name}"
    zone_id                = "${aws_alb.main.zone_id}"
    evaluate_target_health = true
  }
}

resource "aws_route53_record" "integration" {
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "integration.${var.project_name}.${var.hosted_zone}"
  type    = "A"

  alias {
    name                   = "${aws_alb.main.dns_name}"
    zone_id                = "${aws_alb.main.zone_id}"
    evaluate_target_health = true
  }
}


# https://www.terraform.io/docs/providers/aws/r/iam_role.html
# IAM role
#
resource "aws_iam_role" "pac_lambda_execution_role" {
  name = "${var.project_name}-lambda-execution-role"

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

  depends_on = ["aws_iam_role.pac_lambda_execution_role"]
}


resource "aws_iam_policy" "pac_lambda_dynamodb" {
  name = "pac-lambda-dynamodb-policy"
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
  policy_arn = "${aws_iam_policy.pac_lambda_dynamodb.arn}"

  depends_on = ["aws_iam_role.pac_lambda_execution_role"]
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
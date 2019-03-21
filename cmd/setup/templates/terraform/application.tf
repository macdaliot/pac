
resource "aws_vpc" "application" {
  cidr_block = "10.1.0.0/16"

  tags {
    name = "${var.project_name}-application-vpc"
  }
}

# Create var.az_count public subnets, each in a different AZ
resource "aws_subnet" "public" {
  count                   = "${var.az_count}"
  cidr_block              = "${cidrsubnet(aws_vpc.application.cidr_block, 8, var.az_count + count.index)}"
  availability_zone       = "${data.aws_availability_zones.available.names[count.index]}"
  vpc_id                  = "${aws_vpc.application.id}"
  map_public_ip_on_launch = false

  tags {
    name = "${var.project_name}-public-${count.index}"
  }
}

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

resource "aws_alb" "main" {
  name            = "${var.project_name}-load-balancer"
  subnets         = ["${aws_subnet.public.*.id}"]
  security_groups = ["${aws_security_group.lb.id}"]
}

variable "cnames" {
  default = ["api", "jenkins", "selenium", "sonar"]
}

#
# https://www.terraform.io/docs/providers/aws/r/route53_record.html
#
resource "aws_route53_record" "record" {
  

  count   = "${length(var.cnames)}"
  zone_id = "${aws_route53_zone.main.zone_id}"
  name    = "${element(var.cnames,count.index)}"
  type    = "CNAME"
  ttl     = "60"
  records = ["${aws_alb.main.dns_name}"]
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

resource "aws_route_table" "application_vpc" {
  vpc_id = "${aws_vpc.application.id}"

  route {
    cidr_block = "${aws_vpc.management.cidr_block}"
    vpc_peering_connection_id = "${module.peer_vpcs.id}"
  }

  tags = {
    Name = "${var.project_name} vpc peering route table"
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
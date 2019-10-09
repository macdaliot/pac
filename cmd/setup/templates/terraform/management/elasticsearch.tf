#
# http://www.terraform.io/docs/providers/aws/r/elasticsearch_domain.html
#
resource "aws_elasticsearch_domain" "es" {
  domain_name           = "${var.project_name}-management"
  elasticsearch_version = var.es_version

  cluster_config {
    instance_type = var.es_instance_type
  }

  vpc_options {
    subnet_ids = [
      aws_subnet.private[0].id,
    ]
    security_group_ids = [aws_security_group.es.id]
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 10
  }

  snapshot_options {
    automated_snapshot_start_hour = var.es_automated_snapshot_start_hour
  }

  tags = {
    Domain = "${var.project_name} Elasticsearch Domain"
    pac-project-name = var.project_name
    environment      = "management"
  }
}

output "elasticsearch_endpoint" {
  value = aws_elasticsearch_domain.es.endpoint
}

resource "aws_elasticsearch_domain_policy" "vpc_based" {
  domain_name = aws_elasticsearch_domain.es.domain_name

  access_policies = <<POLICIES
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "es:*",
      "Resource": [
          "${aws_elasticsearch_domain.es.arn}/*"
      ]
    }
  ]
}
POLICIES

}

resource "aws_security_group" "es" {
  name = "elasticsearch-${var.project_name}"
  description = "Managed by Terraform"
  vpc_id = aws_vpc.management_vpc.id

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = [
      aws_vpc.management_vpc.cidr_block,
    ]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

#----------------------------------------------------------------------------------------------------------------------
# SEND CLOUDWATCH LOGS TO ELASTICSEARCH
#----------------------------------------------------------------------------------------------------------------------

variable "es_endpoint" {
  type = string
  default = "elasticsearch.endpoint.es.amazonaws.com"
}

variable "cwl_endpoint" {
  type = string
  default = "logs.us-east-2.amazonaws.com"
}

# CloudWatch Logs uses Lambda to deliver log data to Amazon ES. 
# You must specify an IAM role that grants Lambda permission to make calls to Amazon ES. 
# You can choose an existing role or create an IAM role that automatically has the required permissions. 
# To deliver log data to another account, you must specify the Elasticsearch Domain ARN and 
# Elasticsearch Endpoint of other account and ensure permissions are granted to be able to publish to that ARN.
#
resource "aws_iam_role" "lambda_elasticsearch_execution_role" {
  name = "${var.project_name}_management_lambda_es_xrole"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

# When using a non-private Elasticsearch cluster, the Lambda IAM Execution Role needs to have the following permissions 
# policy: AWSLambdaVPCAccessExecutionRole. So we are attaching it here.
resource "aws_iam_role_policy_attachment" "AWSLambdaVPCAccessExecutionRole" {
role  = aws_iam_role.lambda_elasticsearch_execution_role.name

# AWS Mananged role, safe to hard code
policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

# Creates an inline policy on the lambda_elasticsearch_execution_role that allows us to stream logs to Elasticsearch.
resource "aws_iam_role_policy" "lambda_elasticsearch_execution_policy" {
name   = "${var.project_name}_lambda_elasticsearch_execution_policy"
role   = aws_iam_role.lambda_elasticsearch_execution_role.id
policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "es:ESHttpPost",
      "Resource": "${aws_elasticsearch_domain.es.arn}"
    }
  ]
}
EOF

}

resource "aws_lambda_function" "cwl_stream_lambda" {
  filename      = "cwl2eslambda.zip"
  function_name = "LogsToElasticsearch-${var.project_name}-management"
  role          = aws_iam_role.lambda_elasticsearch_execution_role.arn
  handler       = "exports.handler"

  #source_code_hash = "${base64sha256(file("cwl2eslambda.zip"))}"
  runtime = "nodejs8.10"

  environment {
    variables = {
      es_endpoint = aws_elasticsearch_domain.es.endpoint
    }
  }

  vpc_config {
    subnet_ids = [aws_subnet.private[0].id, aws_subnet.private[1].id]
    security_group_ids = [aws_vpc.management_vpc.default_security_group_id]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_lambda_permission" "cloudwatch_allow" {
  statement_id = "cloudwatch_allow"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cwl_stream_lambda.arn
  principal = var.cwl_endpoint
  source_arn = aws_cloudwatch_log_group.{{.projectName}}_log_group.arn
}

resource "aws_cloudwatch_log_subscription_filter" "cloudwatch_logs_to_es" {
  depends_on = [aws_lambda_permission.cloudwatch_allow]
  name = "cloudwatch_logs_to_elasticsearch-management"
  log_group_name = aws_cloudwatch_log_group.{{.projectName}}_log_group.name
  filter_pattern = ""
  destination_arn = aws_lambda_function.cwl_stream_lambda.arn
}

#----------------------------------------------------------------------------------------------------------------------
# DYNAMOBDB ACTIVITY TO ELASTICSEARCH
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_policy" "{{.projectName}}_dev_dynamodb_elasticsearch" {
  name = "${var.project_name}-management-dynamodb-es-policy"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Sid": "",
          "Effect": "Allow",
          "Action": [
              "dynamodb:*"
          ],
          "Resource": "*"
      }
  ]
}
EOF

}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_dev_lambda_elasticsearch_attach_policy" {
  role       = aws_iam_role.lambda_elasticsearch_execution_role.name
  policy_arn = aws_iam_policy.{{.projectName}}_dev_dynamodb_elasticsearch.arn
}

#----------------------------------------------------------------------------------------------------------------------
# JUMPBOX TO ALLOW ACCESS TO KIBANA
#
# https://www.jeremydaly.com/access-aws-vpc-based-elasticsearch-cluster-locally/
# https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html
# https://forums.aws.amazon.com/thread.jspa?threadID=279437
#----------------------------------------------------------------------------------------------------------------------

data "aws_ami" "amzn" {
  most_recent = true

  filter {
    name   = "name"
    values = [var.jumpbox_ami]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["137112412989"] # Amazon
}

resource "aws_security_group" "es_jumpbox" {
  name                   = "${var.project_name}-es-jumpbox"
  description            = "controls access to Kibana"
  vpc_id                 = aws_vpc.management_vpc.id
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "22"
    to_port     = "22"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "443"
    to_port     = "443"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_instance" "jumpbox" {
  ami                         = data.aws_ami.amzn.id
  associate_public_ip_address = true
  instance_type               = "t2.micro"

  # referring to the key pair to be used to SSH into box
  key_name               = "${var.project_name}-worker"
  subnet_id              = aws_subnet.private[0].id
  vpc_security_group_ids = [aws_security_group.es_jumpbox.id]

  tags = {
    Name             = "${var.project_name} Management Jumpbox"
    pac-project-name = var.project_name
    environment      = "management"
  }
}

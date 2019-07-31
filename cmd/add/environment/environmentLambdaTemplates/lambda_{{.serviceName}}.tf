#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket_object.html
# Upload the zip file full of Lambda code to the Lambda S3 bucket
#
resource "aws_s3_bucket_object" "lambda_{{ .serviceName }}_code" {
  bucket     = "${var.environment_name}.${var.project_fqdn}"
  key        = "{{ .serviceName }}.zip"
  source     = "${path.cwd}/../../services/{{ .serviceName }}/function.zip"
  depends_on = [aws_s3_bucket.{{ .environmentName }}]

  tags = {
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda_{{ .serviceName }}" {
  s3_bucket        = "${var.environment_name}.${var.project_fqdn}"
  s3_key           = "{{ .serviceName }}.zip"
  function_name    = "pac-${var.project_name}-${var.environment_abbr}-{{ .serviceName }}"
  role             = data.terraform_remote_state.management.outputs.{{ .projectName }}_lambda_execution_role.arn
  handler          = "lambda.handler"
  source_code_hash = filebase64sha256("${path.cwd}/../../services/{{ .serviceName }}/function.zip")
  runtime          = "nodejs8.10"
  depends_on       = [aws_s3_bucket_object.lambda_{{ .serviceName }}_code]

  environment {
    variables = {
      JWT_ISSUER = data.terraform_remote_state.management.outputs.jwt_issuer.value
      JWT_SECRET = data.terraform_remote_state.management.outputs.jwt_secret.value
      ENV_ABBR   = var.environment_abbr
    }
  }

  tags = {
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

#
# Create target group. Only one lambda can be registered per target group
# http://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
# http://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group" "{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg" {
  name        = "pac-${var.project_name}-${var.environment_abbr}-{{ .serviceName }}"
  port        = "80"
  protocol    = "http"
  vpc_id      = aws_vpc.application_vpc.id
  target_type = "lambda"

  tags = {
    Name = "pac-${var.project_name}-${var.environment_name}-dollhair"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html
#
# Provides the ability to register instances and containers with an Application Load Balancer (ALB) 
# or Network Load Balancer (NLB) target group
#
resource "aws_lambda_permission" "{{ .projectName }}_{{ .serviceName }}_with_lb" {
  statement_id  = "AllowExecutionFromlb"
  action        = "lambda:InvokeFunction"
  function_name = "pac-${var.project_name}-${var.environment_abbr}-{{ .serviceName }}"
  principal     = "elasticloadbalancing.amazonaws.com"

  #source_arn    = "${aws_alb_target_group.pac_lambda_target_group.arn}"
  source_arn = aws_alb_target_group.{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg.id
}

# register with load balancer target group 
resource "aws_alb_target_group_attachment" "{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg_attachment" {
  target_group_arn = aws_alb_target_group.{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg.id
  target_id        = aws_lambda_function.lambda_{{ .serviceName }}.arn
  depends_on       = [aws_lambda_permission.{{ .projectName }}_{{ .serviceName }}_with_lb]
}

#
# http://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html
# alb listener rule
#
resource "aws_lb_listener_rule" "{{ .projectName }}_{{ .serviceName }}_rule_100" {
  listener_arn = aws_lb_listener.{{ .environmentName }}.arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{ .serviceName }}"]
  }
}

resource "aws_lb_listener_rule" "{{ .projectName }}_{{ .serviceName }}_rule_200" {
  listener_arn = aws_lb_listener.{{ .environmentName }}.arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.{{ .projectName }}_{{ .environmentAbbr }}_{{ .serviceName }}_tg.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{ .serviceName }}/*"]
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/dynamodb_table.html
# dynamo_db
#
resource "aws_dynamodb_table" "{{ .projectName }}_dynamodb_table_{{ .serviceName }}" {
  name             = "pac-${var.project_name}-${var.environment_abbr}-{{ .serviceName }}"
  billing_mode     = "PROVISIONED"
  read_capacity    = 5
  write_capacity   = 5
  hash_key         = "id"
  stream_view_type = "NEW_AND_OLD_IMAGES"
  stream_enabled   = true

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name             = "${var.project_name}-dynamodb-table-dollhair"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

resource "aws_lambda_event_source_mapping" "{{ .projectName }}_dynamodb_table_{{ .serviceName }}_source_map" {
  count             = var.enable_elasticsearch == "true" ? 1 : 0
  event_source_arn  = aws_dynamodb_table.{{ .projectName }}_dynamodb_table_{{ .serviceName }}.stream_arn
  function_name     = "DynamoDBToElasticsearch-${var.project_name}"
  starting_position = "LATEST"
}


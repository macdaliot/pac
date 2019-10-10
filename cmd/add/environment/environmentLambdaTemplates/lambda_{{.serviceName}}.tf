#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket_object.html
# Upload the zip file full of Lambda code to the Lambda S3 bucket
#
resource "aws_s3_bucket_object" "lambda_{{.serviceName}}_code" {
  bucket     = aws_s3_bucket.{{.environmentName}}.id
  key        = "{{ .serviceName }}.zip"
  source     = "${path.cwd}/../../services/{{.serviceName}}/function.zip"
  depends_on = [aws_s3_bucket.{{.environmentName}}]
  etag = "${filemd5("${path.cwd}/../../services/{{.serviceName}}/function.zip")}"
  tags = {
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

#
# http://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda_{{.serviceName}}" {
  s3_bucket        = aws_s3_bucket.{{.environmentName}}.id
  s3_key           = "{{.serviceName}}.zip"
  function_name    = "pac-${var.project_name}-${var.environment_abbr}-{{.serviceName}}"
  role             = aws_iam_role.lambda_elasticsearch_execution_role.arn
  handler          = "lambda.handler"
  source_code_hash = filebase64sha256("${path.cwd}/../../services/{{.serviceName}}/function.zip")
  depends_on       = [aws_s3_bucket_object.lambda_{{ .serviceName }}_code]
  memory_size      = 256
  timeout          = 15
  runtime          = "nodejs10.x"

  environment {
    variables = {
      PROJECTNAME = var.project_name
      JWT_ISSUER  = data.terraform_remote_state.management.outputs.jwt_issuer.value
      JWT_SECRET  = data.terraform_remote_state.management.outputs.jwt_secret.value
      ENV_ABBR    = var.environment_abbr
      ES_ENDPOINT = "https://${aws_elasticsearch_domain.es.endpoint}"
    }
  }

  vpc_config {
    subnet_ids = [aws_subnet.public[0].id, aws_subnet.public[1].id]
    security_group_ids = [aws_vpc.application_vpc.default_security_group_id]
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
resource "aws_alb_target_group" "{{.projectName}}_{{.environmentAbbr}}_{{.serviceName}}_tg" {
  name        = "pac-${var.project_name}-${var.environment_abbr}-{{.serviceName}}"
  port        = "80"
  protocol    = "http"
  vpc_id      = aws_vpc.application_vpc.id
  target_type = "lambda"

  tags = {
    Name = "pac-${var.project_name}-${var.environment_name}-{{.serviceName}}"
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
  function_name = aws_lambda_function.lambda_{{.serviceName}}.function_name
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
resource "aws_lb_listener_rule" "{{.projectName}}_{{.serviceName}}_rule_100" {
  listener_arn = aws_lb_listener.{{.environmentName}}.arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.{{.projectName}}_{{.environmentAbbr}}_{{.serviceName}}_tg.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{.serviceName}}"]
  }
}

resource "aws_lb_listener_rule" "{{.projectName}}_{{.serviceName}}_rule_200" {
  listener_arn = aws_lb_listener.{{.environmentName}}.arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.{{.projectName}}_{{.environmentAbbr}}_{{.serviceName}}_tg.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{ .serviceName }}/*"]
  }
}


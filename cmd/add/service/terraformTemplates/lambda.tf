#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket_object.html
# Upload the zip file full of Lambda code to the Lambda S3 bucket
#
resource "aws_s3_bucket_object" "lambda_{{ .serviceName }}_code" {
  bucket = "lambda.${var.project_name}.${var.hosted_zone}"
  key    = "{{ .serviceName }}.zip"
  source = "${path.cwd}/../{{ .serviceName }}/function.zip"
}

#
# http://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda_{{ .serviceName }}" {
  s3_bucket        = "lambda.${var.project_name}.${var.hosted_zone}"
  s3_key           = "{{ .serviceName }}.zip"
  function_name    = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  role             = "${data.terraform_remote_state.app.{{ .projectName }}_lambda_execution_role_arn}"
  handler          = "lambda.handler"
  # source_code_hash = "${base64sha256(file(var.lambda_function_payload))}"
  runtime          = "nodejs8.10"
  depends_on       = ["aws_s3_bucket_object.lambda_{{ .serviceName }}_code"]

  environment {
    variables = {
      JWT_ISSUER = "${data.terraform_remote_state.app.jwt_issuer}"
      JWT_SECRET = "${data.terraform_remote_state.app.jwt_secret}"
    }
  }

  tags {
		pac-project-name = "{{ .projectName }}"
  }
}

#
# Create target group. Only one lambda can be registered per target group
# http://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
# http://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group" "{{ .projectName }}_{{ .serviceName }}_target_group" {
  name        = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  port        = "80"
  protocol    = "http"
  vpc_id      = "${data.terraform_remote_state.app.application_vpc_id}"
  target_type = "lambda"
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
  function_name = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  principal     = "elasticloadbalancing.amazonaws.com"
  #source_arn    = "${aws_alb_target_group.pac_lambda_target_group.arn}"
  source_arn    = "${aws_alb_target_group.{{ .projectName }}_{{ .serviceName }}_target_group.id}"
}

# register with load balancer target group 
resource "aws_alb_target_group_attachment" "{{ .projectName }}_{{ .serviceName }}_target_group_attachment" {
  target_group_arn = "${aws_alb_target_group.{{ .projectName }}_{{ .serviceName }}_target_group.id}"
  target_id        = "${aws_lambda_function.lambda_{{ .serviceName }}.arn}"
  depends_on       = ["aws_lambda_permission.{{ .projectName }}_{{ .serviceName }}_with_lb"]
}

#
# http://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html
# alb listener rule
#
resource "aws_lb_listener_rule" "{{ .projectName }}_{{ .serviceName }}_rule_100" {
  listener_arn = "${data.terraform_remote_state.app.aws_lb_listener_api_arn}"

  action {
    type = "forward"
    target_group_arn = "${aws_alb_target_group.{{ .projectName }}_{{ .serviceName }}_target_group.arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{ .serviceName }}"]
  }
}

resource "aws_lb_listener_rule" "{{ .projectName }}_{{ .serviceName }}_rule_200" {
  listener_arn = "${data.terraform_remote_state.app.aws_lb_listener_api_arn}"


  action {
    type = "forward"
    target_group_arn = "${aws_alb_target_group.{{ .projectName }}_{{ .serviceName }}_target_group.arn}"
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
  name           = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "id"
  stream_view_type = "NEW_AND_OLD_IMAGES"
  stream_enabled = true

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name = "{{ .projectName }}-dynamodb-table-{{ .serviceName }}"
		pac-project-name = "{{ .projectName }}"
  }
}

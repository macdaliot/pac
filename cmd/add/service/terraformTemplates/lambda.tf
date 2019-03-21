#
# https://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda_{{ .serviceName }}" {
  filename         = "function.zip"
  function_name    = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  role             = "${data.terraform_remote_state.pac.pac_lambda_execution_role_arn}"
  handler          = "lambda.handler"
  # source_code_hash = "${base64sha256(file(var.lambda_function_payload))}"
  runtime          = "nodejs8.10"

  environment {
    variables = {
      JWT_ISSUER = "urn:pacAuth"
      JWT_SECRET = "4pWQUrx6RkgU6o2TC"
    }
  }

  tags {
      pac-project-name = "{{ .projectName }}-{{ .serviceName }}"
  }
}

# Create target group. Only one lambda can be registered per target group
# https://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
# https://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group" "{{ .projectName }}_{{ .serviceName }}_target_group" {
  name        = "pac-{{ .projectName }}-i-{{ .serviceName }}"
  # port        = "${var.app_port}"
  # protocol    = "${var.app_protocol}"
  # vpc_id      = "${var.vpc_id}"
  target_type = "lambda"
}

#
# https://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html
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
# https://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html
# alb listener rule
#
resource "aws_lb_listener_rule" "{{ .projectName }}_{{ .serviceName }}_rule_100" {
  listener_arn = "${data.terraform_remote_state.pac.aws_lb_listener_api_arn}"

  priority = "100"

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
  listener_arn = "${data.terraform_remote_state.pac.aws_lb_listener_api_arn}"

  priority = "200"

  action {
    type = "forward"
    target_group_arn = "${aws_alb_target_group.{{ .projectName }}_{{ .serviceName }}_target_group.arn}"
  }

  condition {
    field  = "path-pattern"
    values = ["/api/{{ .serviceName }}"]
  }
}

#
# https://www.terraform.io/docs/providers/aws/r/dynamodb_table.html
# dynamo_db
#
resource "aws_dynamodb_table" "{{ .projectName }}_dynamodb_table_{{ .serviceName }}" {
  name           = "pac-{{ .projectName }}-{{ .serviceName }}"
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
    Name        = "{{ .projectName }}-dynamodb-table-{{ .serviceName }}"
  }
}
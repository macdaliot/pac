#
# https://www.terraform.io/docs/providers/aws/r/s3_bucket_object.html
# Upload the zip file full of Lambda code to the Lambda S3 bucket
#
resource "aws_s3_bucket_object" "lambda_auth_code" {
  bucket = "lambda.${var.project_name}.${var.hosted_zone}"
  key    = "auth.zip"
  source = "${path.cwd}/../auth/function.zip"
}

#
# https://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda_auth" {
  s3_bucket        = "lambda.${var.project_name}.${var.hosted_zone}"
  s3_key           = "auth.zip"
  function_name    = "pac-{{ .projectName }}-i-auth"
  role             = "${data.terraform_remote_state.management.{{ .projectName }}_lambda_execution_role_arn}"
  handler          = "lambda.handler"
  # source_code_hash = "${base64sha256(file(var.lambda_function_payload))}"
  runtime          = "nodejs8.10"
  timeout          = 10
  depends_on       = ["aws_s3_bucket_object.lambda_auth_code"]

  environment {
    variables = {
      APP_ROOT        = "http://integration.{{ .projectName }}.pac.pyramidchallenges.com.s3-website.{{ .region }}.amazonaws.com"
      AUTH0_CLIENT_ID = "PJqs70Pr0VRH67aR2TnHf4Sn6DDldiNR"
      AUTH0_DOMAIN    = "pyramidsystems.auth0.com"
      JWT_SECRET      = "${data.terraform_remote_state.management.jwt_secret}"
      SAML_CALLBACK   = "https://api.{{ .projectName }}.pac.pyramidchallenges.com/api/auth/callback"
      SAML_SIGNER     = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tDQpNSUlDK0RDQ0FlQ2dBd0lCQWdJSlBPV29wemhNQVdPSk1BMEdDU3FHU0liM0RRRUJDd1VBTUNNeElUQWZCZ05WDQpCQU1UR0hCNWNtRnRhV1J6ZVhOMFpXMXpMbUYxZEdnd0xtTnZiVEFlRncweE56QXpNVEV4T0RJd01EbGFGdzB6DQpNREV4TVRneE9ESXdNRGxhTUNNeElUQWZCZ05WQkFNVEdIQjVjbUZ0YVdSemVYTjBaVzF6TG1GMWRHZ3dMbU52DQpiVENDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFNa2pFeDkyZkFscmU0S0xYUXgvDQplQ2lUTFZOY2ljMnlUeXRRb2cveWJsRzJvMEx5L0lGdHhaTmpFaEF2ODQ1MGpiWDdFVG1iNkhrUEZWb0FHcDNsDQp6WGdxNTByL3pRSzRmMVp6Qzc1VC9YL01HVlF3TmRGNDZXUEJlaUs4T0xndzVZa2RUa1J6V1pDQm9zNGRMK2NZDQoyWHc0eGpZQ1MxL28wMExPZzVjaDdCTEtodTRSb0dRRnRWRUZBV0VTTytBZ2xCRVFSR2ljbXQ2WWhVTW5RdTV6DQo4bHdvTXRRNFlZdk9vNVpsbUdLZ2hEOEJzTmhidGdLRkV2UVlwbXpmZVQ5amRScjhGOHN6cHkya0xmTU0wRWMzDQo4cEEvNnFialA0ckhHdzE5ZS9FUmhyRlRTa2FFMFMyVFkya1l0V1RFODdCSHI4SWJFaHYvNDl1OGVkbW5ERlBFDQo3aDBDQXdFQUFhTXZNQzB3REFZRFZSMFRCQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVLMEFuUHpXeGhLOTZzeEUrDQoydjhCQjRBc2ZKMHdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBSWRtYkFzVC9tK01YcFh0MzBuMHk4dU5JcCs4DQppTFVIM3duL3RaQUl6Tjk1WDJ1ZlNuRTZSb1dWaGNzOVAyT0NScjY0b1ZzU29McUJhWjVwSkhINmdtSTVlc0lTDQp6dlRFYXZIUDIyb2QyeDUwcEVlZGpzSyt5Rlh1bkZ2b2xYYnIyeFU2bmMwc3FuSW12OWVCMU1vN3FsU0hxQnhIDQpoUkh5NXhCUGVVcTI5U0ZrM0wwd042UUhaNXNTZGRKV2prYmMvVWdFZnpGeURNQUxqZEhHUGd2ZVB6bU44SlVJDQpWVkxYOHBrYWdQYVJ4RWd2dTN6RGtTRjBBMkEvejZPMFZVWWNXWi9YN1Nad0Z4d3ZRNEpoYTR1ejFOcHJrY0s0DQptNUNkY1pSak41WWlUWUV5YmtBWWR4Z0w3M3dZUUQwYlRYVVAySkRscVlYeVFUVVdPWHFaR0N4NXRFST0NCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0NCg=="
    }
  }

  tags {
      pac-project-name = "{{ .projectName }}"
  }
}

#
# Create target group. Only one lambda can be registered per target group
# https://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
# https://www.terraform.io/docs/providers/aws/r/lb_target_group.html
#
resource "aws_alb_target_group" "{{ .projectName }}_auth_target_group" {
  name        = "pac-{{ .projectName }}-i-auth"
  port        = "80"
  protocol    = "http"
  vpc_id      = "${data.terraform_remote_state.management.application_vpc_id}"
  target_type = "lambda"
}

#
# https://www.terraform.io/docs/providers/aws/r/lb_target_group_attachment.html
#
# Provides the ability to register instances and containers with an Application Load Balancer (ALB) 
# or Network Load Balancer (NLB) target group
#
resource "aws_lambda_permission" "{{ .projectName }}_auth_with_lb" {
  statement_id  = "AllowExecutionFromlb"
  action        = "lambda:InvokeFunction"
  function_name = "pac-{{ .projectName }}-i-auth"
  principal     = "elasticloadbalancing.amazonaws.com"
  # source_arn    = "${aws_alb_target_group.pac_lambda_target_group.arn}"
  source_arn    = "${aws_alb_target_group.{{ .projectName }}_auth_target_group.id}"
}

# register with load balancer target group 
resource "aws_alb_target_group_attachment" "{{ .projectName }}_auth_target_group_attachment" {
  target_group_arn = "${aws_alb_target_group.{{ .projectName }}_auth_target_group.id}"
  target_id        = "${aws_lambda_function.lambda_auth.arn}"
  depends_on       = ["aws_lambda_permission.{{ .projectName }}_auth_with_lb"]
}

#
# https://www.terraform.io/docs/providers/aws/r/lb_listener_rule.html
# alb listener rule
#
resource "aws_lb_listener_rule" "{{ .projectName }}_auth_rule_100" {
  listener_arn = "${data.terraform_remote_state.management.aws_lb_listener_api_arn}"
  action {
    type = "forward"
    target_group_arn = "${aws_alb_target_group.{{ .projectName }}_auth_target_group.arn}"
  }
  condition {
    field  = "path-pattern"
    values = ["/api/auth"]
  }
}

resource "aws_lb_listener_rule" "{{ .projectName }}_auth_rule_200" {
  listener_arn = "${data.terraform_remote_state.management.aws_lb_listener_api_arn}"
  action {
    type = "forward"
    target_group_arn = "${aws_alb_target_group.{{ .projectName }}_auth_target_group.arn}"
  }
  condition {
    field  = "path-pattern"
    values = ["/api/auth/*"]
  }
}

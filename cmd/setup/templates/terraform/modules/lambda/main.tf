# LAMBDA_ARN=$(aws lambda create-function 
# --function-name "$FULL_SERVICE_NAME" 
# --runtime nodejs8.10 
# --role arn:aws:iam::118104210923:role/service-role/god 
# --handler lambda.handler 
# --zip-file fileb://function.zip 
# --region us-east-2 | jq '.FunctionArn')

#
# https://www.terraform.io/docs/providers/aws/r/lambda_function.html
# lambda function
#
resource "aws_lambda_function" "lambda" {
  filename         = "${var.lambda_function_payload}"
  function_name    = "pac-${var.project_name}-i-${var.resource_name}"
  role             = "${var.role}"
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
      pac-project-name = "${var.project_name}-${var.resource_name}"
  }
}
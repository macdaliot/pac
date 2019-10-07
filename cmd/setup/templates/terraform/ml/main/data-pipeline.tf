
#----------------------------------------------------------------------------------------------------------------------
# ML {{ .projectName }} Ingestion Lambda Function Permissions
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "ml_data_ingestion_lambda_xrole" {
  name = "${var.project_name}_lambda_s3_xrole"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      }
    }
  ]
}
EOF
  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
  }
}

resource "aws_iam_policy" "{{ .projectName }}_s3_allow_lambda_access" {
  name = "${var.project_name}-ml-object-upload-allow-s3-access"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": [
        "s3:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_policy" "{{ .projectName }}_dynamodb_allow_lambda_access" {
  name = "${var.project_name}-ml-object-upload-allow-dynamo-access"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": [
        "dynamodb:*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "AWSLambdaVPCAccessExecutionRole" {
  role       = aws_iam_role.ml_data_ingestion_lambda_xrole.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_s3_allow_lambda_access" {
  role       = aws_iam_role.ml_data_ingestion_lambda_xrole.name
  policy_arn = aws_iam_policy.{{ .projectName }}_s3_allow_lambda_access.arn
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_dynamo_allow_lambda_access" {
  role       = aws_iam_role.ml_data_ingestion_lambda_xrole.name
  policy_arn = aws_iam_policy.{{ .projectName }}_dynamodb_allow_lambda_access.arn
}

#----------------------------------------------------------------------------------------------------------------------
# Outputs
#----------------------------------------------------------------------------------------------------------------------

output "ml_data_ingestion_lambda_xrole" {
  value = aws_iam_role.ml_data_ingestion_lambda_xrole
}

resource "aws_lambda_permission" "allow_celebrity_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.development.outputs.celebrity_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_character_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket1"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.development.outputs.character_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_role_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket2"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.development.outputs.roles_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_celebrity_bucket1" {
  statement_id  = "AllowExecutionFromS3Bucket3"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.e2etesting.outputs.celebrity_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_character_bucket1" {
  statement_id  = "AllowExecutionFromS3Bucket4"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.e2etesting.outputs.character_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_role_bucket1" {
  statement_id  = "AllowExecutionFromS3Bucket5"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.e2etesting.outputs.roles_upload_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

# resource "aws_lambda_permission" "allow_celebrity_bucket2" {
#   statement_id  = "AllowExecutionFromS3Bucket6"
#   action        = "lambda:InvokeFunction"
#   function_name = data.terraform_remote_state.production.outputs.celebrity_upload_lambda.function_name
#   principal     = "s3.amazonaws.com"
#   source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
# }

# resource "aws_lambda_permission" "allow_character_bucket2" {
#   statement_id  = "AllowExecutionFromS3Bucket7"
#   action        = "lambda:InvokeFunction"
#   function_name = data.terraform_remote_state.production.outputs.character_upload_lambda.function_name
#   principal     = "s3.amazonaws.com"
#   source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
# }

# resource "aws_lambda_permission" "allow_role_bucket2" {
#   statement_id  = "AllowExecutionFromS3Bucket8"
#   action        = "lambda:InvokeFunction"
#   function_name = data.terraform_remote_state.production.outputs.roles_upload_lambda.function_name
#   principal     = "s3.amazonaws.com"
#   source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
# }

resource "aws_lambda_permission" "allow_result_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket9"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.development.outputs.upload_results_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

resource "aws_lambda_permission" "allow_result_bucket1" {
  statement_id  = "AllowExecutionFromS3Bucket10"
  action        = "lambda:InvokeFunction"
  function_name = data.terraform_remote_state.e2etesting.outputs.upload_results_lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
}

# resource "aws_lambda_permission" "allow_result_bucket2" {
#   statement_id  = "AllowExecutionFromS3Bucket11"
#   action        = "lambda:InvokeFunction"
#   function_name = data.terraform_remote_state.production.outputs.upload_results_lambda.function_name
#   principal     = "s3.amazonaws.com"
#   source_arn    = module.codepipeline_1.ml_codepipeline_data_bucket.arn
# }
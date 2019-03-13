
#
# https://www.terraform.io/docs/providers/aws/r/iam_role.html#assume_role_policy
#

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
#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-jenkins ROLE
#----------------------------------------------------------------------------------------------------------------------

# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{ .projectName }}_{{ .env }}_jenkins" {
    name = "{{ .projectName }}-dev-jenkins"

    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

# Creates an inline policy to allow listing and updating route53 records
resource "aws_iam_role_policy" "{{ .projectName }}_{{ .env }}_jenkins_update_route53" {
  name   = "{{ .projectName }}-dev-jenkins-update-route53-policy"
  role   = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.id}"
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "route53:ChangeResourceRecordSets",
            "Resource": "arn:aws:route53:::hostedzone/*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": "route53:ListHostedZones",
            "Resource": "*"
        }
    ]
}
EOF
}

# Creates an inline policy on the lambda_elasticsearch_execution_role that allows us to stream logs to Elasticsearch.
resource "aws_iam_role_policy" "{{ .projectName }}_{{ .env }}_jenkins_ec2_specifics" {
  name   = "{{ .projectName }}-dev-jenkins-ec2-policy"
  role   = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.id}"
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "",
            "Effect": "Allow",
            "Action": [
                "ec2:CreateTags",
                "ec2:RunInstances"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ec2_container_registry_full_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_s3_read_only_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ecs_full_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonECS_FullAccess"
}

#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-task-execution ROLE
#----------------------------------------------------------------------------------------------------------------------
# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{ .projectName }}_task_execution" {
    name = "${var.project_name}-dev-task-execution"

    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ecs_task_execution" {
  role       = "${aws_iam_role.{{ .projectName }}_task_execution.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# Creates an inline policy to allow reading from System Manager Parameter Store
resource "aws_iam_role_policy" "parameter_store" {
  name   = "${var.project_name}-dev-parameter-store"
  role   = "${aws_iam_role.{{ .projectName }}_task_execution.id}"
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ssm:GetParameters",
                "secretsmanager:GetSecretValue",
                "kms:Decrypt"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_task_ec2_container_registry_full_access" {
  role   = "${aws_iam_role.{{ .projectName }}_task_execution.id}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-lambda-execution ROLE
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "{{ .projectName }}_{{ .env }}_lambda_execution_role" {
  name = "${var.project_name}-dev-lambda-execution-role"
  force_detach_policies = true

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

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_attach_aws_lambda_role" {
  role       = "${var.project_name}-dev-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaRole"

  depends_on = ["aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role"]
}

output "{{ .projectName }}_lambda_execution_role_arn" {
    value = "${aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role.arn}"
}

# output "{{ .projectName }}_{{ .env }}_lambda_execution_role_name" {
#     value = "${aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role.name}"
# }

# output "{{ .projectName }}_lambda_exec_policy_attachment_policy_arn" {
#     value ="${aws_iam_role_policy_attachment.{{ .projectName }}_{{ .env }}_attach_aws_lambda_role.id}"
# }


resource "aws_iam_policy" "{{ .projectName }}_{{ .env }}_lambda_dynamodb" {
  name = "${var.project_name}-dev-lambda-dynamodb-policy"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Sid": "",
          "Effect": "Allow",
          "Action": [
              "dynamodb:BatchGetItem",
              "logs:CreateLogStream",
              "dynamodb:BatchWriteItem",
              "dynamodb:PutItem",
              "dynamodb:GetItem",
              "dynamodb:Scan",
              "dynamodb:Query",
              "dynamodb:UpdateItem",
              "logs:PutLogEvents",
              "logs:CreateLogGroup"
          ],
          "Resource": "*"
      }
  ]
}
EOF
}

# output "{{ .projectName }}_{{ .env }}_lambda_dynamodb_policy_arn" {
#   value = "${aws_iam_policy.{{ .projectName }}_{{ .env }}_lambda_dynamodb.arn}"
# }

# output "{{ .projectName }}_{{ .env }}_lambda_dynamodb_policy_name" {
#   value = "${aws_iam_policy.{{ .projectName }}_{{ .env }}_lambda_dynamodb.name}"
# }

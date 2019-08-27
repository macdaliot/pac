#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-jenkins ROLE
#----------------------------------------------------------------------------------------------------------------------

# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{ .projectName }}_{{ .env }}_jenkins" {
    name = "${var.project_name}-{{ .env }}-jenkins"

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

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

# Creates an inline policy to allow listing and updating route53 records
resource "aws_iam_role_policy" "{{ .projectName }}_{{ .env }}_jenkins_update_route53" {
  name   = "${var.project_name}-{{ .env }}-jenkins-update-route53-policy"
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
  # Tags not supported
}

# Creates an inline policy on the lambda_elasticsearch_execution_role that allows us to stream logs to Elasticsearch.
resource "aws_iam_role_policy" "{{ .projectName }}_{{ .env }}_jenkins_ec2_specifics" {
  name   = "${var.project_name}-{{ .env }}-jenkins-ec2-policy"
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

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ec2_container_registry_full_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_s3_read_only_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ecs_full_access" {
  role       = "${aws_iam_role.{{ .projectName }}_{{ .env }}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonECS_FullAccess"

  # Tags not supported
}

#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-task-execution ROLE
#----------------------------------------------------------------------------------------------------------------------
# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{ .projectName }}_task_execution" {
    name = "${var.project_name}-{{ .env }}-task-execution"

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

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_ecs_task_execution" {
  role       = "${aws_iam_role.{{ .projectName }}_task_execution.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"

  # Tags not supported
}

# Creates an inline policy to allow reading from System Manager Parameter Store
resource "aws_iam_role_policy" "parameter_store" {
  name   = "${var.project_name}-{{ .env }}-parameter-store"
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

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_task_ec2_container_registry_full_access" {
  role   = "${aws_iam_role.{{ .projectName }}_task_execution.id}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"

  # Tags not supported
}

#----------------------------------------------------------------------------------------------------------------------
# {{ .projectName }}-{{ .env }}-lambda-execution ROLE
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "{{ .projectName }}_{{ .env }}_lambda_execution_role" {
  name = "${var.project_name}-{{ .env }}-lambda-execution-role"
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

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_attach_aws_lambda_role" {
  role       = "${var.project_name}-{{ .env }}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaRole"

  depends_on = ["aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role"]

  # Tags not supported
}

output "{{ .projectName }}_lambda_execution_role" {
    value = "${aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role}"
}

# output "{{ .projectName }}_lambda_exec_policy_attachment_policy" {
#     value ="${aws_iam_role_policy_attachment.{{ .projectName }}_{{ .env }}_attach_aws_lambda_role}"
# }

resource "aws_iam_policy" "{{ .projectName }}_{{ .env }}_lambda_dynamodb" {
  name = "${var.project_name}-{{ .env }}-lambda-dynamodb-policy"
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
              "dynamodb:DeleteItem",
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

resource "aws_iam_role_policy_attachment" "{{ .projectName }}_{{ .env }}_lambda_dynamodb_policy_attach" {
  role       = "${var.project_name}-{{ .env }}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "${aws_iam_policy.{{ .projectName }}_{{ .env }}_lambda_dynamodb.arn}"

  depends_on = ["aws_iam_role.{{ .projectName }}_{{ .env }}_lambda_execution_role"]

  # Tags not supported
}

# output "{{ .projectName }}_{{ .env }}_lambda_dynamodb_policy" {
#   value = "${aws_iam_policy.{{ .projectName }}_{{ .env }}_lambda_dynamodb}"
# }


#----------------------------------------------------------------------------------------------------------------------
# ECS EC2 ROLES
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "ecsInstanceRole_{{ .env }}" {
  name = "ecsInstanceRole-{{ .env }}-${var.project_name}"

  assume_role_policy = <<EOF
{
 "Version": "2008-10-17",
 "Statement": [
   {
     "Sid": "",
     "Effect": "Allow",
     "Principal": {
       "Service": "ec2.amazonaws.com"
     },
     "Action": "sts:AssumeRole"
   }
 ]
}
EOF

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_iam_role_policy" "ecsInstanceRolePolicy_{{ .env }}" {
  name = "ecsInstanceRolePolicy-{{ .env }}-${var.project_name}"
  role = "${aws_iam_role.ecsInstanceRole_{{ .env }}.id}"

  policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Effect": "Allow",
     "Action": [
       "ecs:CreateCluster",
       "ecs:DeregisterContainerInstance",
       "ecs:DiscoverPollEndpoint",
       "ecs:Poll",
       "ecs:RegisterContainerInstance",
       "ecs:StartTelemetrySession",
       "ecs:Submit*",
       "ecr:GetAuthorizationToken",
       "ecr:BatchCheckLayerAvailability",
       "ecr:GetDownloadUrlForLayer",
       "ecr:BatchGetImage",
       "logs:CreateLogStream",
       "logs:PutLogEvents"
     ],
     "Resource": "*"
   }
 ]
}
EOF

  # Tags not supported
}

# Create ECS IAM Service Role and Policy
resource "aws_iam_role" "ecsServiceRole_{{ .env }}" {
  name = "ecsServiceRole-{{ .env }}-${var.project_name}"

  assume_role_policy = <<EOF
{
 "Version": "2008-10-17",
 "Statement": [
   {
     "Sid": "",
     "Effect": "Allow",
     "Principal": {
       "Service": "ecs.amazonaws.com"
     },
     "Action": "sts:AssumeRole"
   }
 ]
}
EOF

  tags = {
    pac-project-name = var.project_name
    environment      = "management"
  }
}

resource "aws_iam_role_policy" "ecsServiceRolePolicy_{{ .env }}" {
  name = "ecsServiceRolePolicy-{{ .env }}-${var.project_name}"
  role = "${aws_iam_role.ecsServiceRole_{{ .env }}.id}"

  policy = <<EOF
{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Effect": "Allow",
     "Action": [
       "ec2:AuthorizeSecurityGroupIngress",
       "ec2:Describe*",
       "elasticloadbalancing:DeregisterInstancesFromLoadBalancer",
       "elasticloadbalancing:DeregisterTargets",
       "elasticloadbalancing:Describe*",
       "elasticloadbalancing:RegisterInstancesWithLoadBalancer",
       "elasticloadbalancing:RegisterTargets"
     ],
     "Resource": "*"
   }
 ]
}
EOF

  # Tags not supported
}

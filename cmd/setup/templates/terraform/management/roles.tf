#----------------------------------------------------------------------------------------------------------------------
# {{.projectName}}-{{.env}}-jenkins ROLE
#----------------------------------------------------------------------------------------------------------------------

# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{.projectName}}_{{.env}}_jenkins" {
    name = "${var.project_name}-{{.env}}-jenkins"

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
resource "aws_iam_role_policy" "{{.projectName}}_{{.env}}_jenkins_update_route53" {
  name   = "${var.project_name}-{{.env}}-jenkins-update-route53-policy"
  role   = "${aws_iam_role.{{.projectName}}_{{.env}}_jenkins.id}"
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
resource "aws_iam_role_policy" "{{.projectName}}_{{.env}}_jenkins_ec2_specifics" {
  name   = "${var.project_name}-{{.env}}-jenkins-ec2-policy"
  role   = "${aws_iam_role.{{.projectName}}_{{.env}}_jenkins.id}"
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

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_ec2_container_registry_full_access" {
  role       = "${aws_iam_role.{{.projectName}}_{{.env}}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_s3_read_only_access" {
  role       = "${aws_iam_role.{{.projectName}}_{{.env}}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_ecs_full_access" {
  role       = "${aws_iam_role.{{.projectName}}_{{.env}}_jenkins.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonECS_FullAccess"

  # Tags not supported
}

#----------------------------------------------------------------------------------------------------------------------
# {{.projectName}}-{{.env}}-task-execution ROLE
#----------------------------------------------------------------------------------------------------------------------
# Creates a trust relationship with ecs-tasks.amazonaws.com and assumes that role
resource "aws_iam_role" "{{.projectName}}_task_execution" {
    name = "${var.project_name}-{{.env}}-task-execution"

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

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_ecs_task_execution" {
  role       = "${aws_iam_role.{{.projectName}}_task_execution.name}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"

  # Tags not supported
}

# Creates an inline policy to allow reading from System Manager Parameter Store
resource "aws_iam_role_policy" "parameter_store" {
  name   = "${var.project_name}-{{.env}}-parameter-store"
  role   = "${aws_iam_role.{{.projectName}}_task_execution.id}"
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

resource "aws_iam_role_policy_attachment" "{{.projectName}}_task_ec2_container_registry_full_access" {
  role   = "${aws_iam_role.{{.projectName}}_task_execution.id}"

  # AWS Mananged policy, safe to hard code
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"

  # Tags not supported
}

#----------------------------------------------------------------------------------------------------------------------
# {{.projectName}}-{{.env}}-lambda-execution ROLE
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "{{.projectName}}_{{.env}}_lambda_execution_role" {
  name = "${var.project_name}-{{.env}}-lambda-execution-role"
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
        },
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "ec2.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        },
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "xray.amazonaws.com"
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

output "lambda_exec_role" {
  value = aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role
}

resource "aws_iam_role_policy_attachment" "aws_xray_write_only_access" {
  role = "${var.project_name}-{{.env}}-lambda-execution-role"
  policy_arn = "arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess"
  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]
}

resource "aws_iam_role_policy_attachment" "aws_sagemaker_read_only_access" {
  role = "${var.project_name}-{{.env}}-lambda-execution-role"
  policy_arn = "arn:aws:iam::aws:policy/AmazonSageMakerReadOnly"
  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_attach_aws_lambda_role" {
  role       = "${var.project_name}-{{.env}}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaRole"

  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_attach_aws_lambda_role2" {
  role       = "${var.project_name}-{{.env}}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"

  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_attach_aws_lambda_role3" {
  role       = "${var.project_name}-{{.env}}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/ElasticLoadBalancingFullAccess"

  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]

  # Tags not supported
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_attach_aws_lambda_role4" {
  role       = "${var.project_name}-{{.env}}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2FullAccess"

  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]

  # Tags not supported
}

output "{{.projectName}}_lambda_execution_role" {
    value = "${aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role}"
}

# output "{{.projectName}}_lambda_exec_policy_attachment_policy" {
#     value ="${aws_iam_role_policy_attachment.{{.projectName}}_{{.env}}_attach_aws_lambda_role}"
# }

resource "aws_iam_policy" "{{.projectName}}_{{.env}}_lambda_dynamodb" {
  name = "${var.project_name}-{{.env}}-lambda-dynamodb-policy"
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
              "logs:CreateLogGroup",
              "s3:*"
          ],
          "Resource": "*"
      }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "{{.projectName}}_{{.env}}_lambda_dynamodb_policy_attach" {
  role       = "${var.project_name}-{{.env}}-lambda-execution-role"

  # managed by AWS so we can hard code it
  policy_arn = "${aws_iam_policy.{{.projectName}}_{{.env}}_lambda_dynamodb.arn}"

  depends_on = ["aws_iam_role.{{.projectName}}_{{.env}}_lambda_execution_role"]

  # Tags not supported
}

# output "{{.projectName}}_{{.env}}_lambda_dynamodb_policy" {
#   value = "${aws_iam_policy.{{.projectName}}_{{.env}}_lambda_dynamodb}"
# }


#----------------------------------------------------------------------------------------------------------------------
# ECS EC2 ROLES
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "ecsInstanceRole_{{.env}}" {
  name = "ecsInstanceRole-{{.env}}-${var.project_name}"

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

resource "aws_iam_role_policy" "ecsInstanceRolePolicy_{{.env}}" {
  name = "ecsInstanceRolePolicy-{{.env}}-${var.project_name}"
  role = "${aws_iam_role.ecsInstanceRole_{{.env}}.id}"

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
resource "aws_iam_role" "ecsServiceRole_{{.env}}" {
  name = "ecsServiceRole-{{.env}}-${var.project_name}"

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

resource "aws_iam_role_policy" "ecsServiceRolePolicy_{{.env}}" {
  name = "ecsServiceRolePolicy-{{.env}}-${var.project_name}"
  role = "${aws_iam_role.ecsServiceRole_{{.env}}.id}"

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

resource "aws_iam_service_linked_role" "es" {
  count            = var.create_elasticsearch_linked_service_role == "true" ? 1 : 0
  aws_service_name = "es.amazonaws.com"
}

#----------------------------------------------------------------------------------------------------------------------
# CODEBUILD ROLE
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "codebuild_role" {
  name = "${var.project_name}-codebuild-role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "sts:AssumeRole"
      ],
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "codebuild.amazonaws.com"
        ]
      }
    }
  ]
}
EOF
}

output "codebuild_role" {
  value = aws_iam_role.codebuild_role
}

resource "aws_iam_role_policy" "codebuild_policy" {
  name = "${var.project_name}-codebuild-policy"
  role = aws_iam_role.codebuild_role.id
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ecr:*",
        "codepipeline:*",
        "sagemaker:*",
        "s3:*",
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogStreams",
        "logs:GetLogEvents",
        "ssm:GetParameter",
        "kms:Encrypt",
        "kms:Decrypt",
        "iam:PassRole"
      ],
      "Effect": "Allow",
      "Resource": "*"
    },
    {
      "Action": [
        "iam:PassRole"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF

}

#----------------------------------------------------------------------------------------------------------------------
# SAGEMAKER ROLE
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "sagemaker_role" {
  name = "${var.project_name}-sagemaker-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
      "Action": [
          "sts:AssumeRole"
      ],
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "sagemaker.amazonaws.com"
        ]
      }
    }
  ]
}
EOF
}

output "sagemaker_role" {
  value = aws_iam_role.sagemaker_role
}

# TODO Limit the given roles to only those absolutely needed
resource "aws_iam_role_policy" "sagemaker_policy" {
  name = "${var.project_name}-sagemaker-policy"
  role = aws_iam_role.sagemaker_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject",
        "s3:PutObject",
        "s3:DeleteObject",
        "s3:ListBucket",
        "application-autoscaling:DeleteScalingPolicy",
        "application-autoscaling:DeleteScheduledAction",
        "application-autoscaling:DeregisterScalableTarget",
        "application-autoscaling:DescribeScalableTargets",
        "application-autoscaling:DescribeScalingActivities",
        "application-autoscaling:DescribeScalingPolicies",
        "application-autoscaling:DescribeScheduledActions",
        "application-autoscaling:PutScalingPolicy",
        "application-autoscaling:PutScheduledAction",
        "application-autoscaling:RegisterScalableTarget",
        "aws-marketplace:ViewSubscriptions",
        "cloudwatch:DeleteAlarms",
        "cloudwatch:DescribeAlarms",
        "cloudwatch:GetMetricData",
        "cloudwatch:GetMetricStatistics",
        "cloudwatch:ListMetrics",
        "cloudwatch:PutMetricAlarm",
        "cloudwatch:PutMetricData",
        "codecommit:BatchGetRepositories",
        "codecommit:CreateRepository",
        "codecommit:GetRepository",
        "codecommit:ListBranches",
        "codecommit:ListRepositories",
        "cognito-idp:AdminAddUserToGroup",
        "cognito-idp:AdminCreateUser",
        "cognito-idp:AdminDeleteUser",
        "cognito-idp:AdminDisableUser",
        "cognito-idp:AdminEnableUser",
        "cognito-idp:AdminRemoveUserFromGroup",
        "cognito-idp:CreateGroup",
        "cognito-idp:CreateUserPool",
        "cognito-idp:CreateUserPoolClient",
        "cognito-idp:CreateUserPoolDomain",
        "cognito-idp:DescribeUserPool",
        "cognito-idp:DescribeUserPoolClient",
        "cognito-idp:ListGroups",
        "cognito-idp:ListIdentityProviders",
        "cognito-idp:ListUserPoolClients",
        "cognito-idp:ListUserPools",
        "cognito-idp:ListUsers",
        "cognito-idp:ListUsersInGroup",
        "cognito-idp:UpdateUserPool",
        "cognito-idp:UpdateUserPoolClient",
        "ec2:CreateNetworkInterface",
        "ec2:CreateNetworkInterfacePermission",
        "ec2:CreateVpcEndpoint",
        "ec2:DeleteNetworkInterface",
        "ec2:DeleteNetworkInterfacePermission",
        "ec2:DescribeDhcpOptions",
        "ec2:DescribeNetworkInterfaces",
        "ec2:DescribeRouteTables",
        "ec2:DescribeSecurityGroups",
        "ec2:DescribeSubnets",
        "ec2:DescribeVpcEndpoints",
        "ec2:DescribeVpcs",
        "ecr:BatchCheckLayerAvailability",
        "ecr:BatchGetImage",
        "ecr:CreateRepository",
        "ecr:Describe*",
        "ecr:GetAuthorizationToken",
        "ecr:GetDownloadUrlForLayer",
        "elastic-inference:Connect",
        "elasticfilesystem:DescribeFileSystems",
        "elasticfilesystem:DescribeMountTargets",
        "fsx:DescribeFileSystems",
        "glue:CreateJob",
        "glue:DeleteJob",
        "glue:GetJob",
        "glue:GetJobRun",
        "glue:GetJobRuns",
        "glue:GetJobs",
        "glue:ResetJobBookmark",
        "glue:StartJobRun",
        "glue:UpdateJob",
        "groundtruthlabeling:*",
        "iam:ListRoles",
        "kms:DescribeKey",
        "kms:ListAliases",
        "lambda:ListFunctions",
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:DescribeLogStreams",
        "logs:GetLogEvents",
        "logs:PutLogEvents",
        "sns:ListTopics",
        "ssm:GetParameter"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

#----------------------------------------------------------------------------------------------------------------------
# CODEPIPELINE ROLE
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "codepipeline_role" {
  name = "${var.project_name}-codepipeline-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": ["sts:AssumeRole"],
      "Principal": {
        "Service": "codepipeline.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

output "codepipeline_role" {
  value = aws_iam_role.codepipeline_role
}

resource "aws_iam_role_policy" "codepipeline_policy" {
  name = "${var.project_name}-codepipeline-policy"
  role = aws_iam_role.codepipeline_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*",
        "codebuild:*",
        "codepipeline:*",
        "cloudformation:CreateStack",
        "cloudformation:DescribeStacks",
        "cloudformation:DeleteStack",
        "cloudformation:UpdateStack",
        "cloudformation:CreateChangeSet",
        "cloudformation:ExecuteChangeSet",
        "cloudformation:DeleteChangeSet",
        "cloudformation:DescribeChangeSet",
        "cloudformation:SetStackPolicy",
        "iam:PassRole",
        "sns:Publish"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

#----------------------------------------------------------------------------------------------------------------------
# CLOUDFORMATION ROLE
#----------------------------------------------------------------------------------------------------------------------
resource "aws_iam_role" "cloudformation_role" {
  name = "${var.project_name}-cloudformation-role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "sts:AssumeRole"
      ],
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cloudformation.amazonaws.com"
        ]
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "cloudformation_policy" {
  name = "${var.project_name}-cloudformation-policy"
  role = aws_iam_role.cloudformation_role.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
      {
        "Action": [
          "sagemaker:*",
          "iam:PassRole",
          "s3:*"
        ],
        "Effect": "Allow",
        "Resource": "*"
      }
  ]
}
EOF
}

#----------------------------------------------------------------------------------------------------------------------
# VPC FLOW LOGS
#----------------------------------------------------------------------------------------------------------------------

resource "aws_iam_role" "vpc_flow_log" {
  name = "${var.project_name}-vpc-flow-log"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "vpc-flow-logs.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

output "vpc_flow_log_role" {
  value = aws_iam_role.vpc_flow_log
}

resource "aws_iam_role_policy" "vpc_log_policy" {
  name = "${var.project_name}-vpc-flow-log-policy"
  role = aws_iam_role.vpc_flow_log.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogGroups",
        "logs:DescribeLogStreams"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
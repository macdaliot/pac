#==============================================================================
# IAM
#==============================================================================

locals {
  users = [
    "testalicious1",
    "testalicious2"
  ]

  policies = [
    "arn:aws:iam::aws:policy/PowerUserAccess",
    "arn:aws:iam::aws:policy/IAMFullAccess"
  ]
}

resource "aws_iam_group" "developers" {
  name = "developers"
}

resource "aws_iam_group_policy_attachment" "developers-policies" {
  count      = length(local.policies)
  group      = aws_iam_group.developers.name
  policy_arn = element(local.policies, count.index)
}

resource "aws_iam_user" "u" {
  count = length(local.users)
  name  = element(local.users, count.index)
}

resource "aws_iam_user_group_membership" "developers-membership" {
  count = length(local.users)
  user  = element(local.users, count.index)
  groups = [
    aws_iam_group.developers.name
  ]
}

resource "aws_iam_user_login_profile" "default" {
  count   = length(local.users)
  user    = element(local.users, count.index)
  pgp_key = "keybase:some_person_that_exists"
}

output "password" {
  value = "${aws_iam_user_login_profile.example.encrypted_password}"
}
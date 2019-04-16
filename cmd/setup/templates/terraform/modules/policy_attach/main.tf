#
# https://www.terraform.io/docs/providers/aws/r/iam_role_policy_attachment.html
# attach policy to role
#
resource "aws_iam_role_policy_attachment" "policy_attach" {
  role       = "${var.role_name}"
  policy_arn = "${var.policy_arn}"
}

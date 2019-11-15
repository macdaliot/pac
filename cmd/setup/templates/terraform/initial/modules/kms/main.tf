#
# https://www.terraform.io/docs/providers/aws/r/kms_key.html
#
resource "aws_kms_key" "key" {
  deletion_window_in_days = var.deletion_window
}

#
# https://www.terraform.io/docs/providers/aws/r/kms_alias.html
#
resource "aws_kms_alias" "alias" {
  name          = var.alias_name
  target_key_id = aws_kms_key.key.key_id
}

# https://www.terraform.io/docs/providers/aws/r/kms_key.html
resource "aws_kms_key" "k" {
    deletion_window_in_days = var.deletion_window
}

resource "aws_kms_alias" "a" {
  name          = "alias/pac/${var.project_name}"
  target_key_id = aws_kms_key.k.key_id
}
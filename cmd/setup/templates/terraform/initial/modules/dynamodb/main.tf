#
# https://www.terraform.io/docs/providers/aws/r/dynamodb_table.html
#
resource "aws_dynamodb_table" "table" {
  attribute {
    name = var.attribute_name
    type = var.attribute_type
  }
  billing_mode     = var.billing_mode
  hash_key         = var.attribute_name
  name             = var.table_name
  read_capacity    = var.read_capacity
  stream_enabled   = var.stream_enabled
  stream_view_type = var.stream_view_type
  write_capacity   = var.write_capacity
}

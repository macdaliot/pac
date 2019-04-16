#
# https://www.terraform.io/docs/providers/aws/r/dynamodb_table.html
# dynamo_db
resource "aws_dynamodb_table" "dynamodb_table" {
  name           = "pac-${var.project_name}-${var.resource_name}"
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "id"
  stream_view_type = "NEW_AND_OLD_IMAGES"
  stream_enabled = true

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name        = "dynamodb-table-${var.resource_name}"
  }
}
#
# https://www.terraform.io/docs/providers/aws/r/dynamodb_table.html
resource "aws_dynamodb_table" "t" {
  name             = var.table_name
  billing_mode     = var.billing_mode     
  read_capacity    = var.read_capacity    
  write_capacity   = var.write_capacity   
  hash_key         = var.attribute_name       
  stream_view_type = var.stream_view_type 
  stream_enabled   = var.stream_enabled   

  attribute {
    name = var.attribute_name
    type = var.attribute_type
  }

  # tags = {
  #   Name             = "${var.project_name}-dynamodb-table-celebrity"
  #   pac-project-name = var.project_name
  #   environment      = var.environment_name
  # }
}
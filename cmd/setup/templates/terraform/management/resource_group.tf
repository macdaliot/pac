resource "aws_resourcegroups_group" "main" {
  name = "${var.project_name}"

  resource_query {
    query = <<JSON
{
  "ResourceTypeFilters": [
    "AWS::AllSupported"
  ],
  "TagFilters": [
    {
      "Key": "pac-project-name",
      "Values": ["${var.project_name}"]
    }
  ]
}
JSON
  }
}

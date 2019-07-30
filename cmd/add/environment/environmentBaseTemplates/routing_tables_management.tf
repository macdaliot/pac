# This allows traffic between the management vpc and application vpc.
# Since the application vpc is created at an undetermined time after the management vpc
# we are creating this route table during the application vpc lifecycle.
resource "aws_route_table" "management_vpc" {
  vpc_id = data.terraform_remote_state.management.outputs.management_vpc_id

  route {
    cidr_block                = aws_vpc.application_vpc.cidr_block
    vpc_peering_connection_id = aws_vpc_peering_connection.pc.id
  }

  tags = {
    Name = "${var.project_name} vpc peering route table"
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}


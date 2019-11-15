#
# Creating a VPC does not provide
# an Internet Gateway so we are creating one here.
#
resource "aws_internet_gateway" "igw" {
  vpc_id = var.vpc_id

  tags = {
    "Name" = "Internet Gateway"
  }
}

resource "aws_eip" "gw" {
  vpc        = true
  depends_on = ["aws_internet_gateway.igw"]
}
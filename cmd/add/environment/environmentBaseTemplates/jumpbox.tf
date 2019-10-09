
#----------------------------------------------------------------------------------------------------------------------
# JUMPBOX TO ALLOW ACCESS TO KIBANA AND DOCUMENTDB
#
# https://www.jeremydaly.com/access-aws-vpc-based-elasticsearch-cluster-locally/
# https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html
# https://forums.aws.amazon.com/thread.jspa?threadID=279437
#----------------------------------------------------------------------------------------------------------------------

data "aws_ami" "amzn" {
  most_recent = true

  filter {
    name   = "name"
    values = [var.jumpbox_ami]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["137112412989"] # Amazon
}

resource "aws_security_group" "jumpbox" {
  count                  = (var.enable_jumpbox || var.enable_documentdb) == "true" ? 1 : 0
  name                   = "${var.project_name}-jumpbox"
  description            = "controls access to Kibana and DocumentDB"
  vpc_id                 = aws_vpc.application_vpc.id
  revoke_rules_on_delete = true

  ingress {
    protocol    = "tcp"
    from_port   = "22"
    to_port     = "22"
    cidr_blocks = [var.end_user_cidr]
  }

  ingress {
    protocol    = "tcp"
    from_port   = "443"
    to_port     = "443"
    cidr_blocks = [var.end_user_cidr]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

resource "aws_instance" "jumpbox" {
  count                       = (var.enable_jumpbox || var.enable_documentdb) == "true" ? 1 : 0
  ami                         = data.aws_ami.amzn.id
  associate_public_ip_address = true
  instance_type               = "t2.micro"

  # referring to the key pair to be used to SSH into box
  key_name               = "jumpbox-${var.project_name}"
  subnet_id              = aws_subnet.public[0].id
  vpc_security_group_ids = [aws_security_group.jumpbox[0].id]

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = var.environment_name
  }
}

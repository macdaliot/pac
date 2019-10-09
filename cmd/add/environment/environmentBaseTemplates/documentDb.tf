resource "aws_docdb_cluster_parameter_group" "docdb_params" {
  count       = var.enable_documentdb == "true" ? 1 : 0
  family      = "docdb3.6"
  name        = "${var.project_name}-${var.environment_abbr}"
  description = "Managed by Terraform"

  parameter {
    name  = "tls"
    value = "disabled"
  }
}

resource "aws_docdb_subnet_group" "appsubnets" {
  count      = var.enable_documentdb == "true" ? 1 : 0
  name       = "documentdb-${var.project_name}"
  subnet_ids = [aws_subnet.public[0].id, aws_subnet.public[1].id, aws_subnet.public[2].id]


  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
  }
}

resource "aws_security_group" "documentdb" {
  count       = var.enable_documentdb == "true" ? 1 : 0
  name        = "documentdb-${var.project_name}"
  description = "Managed by Terraform"
  vpc_id      = aws_vpc.application_vpc.id

  ingress {
    from_port = 1025
    to_port   = 65000
    protocol  = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
  }
}

resource "aws_docdb_cluster_instance" "cluster_instances" {
  count              = var.enable_documentdb == "true" ? 2 : 0
  identifier         = "${var.project_name}-${var.environment_abbr}-instance-${count.index}"
  cluster_identifier = aws_docdb_cluster.docdb[0].id
  instance_class     = "db.r5.large"

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
  }

}

resource "aws_docdb_cluster" "docdb" {
  count                   = var.enable_documentdb == "true" ? 1 : 0
  cluster_identifier      = "${var.project_name}-${var.environment_abbr}"
  engine                  = "docdb"
  master_username         = var.documentdb_user
  master_password         = var.documentdb_password
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
  skip_final_snapshot     = true
  db_cluster_parameter_group_name = aws_docdb_cluster_parameter_group.docdb_params[0].name
  vpc_security_group_ids = [aws_security_group.documentdb[0].id]
  db_subnet_group_name = aws_docdb_subnet_group.appsubnets[0].name
}

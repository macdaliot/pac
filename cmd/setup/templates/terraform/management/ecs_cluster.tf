resource "aws_ecs_cluster" "main" {
  name = var.project_name

  tags = {
    Name             = var.project_name
    pac-project-name = var.project_name
    environment      = "management"
  }
}

output "ecs_cluster_name" {
  value = aws_ecs_cluster.main.name
}
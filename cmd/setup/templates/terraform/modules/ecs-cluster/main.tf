#
#https://www.terraform.io/docs/providers/aws/r/ecs_cluster.html
#
resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${data.terraform_remote_state.vpc.project_name}-ecs-cluster"
}
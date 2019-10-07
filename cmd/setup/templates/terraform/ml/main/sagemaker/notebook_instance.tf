# # https://www.terraform.io/docs/providers/aws/r/sagemaker_notebook_instance.html

resource "aws_sagemaker_notebook_instance" "ni" {
  name                  = "${var.project_name}-notebook-instance"
  role_arn              = var.management.outputs.sagemaker_role.arn
  instance_type         = "ml.t2.2xlarge"
  subnet_id             = var.management.outputs.private_subnet[0].id
  security_groups       = [var.management.outputs.secgroup_application_lb.id]
  lifecycle_config_name =  aws_sagemaker_notebook_instance_lifecycle_configuration.lc.name
  tags = {
    Name = "${var.project_name} Jupytr"
  }
}
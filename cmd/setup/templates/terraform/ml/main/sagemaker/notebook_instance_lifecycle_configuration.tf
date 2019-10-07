# https://www.terraform.io/docs/providers/aws/r/sagemaker_notebook_instance_lifecycle_configuration.html
resource "aws_sagemaker_notebook_instance_lifecycle_configuration" "lc" {
  name      = "${var.project_name}-ml-lc-conf"
  on_create = base64encode(data.template_file.create_script.rendered)
}

data "template_file" "create_script" {
  template = file("${path.module}/create-script.sh")
}

# https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint_configuration.html
resource "aws_sagemaker_endpoint_configuration" "ec_pro" {
  name   = "${var.environment_abbr}-${var.project_name}-endpoint-config"
  # region = var.region
  production_variants {
    initial_instance_count = 1
    initial_variant_weight = 1
    instance_type          = "ml.m4.xlarge"
    model_name             = "${var.model_name}-${var.environment_abbr}"
    variant_name           = "AllTraffic"
  }

  tags = {
    Name = "${var.project_name} - Sagemaker Endpoint Configuration"
  }

  depends_on = [aws_sagemaker_model.sagemaker_model_pro]
}
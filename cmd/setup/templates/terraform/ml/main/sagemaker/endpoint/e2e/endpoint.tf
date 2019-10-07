# https://www.terraform.io/docs/providers/aws/r/sagemaker_endpoint.html

resource "aws_sagemaker_endpoint" "endpoint_e2e" {
  name                 = "${var.environment_abbr}-${var.model_name}"
  endpoint_config_name = aws_sagemaker_endpoint_configuration.ec_e2e.name

  tags = {
    Name = "${var.project_name} Sagemaker Endpoint"
  }

  depends_on = [aws_sagemaker_endpoint_configuration.ec_e2e, aws_sagemaker_model.sagemaker_model_e2e]
}
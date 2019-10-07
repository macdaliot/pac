module "sagemaker_1" {
  source = "./sagemaker"
  development               = data.terraform_remote_state.development
  e2etesting                = data.terraform_remote_state.e2etesting
  management                = data.terraform_remote_state.management
  enable_cloudwatch_metrics = var.enable_cloudwatch_metrics
  image                     = var.image
  log_level                 = var.log_level
  project_name              = var.project_name
  region                    = var.region
  sagemaker_image_name  = var.sagemaker_image_name
  stack_name                = var.stack_name
}
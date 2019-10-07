# # https://www.terraform.io/docs/providers/aws/r/sagemaker_model.html
resource "aws_sagemaker_model" "sagemaker_model_e2e" {
  name               = "${var.model_name}-${var.environment_abbr}"
  execution_role_arn = var.sagemaker_role
  # vpc_config {
  #   security_group_ids   = [data.terraform_remote_state.development.outputs.secgroup_application_lb.id]
  #   subnets              = [#"subnet-0fd3e80836abf685c", "subnet-091fcafc779ee3275", "subnet-08f57ab54555c2804"
  #     data.terraform_remote_state.development.outputs.public_subnet[0].id,
  #     data.terraform_remote_state.development.outputs.public_subnet[1].id
  #      ]
  # }


  primary_container {
    image = "${var.image_repo}${var.sagemaker_image_name}"
    model_data_url    =  var.model_data_url
    environment = {
      SAGEMAKER_CONTAINER_LOG_LEVEL       = var.log_level
      SAGEMAKER_ENABLE_CLOUDWATCH_METRICS = var.enable_cloudwatch_metrics
      SAGEMAKER_PROGRAM                   = "bert_embeddings.py"
      SAGEMAKER_REGION                    = var.region
      SAGEMAKER_SUBMIT_DIRECTORY          = var.source_directory
    }

    # 342720212717 = var.342720212717
  }
}

# resource "aws_sagemaker_model" "sagemaker_model_e2e" {
#   name               = "${data.terraform_remote_state.e2etesting.outputs.environment_abbr}-${var.model_name}"
#   execution_role_arn = var.sagemaker_role
  

#   primary_container {
#     image = var.image

#     environment = {
#       SAGEMAKER_CONTAINER_LOG_LEVEL       = var.log_level
#       SAGEMAKER_ENABLE_CLOUDWATCH_METRICS = var.enable_cloudwatch_metrics
#       # SAGEMAKER_PROGRAM                   = var.sagemaker_program
#       SAGEMAKER_REGION                    = var.region
#       # SAGEMAKER_SUBMIT_DIRECTORY          = var.source_directory
#     }

#     # 342720212717 = var.342720212717
#   }
# }

# resource "aws_sagemaker_model" "sagemaker_model" {
#   name               = "${var.environment_abbr}-${var.model_name}"
#   execution_role_arn = var.sagemaker_role
  

#   primary_container {
#     image = var.image

#     environment = {
#       SAGEMAKER_CONTAINER_LOG_LEVEL       = var.log_level
#       SAGEMAKER_ENABLE_CLOUDWATCH_METRICS = var.enable_cloudwatch_metrics
#       # SAGEMAKER_PROGRAM                   = var.sagemaker_program
#       SAGEMAKER_REGION                    = var.region
#       # SAGEMAKER_SUBMIT_DIRECTORY          = var.source_directory
#     }

#     # 342720212717 = var.342720212717
#   }
# }
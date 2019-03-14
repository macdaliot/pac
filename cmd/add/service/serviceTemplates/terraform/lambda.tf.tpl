# Create lambda resource
module "lambda_{{ .Resource }}" {
  source = "./modules/lambda"

  lambda_function_payload = "function.zip"
  resource_name           = "{{ .Resource }}"
  project_name            = "{{ .ProjectName }}"
  role                    = "${module.roles.pac_lambda_execution_role_arn}"
}

# Create target group. Only one lambda can be registered per target group
# https://docs.aws.amazon.com/elasticloadbalancing/latest/application/lambda-functions.html
module "{{ .ProjectName }}_target_group" {
  source = "./modules/alb_target_group"

  project_name = "{{ .ProjectName }}"
  resource_name = "{{ .Resource }}"
}

module "{{ .ProjectName }}_lambda_permission_with_lb" {
  source = "./modules/aws_lambda_permission"

  project_name = "{{ .ProjectName }}"
  resource_name = "{{ .Resource }}"
  source_arn = "${module.{{ .ProjectName }}_target_group.arn}"
}

# register with load balancer target group 
module "{{ .ProjectName }}_lambda_target_group_attachment" {
  source = "./modules/alb_target_group_attachment"

  alb_target_group_arn = "${module.{{ .ProjectName }}_target_group.arn}"
  lambda_arn           = "${module.lambda_{{ .Resource }}.arn}"
}

module "dynamodb_{{ .Resource }}_table" {
  source = "./modules/dynamo_db"

  project_name = "{{ .ProjectName }}"
  resource_name = "{{ .Resource }}"
}
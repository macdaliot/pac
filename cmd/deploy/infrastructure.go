package deploy

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/config"
)

// Infrastructure - Calls on Terraform to create the AWS infrastructure
func Infrastructure() {
	// The directory Terraform should run in relative to the current directory (project directory)
	terraformDirectory := "terraform"

	// Run `terraform init`
	output := terraform.Initialize(terraformDirectory)
	logger.Info(output)

	// Run `terraform plan`
	output = terraform.Plan(terraformDirectory, createPlanConfig())
	logger.Info(output)

	// Run `terraform apply`
	output = terraform.Apply(terraformDirectory)
	logger.Info(output)
}

func createPlanConfig() map[string]string {
	cfg := make(map[string]string)
	cfg["application_cidr_block"] = config.Get("awsApplicationVpcCidrBlock")
	cfg["management_cidr_block"] = config.Get("awsManagementVpcCidrBlock")
	return cfg
}

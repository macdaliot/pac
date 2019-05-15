package deploy

import (
	"os"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/config"
)

// Infrastructure calls on Terraform to create the AWS infrastructure. It takes a parameter which is used to determine
// which folder in the Terraform structure to execute in relative to the project root.
func Infrastructure(path string) {
	// The directory Terraform should run in relative to the current directory (project directory)
	terraformDirectory, err := ValidatePath(path)
	if err != nil {
		validatePathError := "The following error occured when trying to validate a Terraform directory: "
		errors.LogAndQuit(validatePathError + err.Error())
	}

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

// ValidatePath checks that the given path to a Terraform template directory is valid.
// If the path is valid it is returned and eror is set to nil. Otherwise not then a non-nil error is returned.
func ValidatePath(dir string) (path string, err error) {
	path = "terraform"

	switch dir {
	case "dns":
		path += "/dns"
	case "ssl":
		path += "/ssl"

	}

	_, err = os.Stat(path)

	return path, err
}

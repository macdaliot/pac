package deploy

import (
	"os"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/util"
)

// Infrastructure calls on Terraform to create the AWS infrastructure. It takes a parameter which is used to determine
// which folder in the Terraform structure to execute in relative to the project root.
func Infrastructure(path string) {
	util.GoToRootDirectory()

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
	cfg := make(map[string]string)
	output = terraform.Plan(terraformDirectory, cfg)
	logger.Info(output)

	// Run `terraform apply`
	output = terraform.Apply(terraformDirectory)
	logger.Info(output)
}

// ValidatePath checks that the given path to a Terraform template directory is valid.
// If the path is valid, the path is returned and the error is set to nil. Otherwise, a non-nil error is returned.
func ValidatePath(dir string) (path string, err error) {
	path = "terraform/" + dir
	_, err = os.Stat(path)
	return path, err
}

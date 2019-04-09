package setup

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
)

//IsTerraformInstalled attempts to get the Terraform version to demonstrate Terraform is installed and accessible.
//If Terraform is not installed or accessible the execution of the program is stopped.
func IsTerraformInstalled() {
	logger.Info("Checking if Terraform is installed...")
  _, err := commands.Run("terraform version", "")
  if err != nil {
    errors.LogAndQuit("ERROR D")
  }
  logger.Info("Terraform is installed")
}

//SetTerraformEnv sets the environment variable for Terraform automation
func SetTerraformEnv() {
	os.Setenv("TF_IN_AUTOMATION", "NONEMPTYVALUE")
	os.Setenv("TF_VAR_project_name", config.Get("projectName"))
}

//TerraformInitialize initializes the terraform directory
//checks for *.tf files and processes them. By default
//there will only be one *.tf file, which sets up the
//S3 backend where infrastructure state will be stored
func TerraformInitialize() {
  output, err := commands.Run("terraform init -input=false", "terraform")
  if err != nil {
    errors.LogAndQuit("ERROR A")
  }
  logger.Info(output)
}

//TerraformCreate creates tfplan that will be applied by Terraform to create AWS infrastructure.
func TerraformCreate() {
  output, err := commands.Run("terraform plan -out=tfplan -input=false", "terraform")
  if err != nil {
    errors.LogAndQuit("ERROR B")
  }
  logger.Info(output)
}

//TerraformApply applies tfplan
func TerraformApply() {
	defer timeTrack(time.Now(), "Terraform create")
  output, err := commands.Run("terraform apply -input=false tfplan", "terraform")
  if err != nil {
    errors.LogAndQuit("ERROR C")
  }
  logger.Info(output)
}

//TerraformDestroy destroys all resources managed by Terraform
func TerraformDestroy() {
	defer timeTrack(time.Now(), "Terraform destroy")

	destroyLambdas()
	destroyVPCs()
}

func destroyVPCs() {
	logger.Info("Terraform is destroying VPCs...")
	//navgiate to terraform directory
	os.Chdir(config.GetRootDirectory())
	os.Chdir("./terraform")

	if _, err := os.Stat("./.terraform"); os.IsNotExist(err) {
		//Terraform did not initialize, there is nothing to destroy
		return
	}

	// destroy AWS resources managed by Terraform
	cmd := exec.Command("terraform", "destroy", "-auto-approve")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Terraform destroy failed with %s\n", err)
	}

	logger.Info(string(out))
}

func destroyLambdas() {
	logger.Info("Terraform is destroying lambdas(services)...")
	//Navigate to svc directory
	os.Chdir(config.GetRootDirectory())
	os.Chdir("./svc/terraform")

	if _, err := os.Stat("./.terraform"); os.IsNotExist(err) {
		//Terraform did not initialize, there is nothing to destroy
		return
	}

	// destroy AWS resources managed by Terraform
	cmd := exec.Command("terraform", "destroy", "-auto-approve")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Terraform destroy failed with %s\n", err)
	}

	logger.Info(string(out))
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

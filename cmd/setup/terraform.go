package setup

import (
  "fmt"
  "os"
  "time"

  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

//IsTerraformInstalled attempts to get the Terraform version to demonstrate Terraform is installed and accessible.
//If Terraform is not installed or accessible the execution of the program is stopped.
func IsTerraformInstalled() {
  logger.Info("Checking if Terraform is installed...")
  _, err := commands.Run("terraform version", "")
  if err != nil {
    errors.LogAndQuit(str.Concat("ERROR: Checking the Terraform version failed with the following error: ", err.Error()))
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
    errors.LogAndQuit(str.Concat("ERROR: Initializing Terraform failed with the following error: ", err.Error()))
  }
  logger.Info(output)
}

//TerraformPlan creates tfplan that will be applied by Terraform to create AWS infrastructure.
func TerraformPlan(freeVpcCidrBlocks []string) {
  output, err := commands.Run("terraform plan -var management_cidr_block=" + freeVpcCidrBlocks[0] + " -var application_cidr_block=" + freeVpcCidrBlocks[1] + " -var region=us-east-2 -out tfplan", "terraform")
  if err != nil {
    errors.LogAndQuit(str.Concat("ERROR: Planning with Terraform failed with the following error: ", err.Error()))
  }
  logger.Info(output)
}

//TerraformApply applies tfplan
func TerraformApply() {
  defer timeTrack(time.Now(), "Terraform create")
  output, err := commands.Run("terraform apply -input=false tfplan", "terraform")
  if err != nil {
    errors.LogAndQuit(str.Concat("ERROR: Applying the Terraform plan failed with the following error: ", err.Error()))
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
  if files.Exists("terraform/.terraform") {
    logger.Info("Terraform is destroying VPCs...")
    output, err := commands.Run("terraform destroy -auto-approve", "terraform")
    if err != nil {
      errors.LogAndQuit(str.Concat("ERROR: Terraform destroy failed with the following error: ", err.Error()))
    }
    logger.Info(output)
  } else {
    logger.Info("No VPCs were created through Terraform. No project infrastructure to clean")
  }
}

func destroyLambdas() {
  if files.Exists("svc/terraform/.terraform") {
    logger.Info("Terraform is destroying lambdas (services)...")
    output, err := commands.Run("terraform destroy -auto-approve", "svc/terraform")
    if err != nil {
      errors.LogAndQuit(str.Concat("ERROR: Terraform destroy failed with the following error: ", err.Error()))
    }
    logger.Info(output)
  } else {
    logger.Info("No services were created through Terraform. No lambda functions to clean")
  }
}

func timeTrack(start time.Time, name string) {
  elapsed := time.Since(start)
  fmt.Sprintf("%s took %s", name, elapsed)
}

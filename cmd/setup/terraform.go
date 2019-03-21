package setup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
)

//IsTerraformInstalled attempts to get the Terraform version to demonstrate Terraform is installed and accessible.
//If Terraform is not installed or accessible the execution of the program is stopped.
func IsTerraformInstalled() {
	logger.Info("Checking if Terraform is installed...")
	cmd := exec.Command("terraform", "version")

	err := cmd.Run()
	if err != nil {
		log.Fatal("Terraform is not installed: ", err)
	}

	logger.Info("Terraform is installed.\n")
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
	//switch to terraform directory
	path := path.Join(config.GetRootDirectory(), "terraform")
	os.Chdir(path)

	//initialize terraform
	cmd := exec.Command("terraform", "init", "-input=false")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(out))
}

//TerraformCreate creates tfplan that will be applied by Terraform to create AWS infrastructure.
func TerraformCreate() {
	cmd := exec.Command("terraform", "plan", "-out=tfplan", "-input=false")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(out))
}

//TerraformApply applies tfplan
func TerraformApply() {
	defer timeTrack(time.Now(), "Terraform create")
	cmd := exec.Command("terraform", "apply", "-input=false", "tfplan ")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(out))
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

package setup

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
)

//IsTerrafomInstalled attempts to get the Terraform version to demonstrate Terraform is installed and accessible.
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

func SetTerraformEnv() {
	os.Setenv("TF_IN_AUTOMATION", "NONEMPTYVALUE")
}

//InitializeTerraform initializes the terraform directory
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

// CreateTerraform creates tfplan that will be applied by Terraform to create AWS infrastructure.
func TerraformCreate() {
	cmd := exec.Command("terraform", "plan", "-out=tfplan", "-input=false")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(out))
}

// ApplyTerraform applies tfplan
func TerraformApply() {
	cmd := exec.Command("terraform", "apply", "-input=false", "tfplan ")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(out))
}

// DestroyTerraform destroys all resources managed by Terraform
func TerraformDestroy() {
	//navgiate to terraform directory
	os.Chdir(config.GetRootDirectory())
	os.Chdir("./terraform")

	// destory AWS resources managed by Terraform
	cmd := exec.Command("terraform", "destroy", "-auto-approve")

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Terraform destroy failed with %s\n", err)
	}

	fmt.Println(string(out))
}

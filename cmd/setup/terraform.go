package setup

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/PyramidSystemsInc/go/logger"
)

//Provider struct provides fields needed to create terraform provider resource.
type Provider struct {
	ProjectName     string
	Region          string
	AWSVersion      string
	TemplateVersion string
}

//CreateProviderTemplate creates a Provider struct and populates its fields with the provided parameters
//The struct is then merged with a template file and saved as a terraform file.
func CreateTerraformProvider(projectName, region, awsVersion, templateVersion string) error {
	//create an instance of Provider
	p := Provider{ProjectName: projectName, Region: region, AWSVersion: awsVersion, TemplateVersion: templateVersion}

	//create a new template
	tmpl := template.New("providerTemplate")

	//parse content of template file
	content, err := ioutil.ReadFile("providers.tpl")
	if err != nil {
		log.Fatal(err)
	}

	//parse template
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		log.Fatal("Parse: ", err)
		return err
	}

	//create terraform file
	tf, err := os.OpenFile("providers.tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal("unable to create terraform file providers.tf")
	}
	defer tf.Close()

	//apply template to Provider data structure and write to terraform file
	err = tmpl.Execute(tf, p)
	if err != nil {
		log.Fatal("Execute: ", err)
		return err
	}

	logger.Info("The providers.tf has been created.")

	return nil
}

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

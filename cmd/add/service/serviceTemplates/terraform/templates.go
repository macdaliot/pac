package service

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

//Provider struct hold values used to create Terraform AWS Provider
type Provider struct {
	ProjectName     string
	Region          string
	AWSVersion      string
	TemplateVersion string
}

//Lambda struct provides the necessary fields to populate a AWS Lambda resource template
//needed by the Pyramid Application Constructor
type Lambda struct {
	ProjectName  string
	ResourceName string
}

//ParseTemplate takes in a path to a template file and an empty interface
//The file is opened and the template it contains and the instaniated type passed in for the interface
//are merged and saved as a terraform file
func ParseTemplate(resourceName string, file string, lambda interface{}) {
	//create a new template
	tmpl := template.New("myTemplate")

	//parse content of template file
	content, _ := ioutil.ReadFile(file)

	//parse template
	tmpl, err := tmpl.Parse(string(content))
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	//terraform file
	tf, err := os.OpenFile("lambda_"+resourceName+".tf", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal("unable to create terraform file")
	}
	defer tf.Close()

	//apply template to Lambda data structure and write to terraform file
	err = tmpl.Execute(tf, lambda)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}

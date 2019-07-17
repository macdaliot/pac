package authservice

import (
	"strings"

	"github.com/gobuffalo/packr"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
)

func CreateAllTemplatedFiles(cfg map[string]string) {
	// service code
	options := files.TemplateOptions{
		TargetDirectory: "services/auth",
		Box:             packr.NewBox("./authServiceTemplates"),
		Config:          cfg,
	}
	err := files.CreateTemplatedFiles(options)
	errors.QuitIfError(err)

	// Terraform templates
	environmentNames := strings.Split(config.Get("environments"), ",")
	for _, environmentName := range environmentNames {
		cfg["environmentName"] = environmentName
		options = files.TemplateOptions{
			TargetDirectory: str.Concat("terraform/", environmentName),
			Box:             packr.NewBox("./authTerraformTemplates"),
			Config:          cfg,
		}
		err = files.CreateTemplatedFiles(options)
		errors.QuitIfError(err)
	}
}

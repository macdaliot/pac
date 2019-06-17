package authservice

import (
	"github.com/PyramidSystemsInc/go/files"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(config map[string]string) {
	options := files.TemplateOptions{
		TargetDirectory: "services/auth",
		Box:             packr.NewBox("./authServiceTemplates"),
		Config:          config}
	files.CreateTemplatedFiles(options)

	options.TargetDirectory = "services/terraform"
	options.Box = packr.NewBox("./authTerraformTemplates")
	files.CreateTemplatedFiles(options)
}

package service

import (
	"os"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/gobuffalo/packr"
)

//CreateAllTemplatedFiles says comment your code Jeff!
func CreateAllTemplatedFiles(c map[string]string) {
	box := packr.NewBox("./serviceTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(templatePath, template, c)
	}

	os.Chdir(config.GetRootDirectory())
	os.Chdir("./svc/terraform")

	box = packr.NewBox("./terraformTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(templatePath, template, c)
	}

	os.Rename("./lambda.tf", "./lambda_"+c["serviceName"]+".tf")
}

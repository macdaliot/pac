package authservice

import (
	"path/filepath"

	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(config map[string]string) {
  targetDirectory := "services/auth"
	box := packr.NewBox("./authServiceTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
    fullPath := filepath.Join(targetDirectory, templatePath)
    files.EnsurePath(filepath.Dir(fullPath))
    template, _ := box.FindString(templatePath)
    files.CreateFromTemplate(fullPath, template, config)
	}

  targetDirectory = "services/terraform"
  box = packr.NewBox("./authTerraformTemplates")
  for _, templatePath := range box.List() {
    logger.Info(templatePath)
    fullPath := filepath.Join(targetDirectory, templatePath)
    files.EnsurePath(filepath.Dir(fullPath))
    template, _ := box.FindString(templatePath)
    files.CreateFromTemplate(fullPath, template, config)
  }
}

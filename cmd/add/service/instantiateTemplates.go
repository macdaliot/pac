package service

import (
	"path/filepath"

	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(targetDirectory string, config map[string]string) {
	box := packr.NewBox("./serviceTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		fullPath := filepath.Join(targetDirectory, templatePath)
		files.EnsurePath(filepath.Dir(fullPath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(fullPath, template, config)
	}
}

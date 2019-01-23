package frontEndFiles

import (
	"path/filepath"

	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(targetDirectory string, config map[string]string) {
	box := packr.NewBox("./templates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		fullPath := filepath.Join(targetDirectory, templatePath)
		files.EnsurePath(filepath.Dir(fullPath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(fullPath, template, config)
	}
}

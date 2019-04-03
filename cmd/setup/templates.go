package setup

import (
  "os"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func Templates(projectName string, description string, gitAuth string) {
  createRootProjectDirectory(projectName)
  cfg := createConfig(projectName, description, gitAuth)
	createProjectFiles(cfg)
	logger.Info("Created project structure")
}

func createRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	directories.Create(projectDirectory)
  os.Chdir(projectDirectory)
}

func createConfig(projectName string, description string, gitAuth string) map[string]string {
	cfg := make(map[string]string)
	cfg["projectName"] = projectName
	cfg["description"] = description
	cfg["gitAuth"] = gitAuth
  return cfg
}

func createProjectFiles(cfg map[string]string) {
	box := packr.NewBox("./templates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))
		template, _ := box.FindString(templatePath)
		if templatePath != "core/templates/routes-template.ts" {
			files.CreateFromTemplate(templatePath, template, cfg)
		} else {
			files.Write(templatePath, []byte(template))
		}
	}
}

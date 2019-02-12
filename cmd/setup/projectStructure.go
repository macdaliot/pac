package setup

import (
  "os"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func ProjectStructure(projectName string, description string, gitAuth string) {
  createProjectDirectories(projectName)
	createProjectFiles(projectName, description, gitAuth)
	logger.Info("Created project structure")
}

func createProjectDirectories(projectName string) {
	createRootProjectDirectory(projectName)
	directories.Create("app")
	directories.Create("svc")
}

func createRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	directories.Create(projectDirectory)
  os.Chdir(projectDirectory)
}

func createProjectFiles(projectName string, description string, gitAuth string) {
	config := make(map[string]string)
	config["projectName"] = projectName
	config["description"] = description
	config["gitAuth"] = gitAuth
	box := packr.NewBox("./rootTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(templatePath, template, config)
	}
}

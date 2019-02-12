package setup

import (
  "os"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/gobuffalo/packr"
)

func ProjectStructure(projectName string, description string, gitAuth string) string {
	projectDirectory := createProjectDirectories(projectName)
	createProjectFiles(projectDirectory, projectName, description, gitAuth)
	logger.Info("Created project structure")
	return projectDirectory
}

func createProjectDirectories(projectName string) string {
	projectDirectory := createRootProjectDirectory(projectName)
	directories.Create(filepath.Join(projectDirectory, "app"))
	directories.Create(filepath.Join(projectDirectory, "svc"))
	return projectDirectory
}

func createRootProjectDirectory(projectName string) string {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	directories.Create(projectDirectory)
  os.Chdir(projectDirectory)
	return projectDirectory
}

func createProjectFiles(projectDirectory string, projectName string, description string, gitAuth string) {
	config := make(map[string]string)
	config["projectName"] = projectName
	config["description"] = description
	config["gitAuth"] = gitAuth
	box := packr.NewBox("./rootTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		fullPath := filepath.Join(projectDirectory, templatePath)
		files.EnsurePath(filepath.Dir(fullPath))
		template, _ := box.FindString(templatePath)
		files.CreateFromTemplate(fullPath, template, config)
	}
}

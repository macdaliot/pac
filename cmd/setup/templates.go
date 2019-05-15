package setup

import (
	"os"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/gobuffalo/packr"
)

// Templates - creates project directory, config files, and copies project files to project directory.
func Templates() {
	createProjectFiles()
	logger.Info("Created project structure")

	logger.Info("Installing node modules")
	NpmInstall()

	logger.Info("Finished installing node modules")
}

// CreateRootProjectDirectory returns the path to the project root
func CreateRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	directories.Create(projectDirectory)
}

// GoToRootProjectDirectory changes the working directory to the project root
func GoToRootProjectDirectory(projectName string) {
	workingDirectory := directories.GetWorking()
	projectDirectory := filepath.Join(workingDirectory, projectName)
	os.Chdir(projectDirectory)
}

// createProjectFiles "boxes" the template files and replaces Go template strings with the variables stored in the
// .pac.json configuration file
func createProjectFiles() {
	cfg := config.ReadAll()
	box := packr.NewBox("./templates")

	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))

		template, _ := box.FindString(templatePath)

		if templatePath != "core/templates/routes.template.ts" {
			files.CreateFromTemplate(templatePath, template, cfg)
		} else {
			files.Write(templatePath, []byte(template))
		}
	}
}

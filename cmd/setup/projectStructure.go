package setup

import (
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
)

func ProjectStructure(projectName string, description string) string {
  projectDirectory := createProjectDirectories(projectName)
  createProjectFiles(projectDirectory, projectName, description)
  logger.Info("Created project structure")
  return projectDirectory
}

func createProjectDirectories(projectName string) string {
  projectDirectory := createRootProjectDirectory(projectName)
  directories.Create(projectDirectory + "/app/src/assets/png")
  directories.Create(projectDirectory + "/app/src/components/Header")
  directories.Create(projectDirectory + "/app/src/components/Sidebar/parts/Button")
  directories.Create(projectDirectory + "/app/src/components/pages/NotFound")
  directories.Create(projectDirectory + "/app/src/routes")
  directories.Create(projectDirectory + "/app/src/scss")
  directories.Create(projectDirectory + "/svc")
  return projectDirectory
}

func createRootProjectDirectory(projectName string) string {
  workingDirectory := directories.GetWorking()
  projectDirectory := workingDirectory + "/" + projectName
  directories.Create(projectDirectory)
  return projectDirectory
}

func createProjectFiles(projectDirectory string, projectName string, description string) {
  config := make(map[string]string)
  config["projectName"] = projectName
  config["description"] = description
  createReadmeMd(projectDirectory, config)
  createPacFile(projectDirectory, config)
}

func createReadmeMd(projectDirectory string, config map[string]string) {
  const template = `## {{.projectName}}

To get started, try running:

* pac add page --name <page-name> from /app -OR-
* pac add service --name <service-name> from /svc

{{.description}}
`
  files.CreateFromTemplate(projectDirectory + "/README.md", template, config)
}

func createPacFile(projectDirectory string, config map[string]string) {
  const template = `{
  "projectName": "{{.projectName}}"
}
`
  files.CreateFromTemplate(projectDirectory + "/.pac", template, config)
}

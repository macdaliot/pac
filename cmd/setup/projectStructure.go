package setup

import (
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
)

func ProjectStructure(projectName string, description string, gitAuth string) string {
  projectDirectory := createProjectDirectories(projectName)
  createProjectFiles(projectDirectory, projectName, description, gitAuth)
  logger.Info("Created project structure")
  return projectDirectory
}

func createProjectDirectories(projectName string) string {
  projectDirectory := createRootProjectDirectory(projectName)
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

func createProjectFiles(projectDirectory string, projectName string, description string, gitAuth string) {
  config := make(map[string]string)
  config["projectName"] = projectName
  config["description"] = description
  config["gitAuth"] = gitAuth
  createGitIgnore(projectDirectory)
  createReadmeMd(projectDirectory, config)
  createPacFile(projectDirectory, config)
}

func createGitIgnore(projectDirectory string) {
  const template = `svc/*/node_modules
svc/*/server.js
`
  files.CreateFromTemplate(projectDirectory + "/.gitignore", template, nil)
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
  "projectName": "{{.projectName}}",
  "gitAuth": "{{.gitAuth}}"
}
`
  files.CreateFromTemplate(projectDirectory + "/.pac", template, config)
}

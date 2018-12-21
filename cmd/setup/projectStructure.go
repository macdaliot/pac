package setup

import (
  "encoding/base64"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
)

func ProjectStructure(projectName string, description string, gitUser string, gitPass string) string {
  projectDirectory := createProjectDirectories(projectName)
  createProjectFiles(projectDirectory, projectName, description, gitUser, gitPass)
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

func createProjectFiles(projectDirectory string, projectName string, description string, gitUser string, gitPass string) {
  config := make(map[string]string)
  config["projectName"] = projectName
  config["description"] = description
  config["gitBasicAuth"] = base64.StdEncoding.EncodeToString([]byte(gitUser + ":" + gitPass))
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
  "projectName": "{{.projectName}}",
  "gitAuth": "{{.gitBasicAuth}}"
}
`
  files.CreateFromTemplate(projectDirectory + "/.pac", template, config)
}

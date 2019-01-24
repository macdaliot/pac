package setup

import (
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
)

func ProjectStructure(projectName string, description string, gitAuth string) string {
	projectDirectory := createProjectDirectories(projectName)
	createProjectFiles(projectDirectory, projectName, description, gitAuth)
	logger.Info("Created project structure")
	return projectDirectory
}

func createProjectDirectories(projectName string) string {
  projectDirectory := createRootProjectDirectory(projectName)
  directories.Create(str.Concat(projectDirectory, "/app/src/components/Header"))
  directories.Create(str.Concat(projectDirectory, "/app/src/components/Sidebar/parts/Button"))
  directories.Create(str.Concat(projectDirectory, "/app/src/components/pages/NotFound"))
  directories.Create(str.Concat(projectDirectory, "/app/src/routes"))
  directories.Create(str.Concat(projectDirectory, "/app/src/services"))
  directories.Create(str.Concat(projectDirectory, "/app/src/scss"))
  directories.Create(str.Concat(projectDirectory, "/svc"))
  return projectDirectory
}

func createRootProjectDirectory(projectName string) string {
	workingDirectory := directories.GetWorking()
	projectDirectory := str.Concat(workingDirectory, "/", projectName)
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
	const template = `app/node_modules
svc/*/node_modules
svc/*/dist
svc/*/.env
`
	files.CreateFromTemplate(str.Concat(projectDirectory, "/.gitignore"), template, nil)
}

func createReadmeMd(projectDirectory string, config map[string]string) {
	const template = `## {{.projectName}}

### Getting Started
1. Navigate to your Jenkins' IP address. Wait until Jenkins has a "Credentials" link on the left sidebar. When the link appears, run from your project's root directory:

* pac automate jenkins

2. Create one or more microservices with:

* cd svc
* pac add service --name <service-name>

3. Push to GitHub:

* git add --all
* git commit -m "Initial commit"
* git push origin master

4. Manually run each newly created microservice pipeline from the Jenkins GUI (only needed on the first run)

The above workflow will set up your project and deploys it to AWS

{{.description}}
`
	files.CreateFromTemplate(str.Concat(projectDirectory, "/README.md"), template, config)
}

func createPacFile(projectDirectory string, config map[string]string) {
	const template = `{
  "projectName": "{{.projectName}}",
  "gitAuth": "{{.gitAuth}}"
}
`
	files.CreateFromTemplate(str.Concat(projectDirectory, "/.pac"), template, config)
}

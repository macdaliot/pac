package setup

import (
  "runtime"
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
  createCoverageSh(projectDirectory, projectName)
  createJenkinsfile(projectDirectory, config)
}

func createGitIgnore(projectDirectory string) {
	const template = `app/node_modules
svc/*/node_modules
svc/*/dist
svc/**/.env
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

The above workflow will set up your project and deploys it to AWS

### Helpful Links
* [http://integration.{{.projectName}}.pac.pyramidchallenges.com](http://integration.{{.projectName}}.pac.pyramidchallenges.com) (takes ~15 minutes to provision)
* [http://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>](http://api.{{.projectName}}.pac.pyramidchallenges.com/api/<service-name>)
* [http://jenkins.{{.projectName}}.pac.pyramidchallenges.com:8080](http://jenkins.{{.projectName}}.pac.pyramidchallenges.com:8080)
* [http://sonarqube.{{.projectName}}.pac.pyramidchallenges.com:9000](http://sonarqube.{{.projectName}}.pac.pyramidchallenges.com:9000)
* [http://selenium.{{.projectName}}.pac.pyramidchallenges.com:4444](http://selenium.{{.projectName}}.pac.pyramidchallenges.com:4444)

### Description
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

func createCoverageSh(projectDirectory string, projectName string) {
	const template = `#!/bin/sh
TARGET=$1

echo $TARGET
cd $TARGET
ls -la
npm i
npm run-script test-coverage
`
	files.CreateFromTemplate(str.Concat(projectDirectory, "/coverage.sh"), template, nil)
	if runtime.GOOS == "windows" {
		files.ChangePermissions(str.Concat(".\\", projectName, "\\coverage.sh"), 0755)
	} else {
		files.ChangePermissions(str.Concat("./", projectName, "/coverage.sh"), 0755)
	}
}

func createJenkinsfile(projectDirectory string, config map[string]string) {
	const template = `def pipelineComponents
withCredentials([string(credentialsId: 'PipelineComponents', variable: 'COMPONENTS')]) {
    pipelineComponents = env.COMPONENTS.split(",")
}

def builds = [:]
def unitTests = [:]
def integrationTests = [:]
def deployments = [:]

for (x in pipelineComponents){
    def jobName = x.trim()
    // create builds
    builds[jobName] = {
        build job: jobName,
            parameters: [
                    booleanParam(name: 'EXEC_DEPLOY', value: false),
                    booleanParam(name: 'EXEC_INTEGRATIONTEST', value: false),
                    booleanParam(name: 'EXEC_UNITTEST', value: false)
                ]
    }
    // create unit test runs
    unitTests[jobName] = {
        build job: jobName,
            parameters: [
                    booleanParam(name: 'EXEC_DEPLOY', value: false),
                    booleanParam(name: 'EXEC_INTEGRATIONTEST', value: false)
                ]
    }
    // create integration test runs
    integrationTests[jobName] = {
        build job: jobName,
            parameters: [
                    booleanParam(name: 'EXEC_BUILD', value: false),
                    booleanParam(name: 'EXEC_DEPLOY', value: false),
                    booleanParam(name: 'EXEC_UNITTEST', value: false)
                ]
    }
    // create deployments
    deployments[jobName] = {
        build job: jobName,
            parameters: [
                    booleanParam(name: 'EXEC_INTEGRATIONTEST', value: false),
                    booleanParam(name: 'EXEC_UNITTEST', value: false)
                ]
    }
}

node {
    stage("Git Checkout") {
        checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [[$class: 'CleanBeforeCheckout']], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'gitcredentials', url: 'http://github.com/PyramidSystemsInc/{{.projectName}}.git']]])
    }
    stage("Build") {
        parallel builds
    }
    stage("Unit Tests") {
        parallel unitTests
    }
    stage("Coverage") {
        script {
            // scan services for test coverage
            sh 'find svc -type d -maxdepth 1 -mindepth 1 -exec sh coverage.sh {} \\;'
            sh 'rm -rf coverage ; mkdir coverage ; cd svc ; npm i ; npx lcov-result-merger "**/lcov.info" ../coverage/all.info'
        }
    }
    stage("Deploy"){
        parallel deployments
    }
    stage("Inspect") {
        withCredentials([string(credentialsId: 'SONARLOGIN', variable: 'SONARLOGIN')]) {
            script {
                sh 'npm i typescript'
                sh 'sonar-scanner \
                    -Dsonar.projectKey=app \
                    -Dsonar.sources=. \
                    -Dsonar.host.url=http://sonarqube.{{.projectName}}.pac.pyramidchallenges.com:9000 \
                    -Dsonar.login=$SONARLOGIN'
            }
        }
    }
    stage("Integration Tests"){
        parallel integrationTests
    }
}
`
	files.CreateFromTemplate(str.Concat(projectDirectory, "/Jenkinsfile"), template, config)
}

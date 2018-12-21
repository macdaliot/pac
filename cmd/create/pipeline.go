package create

import (
  "encoding/base64"
  "os"
  "strings"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
)

func Pipeline(pipelineName string, branches string, description string) {
  // TODO: Allow multiple branches
  pipelineDir := "./jenkins/" + pipelineName
  pacFile := readPacFile()
  jenkinsUrl := "http://" + pacFile.JenkinsUrl
  config := createPipelineConfig(pipelineName, pacFile.ProjectName, branches, description)
  createJobXml(pipelineDir, config)
  downloadJenkinsCliJar(jenkinsUrl)
  jenkinsCliCommandStart := "java -jar jenkins-cli.jar -s " + jenkinsUrl + " -auth pyramid:systems"
  saveGitCredentialsToJenkinsIfDoesNotExist(pacFile, jenkinsCliCommandStart)
  createJenkinsPipelineJob(pipelineName, pipelineDir, jenkinsCliCommandStart)
  // performInitialBuildOfNewJob(pipelineName, jenkinsCliCommandStart)
  cleanUp()
  logger.Info("Created " + pipelineName + " pipeline job in Jenkins")
}

func createPipelineConfig(pipelineName string, projectName string, branches string, description string) map[string]string {
  config := make(map[string]string)
  config["pipelineName"] = pipelineName
  config["projectName"] = projectName
  config["branches"] = branches
  config["description"] = description
  return config
}

func createJobXml(pipelineDir string, config map[string]string) {
  const template = `<?xml version='1.1' encoding='UTF-8'?>
<flow-definition plugin="workflow-job@2.31">
  <description>{{.description}}</description>
  <keepDependencies>false</keepDependencies>
  <properties>
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.29.3">
      <projectUrl>https://github.com/PyramidSystemsInc/{{.projectName}}/</projectUrl>
      <displayName></displayName>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
    <org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
      <triggers>
        <com.cloudbees.jenkins.GitHubPushTrigger plugin="github@1.29.3">
          <spec></spec>
        </com.cloudbees.jenkins.GitHubPushTrigger>
      </triggers>
    </org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.61">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.9.1">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <url>https://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
          <credentialsId>gitcredentials</credentialsId>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>*/{{.branches}}</name>
        </hudson.plugins.git.BranchSpec>
      </branches>
      <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
      <submoduleCfg class="list"/>
      <extensions/>
    </scm>
    <scriptPath>jenkins/{{.pipelineName}}/Jenkinsfile</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>
`
  files.CreateFromTemplate(pipelineDir + "/job.xml", template, config)
}

func downloadJenkinsCliJar(jenkinsUrl string) {
  commands.Run("wget " + jenkinsUrl + "/jnlpJars/jenkins-cli.jar", "")
}

func createCredentialsXml(pacFile PacFile) {
  config := createGitCredentialsConfig(pacFile)
  const template = `<com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl plugin="credentials@2.1.18">
  <scope>GLOBAL</scope>
  <id>gitcredentials</id>
  <description></description>
  <username>{{.gitUser}}</username>
  <password>{{.gitPass}}</password>
</com.cloudbees.plugins.credentials.impl.UsernamePasswordCredentialsImpl>
`
  files.CreateFromTemplate(directories.GetWorking() + "/credentials.xml", template, config)
}

func createGitCredentialsConfig(pacFile PacFile) map[string]string {
  decoded, err := base64.StdEncoding.DecodeString(pacFile.GitAuth)
  errors.LogIfError(err)
  separatorIndex := strings.Index(string(decoded), ":")
  config := make(map[string]string)
  config["gitUser"] = string(decoded)[0:separatorIndex]
  config["gitPass"] = string(decoded)[separatorIndex + 1:len(string(decoded))]
  return config
}

func saveGitCredentialsToJenkinsIfDoesNotExist(pacFile PacFile, jenkinsCliCommandStart string) {
  if !doJenkinsCredentialsExist(jenkinsCliCommandStart) {
    createCredentialsXml(pacFile)
    credentialData := files.Read("./credentials.xml")
    commands.RunWithStdin(jenkinsCliCommandStart + " create-credentials-by-xml system::system::jenkins (global)", string(credentialData), "")
  }
}

func doJenkinsCredentialsExist(jenkinsCliCommandStart string) bool {
  results := commands.Run(jenkinsCliCommandStart + " list-credentials system::system::jenkins", "")
  return strings.Contains(results, "gitcredentials")
}

func createJenkinsPipelineJob(pipelineName string, pipelineDir string, jenkinsCliCommandStart string) {
  jobData := files.Read(pipelineDir + "/job.xml")
  createJobCommand := jenkinsCliCommandStart + " create-job " + pipelineName
  commands.RunWithStdin(createJobCommand, string(jobData), "")
}

func performInitialBuildOfNewJob(pipelineName string, jenkinsCliCommandStart string) {
  buildJobCommand := jenkinsCliCommandStart + " build " + pipelineName
  commands.Run(buildJobCommand, "")
}

func cleanUp() {
  if _, err := os.Stat("credentials.xml"); !os.IsNotExist(err) {
    commands.Run("rm credentials.xml", "")
  }
  commands.Run("rm jenkins-cli.jar", "")
}

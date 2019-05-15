package setup

import (
	"encoding/json"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"strings"
)

func AutomateAzure() {
	//TODO Retrieve azureDevopsOrg/projectName from pac file
	//58d2631a7a7a5af5e7d4438afa4077dea8ffc60a  joe token
	projectName := "CLIPROJECT"//config.Get("projectName")
	azureDevopsOrg := "https://dev.azure.com/jgonzalezdevops/"//config.Get("azureDevopsOrg")
	repoUrl := "https://github.com/PyramidSystemsInc/AzurePipeline/"
	accessToken := "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"


	login()
	//createDevopsProject(projectName, azureDevopsOrg)
	endpointId := createAndRetrieveServiceEndpoint(projectName, accessToken, repoUrl, azureDevopsOrg)
	createPipeline(projectName, repoUrl, azureDevopsOrg, endpointId)

	//createEntrypointJobXml(projectName)
	//createMasterPipelineXml(projectName)
	//createFrontEndPipelineXml(projectName)
	//createServicesPipelineXml(projectName)
	//createPipelineJobs(jenkinsURL, projectName, jenkinsCliCommandStart)
	logger.Info("Azure DevOps should now be completely configured.")
}

func createPipeline(projName string, repoUrl string, devopsOrg string, endpointConnection string) {
	output,err := commands.Run(buildCreatePipelineCommand(projName, repoUrl, devopsOrg, endpointConnection), "")
	logger.Info(output)
	if err != nil {
		logger.Err(err.Error())
	}
}

func buildCreatePipelineCommand(projName string, repoUrl string, devopsOrg string,  endpointConnection string) string {
	//az pipelines create --name 'ClIPipelin662' --description 'Pipeline for template CLI project' --repository https://github.com/PyramidSystemsInc/AzurePipeline/ --branch master --yml-path pipeline-definition.yml --service-connection ec8529d1-3d07-4189-8690-fdd2a15a28d8  --organization https://dev.azure.com/jgonzalezdevops --project CLIProject

	var sb strings.Builder
	sb.WriteString("az pipelines create --branch master --yml-path pipeline-definition.yml --repository-type github --name ")
	sb.WriteString("joeTest")
	sb.WriteString("Pipeline --repository ")
	sb.WriteString(repoUrl)
	sb.WriteString(" --organization ")
	sb.WriteString(devopsOrg)
	sb.WriteString(" --project ")
	sb.WriteString(projName)
	sb.WriteString(" --service-connection ")
	sb.WriteString(endpointConnection)
	logger.Info(sb.String())
	return sb.String()
}

func login() {
	//TODO swap out cred for a key
	s, e := commands.Run("az login -u jgonzalez@psi-it.com -p Lastthe.88", "")
	logger.Info(s)
	if e != nil {
		logger.Err(e.Error())
	}
}

func createAndRetrieveServiceEndpoint(projName string, accessToken string, repoUrl string, devopsOrg string) string {
	output,err := commands.Run(buildCreateServiceEnpointCommand(projName, accessToken, repoUrl, devopsOrg), "")
	logger.Info(output)
	if err != nil {
		logger.Err(err.Error())
	}
	result := make(map[string]string)
	json.Unmarshal([]byte(output), &result)
	return result["id"]
}

func buildCreateServiceEnpointCommand(projName string, accessToken string, repoUrl string, devopsOrg string) string  {
	var sb strings.Builder
	sb.WriteString("az devops service-endpoint create --authorization-scheme PersonalAccessToken --service-endpoint-type github --github-access-token \"")
	sb.WriteString(accessToken)
	sb.WriteString("\" --github-url ")
	sb.WriteString(repoUrl)
	sb.WriteString(" --org ")
	sb.WriteString(devopsOrg)
	sb.WriteString(" --project ")
	sb.WriteString(projName)
	sb.WriteString(" --name ")
	sb.WriteString(projName)
	sb.WriteString("Endpoint")
	return sb.String()
}

func createDevopsProject(projectName string, azureDevopsOrg string) {
	out,err := commands.Run(buildDevopsProjectCreateCommand(projectName, azureDevopsOrg), "")
	//TODO handle correct logging/error handling
	logger.Info(out)
	if err != nil {
		logger.Err(err.Error())
	}
}

func buildDevopsProjectCreateCommand(projectName string, azureDevopsOrg string) string {
	var sb strings.Builder
	sb.WriteString("az devops project create --name ")
	sb.WriteString(projectName)
	sb.WriteString(" --org ")
	sb.WriteString(azureDevopsOrg)
	sb.WriteString("\" --source-control git --visibility private")
	return sb.String()
}

func createMasterPipelineYml(projectName string) {
	filePath := "master-pipeline.xml"
	config := make(map[string]string)
	config["projectName"] = projectName
	const template = `<?xml version="1.1" encoding="UTF-8"?><flow-definition plugin="workflow-job@2.31">
  <description/>
  <keepDependencies>false</keepDependencies>
  <properties>
    <jenkins.model.BuildDiscarderProperty>
      <strategy class="hudson.tasks.LogRotator">
        <daysToKeep>-1</daysToKeep>
        <numToKeep>10</numToKeep>
        <artifactDaysToKeep>-1</artifactDaysToKeep>
        <artifactNumToKeep>-1</artifactNumToKeep>
      </strategy>
    </jenkins.model.BuildDiscarderProperty>
    <org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty/>
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.29.3">
      <projectUrl>http://github.com/PyramidSystemsInc/{{.projectName}}/</projectUrl>
      <displayName/>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.61">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.9.1">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <url>http://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
          <credentialsId>gitcredentials</credentialsId>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>*/master</name>
        </hudson.plugins.git.BranchSpec>
      </branches>
      <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
      <submoduleCfg class="list"/>
      <extensions/>
    </scm>
    <scriptPath>Jenkinsfile</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>
`
	files.CreateFromTemplate(filePath, template, config)
}


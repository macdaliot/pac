package setup

import (
	"time"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
)

const projName string = "projectName"

// TODO: Needs to be cleaned up and made into a packr box as a separate task
// TODO: Should either be automated as a part of `pac setup` or moved to `cmd/automate/jenkins.go`
func AutomateJenkins() {
	projectName := config.Get(projName)
	jenkinsURL := config.Get("jenkinsURL")
	jenkinsCliCommandStart := str.Concat("java -jar jenkins-cli.jar -s https://", jenkinsURL, " -auth pyramid:systems")
	DownloadJenkinsCliJar(jenkinsURL)
	createEntrypointJobXml(projectName)
	createMasterPipelineXml(projectName)
	createFrontEndPipelineXml(projectName)
	createServicesPipelineXml(projectName)
	createProductionPipelineXml(projectName)
	createPipelineJobs(jenkinsURL, projectName, jenkinsCliCommandStart)
	cleanUp()
	logger.Info("Jenkins should now be completely configured.")
}

func DownloadJenkinsCliJar(jenkinsURL string) {
	corruptJenkinsCliError := "Error: Invalid or corrupt jarfile jenkins-cli.jar"
	jenkinsCliPath := "./jenkins-cli.jar"
	err := files.Download(str.Concat("https://", jenkinsURL, "/jnlpJars/jenkins-cli.jar"), jenkinsCliPath)
	if err != nil {
		if err.Error() == corruptJenkinsCliError {
			files.Delete(jenkinsCliPath)
			logger.Info("Received a corrupt Jenkins CLI. Attempting another download")
			time.Sleep(20 * time.Second)
			DownloadJenkinsCliJar(jenkinsURL)
		} else {
			errors.LogAndQuit(str.Concat("Downloading the Jenkins CLI failed with the following error: ", err.Error()))
		}
	}
}

func createEntrypointJobXml(projectName string) {
	filePath := "entrypoint-job.xml"
	config := make(map[string]string)
	config[projName] = projectName
	const template = `<?xml version='1.1' encoding='UTF-8'?>
<project>
  <actions/>
  <description></description>
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
  </properties>
  <scm class="hudson.plugins.git.GitSCM" plugin="git@4.0.0-rc">
    <configVersion>2</configVersion>
    <userRemoteConfigs>
      <hudson.plugins.git.UserRemoteConfig>
        <url>http://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
        <credentialsId>gitcredentials</credentialsId>
      </hudson.plugins.git.UserRemoteConfig>
    </userRemoteConfigs>
    <branches>
      <hudson.plugins.git.BranchSpec>
        <name>master</name>
      </hudson.plugins.git.BranchSpec>
    </branches>
    <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
    <submoduleCfg class="list"/>
    <extensions/>
  </scm>
  <canRoam>true</canRoam>
  <disabled>false</disabled>
  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>
  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>
  <triggers>
    <com.cloudbees.jenkins.GitHubPushTrigger plugin="github@1.29.3">
      <spec></spec>
    </com.cloudbees.jenkins.GitHubPushTrigger>
  </triggers>
  <concurrentBuild>false</concurrentBuild>
  <builders>
    <hudson.tasks.Shell>
      <command>java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems build release-through-staging</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers/>
  <buildWrappers />
</project>
`
	files.CreateFromTemplate(filePath, template, config)
}

func createServicesPipelineXml(projectName string) {
	filePath := "services-pipeline.xml"
	config := make(map[string]string)
	config[projName] = projectName
	const template = `<?xml version="1.1" encoding="UTF-8"?><flow-definition plugin="workflow-job@2.31">
  <description/>
  <keepDependencies>false</keepDependencies>
  <properties>
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
    <scriptPath>services/Jenkinsfile</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>
`
	files.CreateFromTemplate(filePath, template, config)
}

func createFrontEndPipelineXml(projectName string) {
	filePath := "front-end-pipeline.xml"
	config := make(map[string]string)
	config[projName] = projectName
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
    <scriptPath>app/Jenkinsfile</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>
`
	files.CreateFromTemplate(filePath, template, config)
}

func createMasterPipelineXml(projectName string) {
	const template = `<?xml version='1.1' encoding='UTF-8'?>
<flow-definition plugin="workflow-job@2.32">
  <actions/>
  <description>For deploying to the development and staging environments</description>
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
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.29.4">
      <projectUrl>http://github.com/PyramidSystemsInc/{{.projectName}}</projectUrl>
      <displayName></displayName>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
    <hudson.model.ParametersDefinitionProperty>
      <parameterDefinitions>
        <hudson.model.StringParameterDefinition>
          <name>GIT_COMMIT</name>
          <description>An optional commit hash used to deploy from a specific commit.</description>
          <defaultValue>master</defaultValue>
          <trim>true</trim>
        </hudson.model.StringParameterDefinition>
      </parameterDefinitions>
    </hudson.model.ParametersDefinitionProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.70">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.10.0">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <url>http://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
          <credentialsId>gitcredentials</credentialsId>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>master</name>
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
	cfg := make(map[string]string)
	cfg[projName] = projectName
	files.CreateFromTemplate("master.xml", template, cfg)
}


func createProductionPipelineXml(projectName string) {
 const template = `<?xml version='1.1' encoding='UTF-8'?>
<flow-definition plugin="workflow-job@2.33">
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties>
    <hudson.model.ParametersDefinitionProperty>
      <parameterDefinitions>
        <hudson.model.StringParameterDefinition>
          <name>BUILD_NUMBER</name>
          <description>Build number from the master pipeline that produced a release candidate</description>
          <defaultValue></defaultValue>
          <trim>true</trim>
        </hudson.model.StringParameterDefinition>
      </parameterDefinitions>
    </hudson.model.ParametersDefinitionProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.70">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@3.10.0">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <url>https://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
          <credentialsId>gitcredentials</credentialsId>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>master</name>
        </hudson.plugins.git.BranchSpec>
      </branches>
      <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
      <submoduleCfg class="list"/>
      <extensions/>
    </scm>
    <scriptPath>Jenkinsfile-Prod</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>`
	cfg := make(map[string]string)
	cfg["projectName"] = projectName
	files.CreateFromTemplate("prod.xml", template, cfg)
}

func createPipelineJobs(jenkinsURL string, projectName string, jenkinsCliCommandStart string) {
	jobData := files.Read("entrypoint-job.xml")
	createJobCommand := str.Concat(jenkinsCliCommandStart, " create-job entrypoint")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
	jobData = files.Read("services-pipeline.xml")
	createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job services")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
	jobData = files.Read("front-end-pipeline.xml")
	createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job front-end")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
	jobData = files.Read("master.xml")
	createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job release-through-staging")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
	jobData = files.Read("prod.xml")
	createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job release-to-production")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
}

func cleanUp() {
	// TODO: Change to os.Rm() or something in order to support Windows CMD
	commands.Run("rm jenkins-cli.jar entrypoint-job.xml front-end-pipeline.xml services-pipeline.xml master.xml prod.xml", "")
}

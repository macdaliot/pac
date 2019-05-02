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

func AutomateJenkins() {
	time.Sleep(60 * time.Second)
	projectName := config.Get("projectName")
	jenkinsURL := config.Get("jenkinsURL")
	DownloadJenkinsCliJar(jenkinsURL)
	createEntrypointJobXml(projectName)
	createMasterPipelineXml(projectName)
	createFrontEndPipelineXml(projectName)
	createServicesPipelineXml(projectName)
	jenkinsCliCommandStart := str.Concat("java -jar jenkins-cli.jar -s http://", jenkinsURL, " -auth pyramid:systems")
	createPipelineJobs(jenkinsURL, projectName, jenkinsCliCommandStart)
	cleanUp()
	logger.Info("Jenkins should now be completely configured.")
}

func DownloadJenkinsCliJar(jenkinsURL string) {
	corruptJenkinsCliError := "Error: Invalid or corrupt jarfile jenkins-cli.jar"
	jenkinsCliPath := "./jenkins-cli.jar"
	err := files.Download(str.Concat("http://", jenkinsURL, "/jnlpJars/jenkins-cli.jar"), jenkinsCliPath)
	if err != nil {
		logger.Info("Made it in the if")
		if err.Error() == corruptJenkinsCliError {
			logger.Info("Made it in the nested if")
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
	config["projectName"] = projectName
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
        <name>*/master</name>
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
      <command>java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems build master</command>
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
	config["projectName"] = projectName
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
	jobData = files.Read("master-pipeline.xml")
	createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job master")
	commands.RunWithStdin(createJobCommand, string(jobData), "")
}

func cleanUp() {
	// TODO: Change to os.Rm() or something in order to support Windows CMD
	commands.Run("rm jenkins-cli.jar entrypoint-job.xml front-end-pipeline.xml services-pipeline.xml master-pipeline.xml", "")
}

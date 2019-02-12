package automate

import (
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

func Jenkins() {
  projectName := config.Get("projectName")
  jenkinsUrl := config.Get("jenkinsUrl")
  downloadJenkinsCliJar(jenkinsUrl)
  createPipelineProvisionerXml(projectName)
  createS3PipelineXml(projectName)
  createWholePipelineXml(projectName)
  jenkinsCliCommandStart := str.Concat("java -jar jenkins-cli.jar -s http://", jenkinsUrl, " -auth pyramid:systems")
  createPipelineJobs(jenkinsUrl, projectName, jenkinsCliCommandStart)
  createPipelineComponentsSecret(jenkinsUrl, jenkinsCliCommandStart)
  cleanUp()
  logger.Info("Jenkins is now configured to create individual pipelines for the front-end and each microservice")
}

func downloadJenkinsCliJar(jenkinsUrl string) {
  files.Download(str.Concat("http://", jenkinsUrl, "/jnlpJars/jenkins-cli.jar"), "./jenkins-cli.jar")
}

func createPipelineProvisionerXml(projectName string) {
  filePath := "pipeline-provisioner-job.xml"
  config := make(map[string]string)
  config["projectName"] = projectName
  const template = `<?xml version='1.1' encoding='UTF-8'?>
<project>
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties/>
  <scm class="hudson.plugins.git.GitSCM" plugin="git@4.0.0-rc">
    <configVersion>2</configVersion>
    <userRemoteConfigs>
      <hudson.plugins.git.UserRemoteConfig>
        <url>https://github.com/PyramidSystemsInc/{{.projectName}}.git</url>
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
      <command>if [ -d ./svc ]; then
  cd svc
  NEW_PIPELINES=&quot;&quot;
  for DIR in *; do
    if [ -d &quot;${DIR}&quot; ]; then
      if java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems get-job &quot;${DIR}&quot;; then
        echo &quot;${DIR} Jenkins Pipeline Exists. Skipping&quot;
      else
        PROJECT_NAME=$(sed -e &apos;s/.*\///g&apos; -e &apos;s/.git$//g&apos; &lt;&lt;&lt; $(echo &quot;$GIT_URL&quot;))
        if [ -z $NEW_PIPELINES ]; then
          NEW_PIPELINES=${DIR}
        else
          NEW_PIPELINES=$NEW_PIPELINES,${DIR}
        fi
cat &lt;&lt;- EOF &gt; job.xml
&lt;?xml version=&apos;1.1&apos; encoding=&apos;UTF-8&apos;?&gt;
&lt;flow-definition plugin=&quot;workflow-job@2.31&quot;&gt;
  &lt;description&gt;&lt;/description&gt;
  &lt;keepDependencies&gt;false&lt;/keepDependencies&gt;
  &lt;properties&gt;
    &lt;com.coravy.hudson.plugins.github.GithubProjectProperty plugin=&quot;github@1.29.3&quot;&gt;
      &lt;projectUrl&gt;https://github.com/PyramidSystemsInc/$PROJECT_NAME/&lt;/projectUrl&gt;
      &lt;displayName&gt;&lt;/displayName&gt;
    &lt;/com.coravy.hudson.plugins.github.GithubProjectProperty&gt;
  &lt;/properties&gt;
  &lt;definition class=&quot;org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition&quot; plugin=&quot;workflow-cps@2.61&quot;&gt;
    &lt;scm class=&quot;hudson.plugins.git.GitSCM&quot; plugin=&quot;git@3.9.1&quot;&gt;
      &lt;configVersion&gt;2&lt;/configVersion&gt;
      &lt;userRemoteConfigs&gt;
        &lt;hudson.plugins.git.UserRemoteConfig&gt;
          &lt;url&gt;https://github.com/PyramidSystemsInc/$PROJECT_NAME.git&lt;/url&gt;
          &lt;credentialsId&gt;gitcredentials&lt;/credentialsId&gt;
        &lt;/hudson.plugins.git.UserRemoteConfig&gt;
      &lt;/userRemoteConfigs&gt;
      &lt;branches&gt;
        &lt;hudson.plugins.git.BranchSpec&gt;
          &lt;name&gt;*/master&lt;/name&gt;
        &lt;/hudson.plugins.git.BranchSpec&gt;
      &lt;/branches&gt;
      &lt;doGenerateSubmoduleConfigurations&gt;false&lt;/doGenerateSubmoduleConfigurations&gt;
      &lt;submoduleCfg class=&quot;list&quot;/&gt;
      &lt;extensions/&gt;
    &lt;/scm&gt;
    &lt;scriptPath&gt;svc/${DIR}/Jenkinsfile&lt;/scriptPath&gt;
    &lt;lightweight&gt;true&lt;/lightweight&gt;
  &lt;/definition&gt;
  &lt;triggers/&gt;
  &lt;disabled&gt;false&lt;/disabled&gt;
&lt;/flow-definition&gt;
EOF
        java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems create-job &quot;${DIR}&quot; &lt; job.xml
      fi
    fi
  done
fi
if [ -z $NEW_PIPELINES ]; then
  echo &quot;No new pipeline components to trigger&quot;
else
cat &lt;&lt;- EOF &gt; pipeline-components.xml
&lt;org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl plugin=&quot;plain-credentials@1.5&quot;&gt;
  &lt;scope&gt;GLOBAL&lt;/scope&gt;
  &lt;id&gt;PipelineComponents&lt;/id&gt;
  &lt;description&gt;PipelineComponents&lt;/description&gt;
  &lt;secret&gt;$PipelineComponents,$NEW_PIPELINES&lt;/secret&gt;
&lt;/org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl&gt;
EOF
  java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems update-credentials-by-xml system::system::jenkins \(global\) PipelineComponents &lt; pipeline-components.xml
fi
java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems build {{.projectName}}</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers/>
  <buildWrappers>
    <org.jenkinsci.plugins.credentialsbinding.impl.SecretBuildWrapper plugin="credentials-binding@1.17">
      <bindings>
        <org.jenkinsci.plugins.credentialsbinding.impl.StringBinding>
          <credentialsId>PipelineComponents</credentialsId>
          <variable>PipelineComponents</variable>
        </org.jenkinsci.plugins.credentialsbinding.impl.StringBinding>
      </bindings>
    </org.jenkinsci.plugins.credentialsbinding.impl.SecretBuildWrapper>
  </buildWrappers>
</project>
`
  files.CreateFromTemplate(filePath, template, config)
}

func createS3PipelineXml(projectName string) {
  filePath := "s3-pipeline-job.xml"
  config := make(map[string]string)
  config["projectName"] = projectName
  const template = `<?xml version="1.1" encoding="UTF-8"?><flow-definition plugin="workflow-job@2.31">
  <description/>
  <keepDependencies>false</keepDependencies>
  <properties>
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.29.3">
      <projectUrl>https://github.com/PyramidSystemsInc/{{.projectName}}/</projectUrl>
      <displayName/>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
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

func createWholePipelineXml(projectName string) {
  filePath := "whole-pipeline-job.xml"
  config := make(map[string]string)
  config["projectName"] = projectName
  const template = `<?xml version="1.1" encoding="UTF-8"?><flow-definition plugin="workflow-job@2.31">
  <description/>
  <keepDependencies>false</keepDependencies>
  <properties>
    <com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.29.3">
      <projectUrl>https://github.com/PyramidSystemsInc/{{.projectName}}/</projectUrl>
      <displayName/>
    </com.coravy.hudson.plugins.github.GithubProjectProperty>
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

func createPipelineJobs(jenkinsUrl string, projectName string, jenkinsCliCommandStart string) {
  jobData := files.Read("pipeline-provisioner-job.xml")
  createJobCommand := str.Concat(jenkinsCliCommandStart, " create-job pipeline-provisioner")
  commands.RunWithStdin(createJobCommand, string(jobData), "")
  jobData = files.Read("s3-pipeline-job.xml")
  createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job front-end-integration")
  commands.RunWithStdin(createJobCommand, string(jobData), "")
  jobData = files.Read("whole-pipeline-job.xml")
  createJobCommand = str.Concat(jenkinsCliCommandStart, " create-job ", projectName)
  commands.RunWithStdin(createJobCommand, string(jobData), "")
}

func createPipelineComponentsSecret(jenkinsUrl string, jenkinsCliCommandStart string) {
  createCredentialsCommand := str.Concat(jenkinsCliCommandStart, " create-credentials-by-xml system::system::jenkins (global)")
  pipelineComponentsXml := []byte(`<org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl plugin="plain-credentials@1.5">
  <scope>GLOBAL</scope>
  <id>PipelineComponents</id>
  <description>PipelineComponents</description>
  <secret>front-end-integration</secret>
</org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl>
`)
  commands.RunWithStdin(createCredentialsCommand, string(pipelineComponentsXml), "")
}

func cleanUp() {
  // TODO: Change to os.Rm() or something in order to support Windows CMD
  commands.Run("rm jenkins-cli.jar pipeline-provisioner-job.xml s3-pipeline-job.xml whole-pipeline-job.xml", "")
}

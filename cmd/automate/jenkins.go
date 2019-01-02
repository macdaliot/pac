package automate

import (
  "encoding/json"
  "io/ioutil"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/files"
)

type PacFile struct {
  ProjectName      string  `json:"projectName"`
  GitAuth          string  `json:"gitAuth"`
  JenkinsUrl       string  `json:"jenkinsUrl"`
  LoadBalancerArn  string  `json:"loadBalancerArn"`
  ListenerArn      string  `json:"listenerArn"`
  ServiceUrl       string  `json:"serviceUrl"`
}

func Jenkins() {
  pacFile := readPacFile()
  jenkinsUrl := pacFile.JenkinsUrl
  downloadJenkinsCliJar(jenkinsUrl)
  createPipelineProvisionerXml(pacFile.ProjectName, jenkinsUrl)
  createPipelineProvisionerJob(jenkinsUrl)
  cleanUp()
}

func readPacFile() PacFile {
  // TODO: Should run from anywhere
  // TODO: Should not depend on pacFile for git
  var pacFile PacFile
  pacFileData, err := ioutil.ReadFile(".pac")
  errors.QuitIfError(err)
  json.Unmarshal(pacFileData, &pacFile)
  return pacFile
}

func downloadJenkinsCliJar(jenkinsUrl string) {
  commands.Run("wget " + jenkinsUrl + "/jnlpJars/jenkins-cli.jar", "")
}

func createPipelineProvisionerXml(projectName string, jenkinsUrl string) {
  filePath := "pipeline-provisioner-job.xml"
  config := make(map[string]string)
  config["projectName"] = projectName
  const template = `<?xml version='1.1' encoding='UTF-8'?>
<project>
  <actions/>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties/>
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
  for DIR in *; do
    if [ -d &quot;${DIR}&quot; ]; then
      if java -jar ~/jenkins-cli.jar -s http://localhost:8080 -auth pyramid:systems get-job &quot;${DIR}&quot;; then
        echo &quot;${DIR} Jenkins Pipeline Exists. Skipping&quot;
      else
cat &lt;&lt;- EOF &gt; job.xml
&lt;?xml version=&apos;1.1&apos; encoding=&apos;UTF-8&apos;?&gt;
&lt;flow-definition plugin=&quot;workflow-job@2.31&quot;&gt;
  &lt;description&gt;&lt;/description&gt;
  &lt;keepDependencies&gt;false&lt;/keepDependencies&gt;
  &lt;properties&gt;
    &lt;com.coravy.hudson.plugins.github.GithubProjectProperty plugin=&quot;github@1.29.3&quot;&gt;
      &lt;projectUrl&gt;https://github.com/PyramidSystemsInc/foxtrot/&lt;/projectUrl&gt;
      &lt;displayName&gt;&lt;/displayName&gt;
    &lt;/com.coravy.hudson.plugins.github.GithubProjectProperty&gt;
    &lt;org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty&gt;
      &lt;triggers&gt;
        &lt;com.cloudbees.jenkins.GitHubPushTrigger plugin=&quot;github@1.29.3&quot;&gt;
          &lt;spec&gt;&lt;/spec&gt;
        &lt;/com.cloudbees.jenkins.GitHubPushTrigger&gt;
      &lt;/triggers&gt;
    &lt;/org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty&gt;
  &lt;/properties&gt;
  &lt;definition class=&quot;org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition&quot; plugin=&quot;workflow-cps@2.61&quot;&gt;
    &lt;scm class=&quot;hudson.plugins.git.GitSCM&quot; plugin=&quot;git@3.9.1&quot;&gt;
      &lt;configVersion&gt;2&lt;/configVersion&gt;
      &lt;userRemoteConfigs&gt;
        &lt;hudson.plugins.git.UserRemoteConfig&gt;
          &lt;url&gt;https://github.com/PyramidSystemsInc/foxtrot.git&lt;/url&gt;
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
fi</command>
    </hudson.tasks.Shell>
  </builders>
  <publishers/>
  <buildWrappers/>
</project>
`
  files.CreateFromTemplate(filePath, template, config)
}

func createPipelineProvisionerJob(jenkinsUrl string) {
  jenkinsCliCommandStart := "java -jar jenkins-cli.jar -s http://" + jenkinsUrl + " -auth pyramid:systems"
  jobData := files.Read("pipeline-provisioner-job.xml")
  createJobCommand := jenkinsCliCommandStart + " create-job pipeline-provisioner"
  commands.RunWithStdin(createJobCommand, string(jobData), "")
}

func cleanUp() {
  commands.Run("rm jenkins-cli.jar pipeline-provisioner-job.xml", "")
}

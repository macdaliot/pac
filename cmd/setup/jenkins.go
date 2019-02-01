package setup

import (
  "encoding/json"
  "io/ioutil"
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

type PacFile struct {
  ProjectName      string  `json:"projectName"`
  GitAuth          string  `json:"gitAuth"`
  JenkinsUrl       string  `json:"jenkinsUrl"`
  LoadBalancerArn  string  `json:"loadBalancerArn"`
  ListenerArn      string  `json:"listenerArn"`
  ServiceUrl       string  `json:"serviceUrl"`
}

func Jenkins(projectName string, projectFqdn string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-jenkins")
  imageName := "pac-jenkins"
  securityGroupName := "pac-jenkins"
  awsSession := aws.CreateAwsSession(region)
  ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
    {
      Name: imageName,
      ImageName: imageName,
      Essential: true,
    },
  })
  publicIp := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  saveJenkinsIpToPacFile(projectName, publicIp)
  jenkinsUrl := publicIp
  if projectFqdn != str.Concat(projectName, ".") {
    jenkinsFqdn := str.Concat("jenkins.", projectFqdn)
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "A", jenkinsFqdn, []string{publicIp}, ttl, awsSession)
    jenkinsUrl = jenkinsFqdn
  }
  ecs.TagCluster(clusterName, "pac-project-name", projectName, awsSession)
  logger.Info(str.Concat("Jenkins will start up in a minute or so running at ", jenkinsUrl, ":8080"))
}

func saveJenkinsIpToPacFile(projectName string, publicIp string) {
  pacFile := readPacFile(projectName)
  pacFile.JenkinsUrl = str.Concat(publicIp, ":8080")
  writePacFile(pacFile)
}

func readPacFile(projectName string) PacFile {
  // TODO: Should run from anywhere
  // TODO: Should not depend on pacFile for git
  var pacFile PacFile
  pacFileData, err := ioutil.ReadFile(str.Concat(projectName, "/.pac"))
  errors.QuitIfError(err)
  json.Unmarshal(pacFileData, &pacFile)
  return pacFile
}

func writePacFile(pacFile PacFile) {
  pacFileData, err := json.Marshal(pacFile)
  errors.QuitIfError(err)
  err = ioutil.WriteFile(str.Concat(pacFile.ProjectName, "/.pac"), []byte(pacFileData), 0644)
  errors.QuitIfError(err)
}

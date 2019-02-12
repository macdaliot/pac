package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

func Jenkins(projectName string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-jenkins")
  imageName := "pac-jenkins"
  securityGroupName := "pac-jenkins"
  awsSession := aws.CreateAwsSession(region)
  taskDefinitionArn := ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
    {
      Name: imageName,
      ImageName: imageName,
      Essential: true,
    },
  })
  tagKey := "pac-project-name"
  ecs.TagTaskDefinition(taskDefinitionArn, tagKey, projectName, awsSession)
  jenkinsUrl := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  config.Set("jenkinsUrl", str.Concat(jenkinsUrl, ":8080"))
  projectFqdn := config.Get("projectFqdn")
  if projectFqdn != str.Concat(projectName, ".") {
    jenkinsFqdn := str.Concat("jenkins.", projectFqdn)
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "A", jenkinsFqdn, []string{jenkinsUrl}, ttl, awsSession)
    jenkinsUrl = jenkinsFqdn
  }
  ecs.TagCluster(clusterName, tagKey, projectName, awsSession)
  logger.Info(str.Concat("Jenkins will start up in a minute or so running at ", jenkinsUrl, ":8080"))
}

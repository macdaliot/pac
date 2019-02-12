package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

func Selenium(projectName string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-selenium")
  securityGroupName := "pac-selenium"
  awsSession := aws.CreateAwsSession(region)
  taskDefinitionArn := ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
    {
      Name: "selenium",
      ImageName: "selenium",
      Essential: true,
    },
  })
  tagKey := "pac-project-name"
  ecs.TagTaskDefinition(taskDefinitionArn, tagKey, projectName, awsSession)
  seleniumUrl := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  projectFqdn := config.Get("projectFqdn")
  if projectFqdn != str.Concat(projectName, ".") {
    seleniumFqdn := str.Concat("selenium.", projectFqdn)
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "A", seleniumFqdn, []string{seleniumUrl}, ttl, awsSession)
    seleniumUrl = seleniumFqdn
  }
  ecs.TagCluster(clusterName, tagKey, projectName, awsSession)
  logger.Info(str.Concat("Selenium will start up in a minute or so running at ", seleniumUrl, ":4444"))
}

package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func Selenium(projectName string, projectFqdn string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-selenium")
  securityGroupName := "pac-selenium"
  awsSession := aws.CreateAwsSession(region)
  ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
    {
      Name: "selenium",
      ImageName: "selenium",
      Essential: true,
    },
  })
  publicIp := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  seleniumUrl := publicIp
  if projectFqdn != str.Concat(projectName, ".") {
    seleniumFqdn := str.Concat("selenium.", projectFqdn)
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "A", seleniumFqdn, []string{publicIp}, ttl, awsSession)
    seleniumUrl = seleniumFqdn
  }
  ecs.TagCluster(clusterName, "pac-project-name", projectName, awsSession)
  logger.Info(str.Concat("Selenium will start up in a minute or so running at ", seleniumUrl, ":4444"))
}

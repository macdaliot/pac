package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func Selenium(projectName string) {
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
  logger.Info(str.Concat("Selenium will start up in a minute or so running at ", publicIp, ":4444"))
}

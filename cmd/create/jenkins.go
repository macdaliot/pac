package create

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecr"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/logger"
)

func Jenkins(projectName string) {
  region := "us-east-2"
  ecr.Login(region)
  awsSession := aws.CreateAwsSession(region)
  clusterName := "pac-" + projectName
  familyName := clusterName + "-jenkins"
  ecs.RegisterFargateTaskDefinition(familyName, awsSession, "pac-jenkins")
  publicIp := ecs.LaunchFargateContainer(familyName, clusterName, awsSession)
  logger.Info("Jenkins will start up in a minute or so running at " + publicIp + ":8080")
}

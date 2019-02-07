package cleanCloud

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/resourcegroups"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func DeleteAllResources(projectName string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
  resourcegroups.DeleteAllResources(projectName, awsSession)
  logger.Info(str.Concat("Deleted all resources in the PAC project ", projectName))
  logger.Warn("Route53 and Cloudfront resources are not yet supported by this command and were not deleted")
}

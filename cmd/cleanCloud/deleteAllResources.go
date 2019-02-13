package cleanCloud

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/cloudfront"
  "github.com/PyramidSystemsInc/go/aws/resourcegroups"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func DeleteAllResources(projectName string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  resourcegroups.Create(projectName, "pac-project-name", projectName, awsSession)
  resourcegroups.DeleteAllResources(projectName, awsSession)
  logger.Info(str.Concat("Finished deleting all AWS resources tagged as part of PAC project ", projectName))
  // AWS Resource Groups does not have the ability to find certain types of
  // resources even if they are properly tagged. This is why Route53 and
  // Cloudfront are handled separately below
  fqdn := str.Concat(projectName, ".pac.pyramidchallenges.com.")
  route53.DeleteHostedZone(fqdn, awsSession)
  route53.DeleteRecord("pac.pyramidchallenges.com", fqdn, awsSession)
  logger.Info("Finished deleting Route53 resources")
  cloudfront.DisableDistribution(str.Concat("integration.", projectName, ".pac.pyramidchallenges.com"), awsSession)
  cloudfront.DisableDistribution(str.Concat("demo.", projectName, ".pac.pyramidchallenges.com"), awsSession)
  logger.Info("Finished disabling the Cloudfront distributions")
  logger.Warn("    Cloudfront distributions have been disabled, not deleted")
  logger.Warn("    There is a Cloudfront origin access identity that PAC does not yet support deleting")
}

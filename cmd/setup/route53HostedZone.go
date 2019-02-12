package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

func Route53HostedZone(projectName string, hostedZone string) {
  region := "us-east-2"
  var ttl int64 = 172800
  awsSession := aws.CreateAwsSession(region)
  projectFqdn := str.Concat(projectName, ".", hostedZone)
  nameServers := route53.CreateHostedZone(projectFqdn, awsSession)
  route53.TagHostedZone(projectFqdn, "pac-project-name", projectName, awsSession)
  route53.ChangeRecord(hostedZone, "NS", projectFqdn, nameServers, ttl, awsSession)
  logger.Info(str.Concat("Created name server records for ", projectFqdn))
  config.Set("projectFqdn", projectFqdn)
}

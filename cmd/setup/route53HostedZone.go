package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func Route53HostedZone(projectName string, hostedZone string) {
  if hostedZone != "" {
    region := "us-east-2"
    var ttl int64 = 172800
    awsSession := aws.CreateAwsSession(region)
    domainName := str.Concat(projectName, ".", hostedZone)
    nameServers := route53.CreateHostedZone(domainName, awsSession)
    route53.ChangeRecord(hostedZone, "NS", domainName, nameServers, ttl, awsSession)
    logger.Info(str.Concat("Created name server records for ", domainName))
  }
}

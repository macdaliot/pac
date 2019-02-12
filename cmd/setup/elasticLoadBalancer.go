package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/elbv2"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/config"
)

func ElasticLoadBalancer(projectName string) {
  region := "us-east-2"
  name := str.Concat("pac-", projectName, "-i")
  awsSession := aws.CreateAwsSession(region)
  loadBalancerArn, listenerArn, serviceUrl := elbv2.Create(name, awsSession)
  elbv2.Tag(name, "pac-project-name", projectName, awsSession)
  config.Set("loadBalancerArn", loadBalancerArn)
  config.Set("listenerArn", listenerArn)
  config.Set("serviceUrl", serviceUrl)
  projectFqdn := config.Get("projectFqdn")
  if projectFqdn != str.Concat(projectName, ".") {
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "CNAME", str.Concat("api.", projectFqdn), []string{serviceUrl}, ttl, awsSession)
  }
  logger.Info("Created application load balancer and HTTP listener to support future microservices")
}

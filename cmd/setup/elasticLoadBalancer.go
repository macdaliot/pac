package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/elbv2"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func ElasticLoadBalancer(projectName string, projectFqdn string) {
  region := "us-east-2"
  name := str.Concat("pac-", projectName, "-i")
  awsSession := aws.CreateAwsSession(region)
  loadBalancerArn, listenerArn, serviceUrl := elbv2.CreateLoadBalancer(name, awsSession)
  pacFile := readPacFile(projectName)
  pacFile.LoadBalancerArn = loadBalancerArn
  pacFile.ListenerArn = listenerArn
  pacFile.ServiceUrl = serviceUrl
  writePacFile(pacFile)
  if projectFqdn != str.Concat(projectName, ".") {
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "CNAME", str.Concat("api.", projectFqdn), []string{serviceUrl}, ttl, awsSession)
  }
  logger.Info("Created application load balancer and HTTP listener to support future microservices")
}

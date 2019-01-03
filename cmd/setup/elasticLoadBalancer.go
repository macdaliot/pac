package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/elbv2"
  "github.com/PyramidSystemsInc/go/logger"
)

func ElasticLoadBalancer(projectName string) {
  region := "us-east-2"
  awsSession := aws.CreateAwsSession(region)
  name := "pac-" + projectName
  loadBalancerArn, listenerArn, serviceUrl := elbv2.CreateLoadBalancer(name, awsSession)
  pacFile := readPacFile(projectName)
  pacFile.LoadBalancerArn = loadBalancerArn
  pacFile.ListenerArn = listenerArn
  pacFile.ServiceUrl = serviceUrl
  writePacFile(pacFile)
  logger.Info("Created application load balancer and HTTP listener to support future microservices")
}

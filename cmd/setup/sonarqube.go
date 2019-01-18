package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func SonarQube(projectName string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-sonarqube")
  securityGroupName := "pac-sonarqube"
  awsSession := aws.CreateAwsSession(region)
  ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
    {
      Name: "sonar-db",
      ImageName: "pac-sonar-db",
      Essential: true,
      EnvironmentVars: map[string]string{
        "POSTGRES_PASSWORD": "pyramid",
      },
    },
    {
      Name: "sonarqube",
      ImageName: "sonarqube",
      Essential: false,
      EnvironmentVars: map[string]string{
        "sonar.jdbc.username": "sonar",
        "sonar.jdbc.password": "sonar",
        "sonar.jdbc.url": "jdbc:postgresql://localhost/sonar",
      },
    },
  })
  publicIp := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  // saveJenkinsIpToPacFile(projectName, publicIp)
  logger.Info(str.Concat("SonarQube will start up in a minute or so running at ", publicIp, ":9000"))
}

/*
func saveJenkinsIpToPacFile(projectName string, publicIp string) {
  pacFile := readPacFile(projectName)
  pacFile.JenkinsUrl = str.Concat(publicIp, ":8080")
  writePacFile(pacFile)
}
*/

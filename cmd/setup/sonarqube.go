package setup

import (
  "github.com/PyramidSystemsInc/go/aws"
  "github.com/PyramidSystemsInc/go/aws/ecs"
  "github.com/PyramidSystemsInc/go/aws/route53"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func SonarQube(projectName string, projectFqdn string) {
  region := "us-east-2"
  clusterName := str.Concat("pac-", projectName)
  familyName := str.Concat(clusterName, "-sonarqube")
  securityGroupName := "pac-sonarqube"
  awsSession := aws.CreateAwsSession(region)
  taskDefinitionArn := ecs.RegisterFargateTaskDefinition(familyName, awsSession, []ecs.Container{
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
  tagKey := "pac-project-name"
  ecs.TagTaskDefinition(taskDefinitionArn, tagKey, projectName, awsSession)
  publicIp := ecs.LaunchFargateContainer(familyName, clusterName, securityGroupName, awsSession)
  sonarqubeUrl := publicIp
  if projectFqdn != str.Concat(projectName, ".") {
    sonarqubeFqdn := str.Concat("sonarqube.", projectFqdn)
    var ttl int64 = 300
    route53.ChangeRecord(projectFqdn, "A", sonarqubeFqdn, []string{publicIp}, ttl, awsSession)
    sonarqubeUrl = sonarqubeFqdn
  }
  ecs.TagCluster(clusterName, tagKey, projectName, awsSession)
  logger.Info(str.Concat("SonarQube will start up in a minute or so running at ", sonarqubeUrl, ":9000"))
}

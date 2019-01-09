package add

import (
  "encoding/json"
  "io/ioutil"
  "strings"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/pac/cmd/add/service"
)

type PacFile struct {
  ProjectName      string  `json:"projectName"`
  GitAuth          string  `json:"gitAuth"`
  JenkinsUrl       string  `json:"jenkinsUrl"`
  LoadBalancerArn  string  `json:"loadBalancerArn"`
  ListenerArn      string  `json:"listenerArn"`
  ServiceUrl       string  `json:"serviceUrl"`
}

// Service adds a new service to the application
func Service(serviceName string) {
  ensureRunningFromServicesDirectory()
  createServiceDirectory(serviceName)
  config := createTemplateConfig(serviceName)
  createServiceFiles(serviceName, config)
  createDynamoDbTable(serviceName)
  launchMicroservice(serviceName)
}

func ensureRunningFromServicesDirectory() {
  workingDirectory := directories.GetWorking()
  if !strings.HasSuffix(workingDirectory, "svc") {
    errors.LogAndQuit("`pac add service [flags]` must be run from the svc directory")
  }
}

func createTemplateConfig(serviceName string) map[string]string {
  config := make(map[string]string)
  config["projectName"] = readPacFile().ProjectName
  config["serviceName"] = serviceName
  return config
}

func createServiceDirectory(serviceName string) {
  workingDirectory := directories.GetWorking()
  directories.Create(str.Concat(workingDirectory, "/", serviceName))
}

func createServiceFiles(serviceName string, config map[string]string) {
  service.CreatePackageJSON(str.Concat(serviceName, "/package.json"), config)
  service.CreateDockerfile(str.Concat(serviceName, "/Dockerfile"))
  service.CreateServerTs(str.Concat(serviceName, "/server.ts"), config)
  service.CreateDynamoConfigJSON(str.Concat(serviceName, "/dynamoConfig.json"), config)
  service.CreateAwsSdkConfigJs(str.Concat(serviceName, "/awsSdkConfig.js"), config)
  service.CreateLaunchSh(str.Concat(serviceName, "/launch.sh"), config)
  service.CreateJenkinsfile(str.Concat(serviceName, "/Jenkinsfile"))
  service.CreateBuildSh(str.Concat(serviceName, "/.build.sh"), serviceName)
  service.CreateDeploySh(str.Concat(serviceName, "/.deploy.sh"), config)
  logger.Info(str.Concat("Created ", serviceName, " Express microservice files"))
}

func createDynamoDbTable(serviceName string) {
  workingDirectory := directories.GetWorking()
  commands.Run(str.Concat("aws dynamodb create-table --cli-input-json file://", workingDirectory, "/", serviceName, "/dynamoConfig.json --endpoint-url http://localhost:8000"), "")
  logger.Info(str.Concat("Created ", serviceName, " DynamoDB table locally"))
}

func launchMicroservice(serviceName string) {
  commands.Run("./launch.sh", str.Concat("./", serviceName))
  logger.Info(str.Concat("Launched ", serviceName, " microservice Docker container locally (available at localhost:3000/api/", serviceName, ")"))
}

func readPacFile() PacFile {
  // TODO: Should run from anywhere
  // TODO: Should not depend on pacFile for git
  var pacFile PacFile
  pacFileData, err := ioutil.ReadFile("../.pac")
  errors.QuitIfError(err)
  json.Unmarshal(pacFileData, &pacFile)
  return pacFile
}

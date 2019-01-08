package add

import (
  "encoding/json"
  "io/ioutil"
  "strings"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
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
  directories.Create(workingDirectory + "/" + serviceName)
}

func createServiceFiles(serviceName string, config map[string]string) {
  service.CreatePackageJSON(serviceName + "/package.json", config)
  service.CreateDockerfile(serviceName + "/Dockerfile")
  service.CreateServerTs(serviceName + "/server.ts", config)
  service.CreateDynamoConfigJSON(serviceName + "/dynamoConfig.json", config)
  service.CreateAwsSdkConfigJs(serviceName + "/awsSdkConfig.js", config)
  service.CreateLaunchSh(serviceName + "/launch.sh", config)
  service.CreateJenkinsfile(serviceName + "/Jenkinsfile")
  service.CreateBuildSh(serviceName + "/.build.sh", serviceName)
  service.CreateDeploySh(serviceName + "/.deploy.sh", config)
  logger.Info("Created " + serviceName + " Express microservice files")
}

func createDynamoDbTable(serviceName string) {
  workingDirectory := directories.GetWorking()
  commands.Run("aws dynamodb create-table --cli-input-json file://" + workingDirectory + "/" + serviceName + "/dynamoConfig.json --endpoint-url http://localhost:8000", "")
  logger.Info("Created " + serviceName + " DynamoDB table locally")
}

func launchMicroservice(serviceName string) {
  commands.Run("./launch.sh", "./" + serviceName)
  logger.Info("Launched " + serviceName + " microservice Docker container locally (available at localhost:3000/api/" + serviceName + ")")
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

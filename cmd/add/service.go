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
  ProjectName     string `json:"projectName"`
  GitAuth         string `json:"gitAuth"`
  JenkinsUrl      string `json:"jenkinsUrl"`
  LoadBalancerArn string `json:"loadBalancerArn"`
  ListenerArn     string `json:"listenerArn"`
  ServiceUrl      string `json:"serviceUrl"`
}

var _serviceName string

const _testPath string = "/tests"

// Service adds a new service to the application
func Service(serviceName string) {
  _serviceName = serviceName
  ensureRunningFromServicesDirectory()
  createServiceDirectory(serviceName)
  config := createTemplateConfig(serviceName)
  createServiceFiles(serviceName, config)
  createTestsDirectory(serviceName)
  createTestFiles(config)
  // createDynamoDbTable(serviceName)
  // launchMicroservice(serviceName)
  commands.Run("npm i", str.Concat("./", serviceName))
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
  var serviceDirectory string = str.Concat(workingDirectory, "/", serviceName)
  directories.Create(serviceDirectory)
  directories.Create(str.Concat(serviceDirectory, "/src"))
}

func createTestsDirectory(serviceName string) {
  workingDirectory := directories.GetWorking()
  directories.Create(str.Concat(workingDirectory, "/", getPathForTestsFolder()))
}

func createTestFiles(config map[string]string) {
  /* add more tests as needed */
  service.CreateTestFile(getPathForTest("server.spec.ts"), config)
}

func getPathForTestsFolder() string {
  return str.Concat(_serviceName, _testPath)
}
func getPathForTest(filePath string) string {
  return str.Concat(getPathForTestsFolder(), "/", filePath)
}

func createServiceFiles(serviceName string, config map[string]string) {
  service.CreatePackageJSON(str.Concat(serviceName, "/package.json"), config)
  service.CreateDockerfile(str.Concat(serviceName, "/Dockerfile"))
  service.CreateTsConfig(str.Concat(serviceName, "/tsconfig.json"), config)
  service.CreateServerTs(str.Concat(serviceName, "/src", "/server.ts"), config)
  service.CreateDynamoConfigJSON(str.Concat(serviceName, "/dynamoConfig.json"), config)
  service.CreateAwsSdkConfigJs(str.Concat(serviceName, "/src", "/awsSdkConfig.js"), config)
  service.CreateLaunchBat(str.Concat(serviceName, "/launch.bat"), config)
  service.CreateLaunchSh(str.Concat(serviceName, "/launch.sh"), config)
  service.CreateJenkinsfile(str.Concat(serviceName, "/Jenkinsfile"))
  service.CreateBuildSh(str.Concat(serviceName, "/.build.sh"), serviceName)
  service.CreateTestSh(str.Concat(serviceName, "/.test.sh"), serviceName)
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

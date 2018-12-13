package add

import (
  "strings"
  "github.com/PyramidSystemsInc/pac/cmd/add/service"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
)

func Service(serviceName string) {
  ensureRunningFromServicesDirectory()
  createServiceDirectory(serviceName)
  config := createConfig(serviceName)
  service.CreatePackageJson(serviceName + "/package.json", config)
  service.CreateDockerfile(serviceName + "/Dockerfile")
  service.CreateServerTs(serviceName + "/server.ts", config)
  service.CreateDynamoConfigJson(serviceName + "/dynamoConfig.json", config)
  logger.Info("Created " + serviceName + " Express microservice")
  createDynamoDbTable(serviceName)
}

func ensureRunningFromServicesDirectory() {
  workingDirectory := directories.GetWorking()
  if !strings.HasSuffix(workingDirectory, "svc") {
    errors.LogAndQuit("`pac add service [flags]` must be run from the svc directory")
  }
}

func createConfig(serviceName string) map[string]string {
  config := make(map[string]string)
  config["serviceName"] = serviceName
  return config
}

func createServiceDirectory(serviceName string) {
  workingDirectory := directories.GetWorking()
  directories.Create(workingDirectory + "/" + serviceName)
}

func createDynamoDbTable(serviceName string) {
  workingDirectory := directories.GetWorking()
  commands.Run("aws dynamodb create-table --cli-input-json file://" + workingDirectory + "/" + serviceName + "/dynamoConfig.json --endpoint-url http://localhost:8000", "")
  logger.Info("Created " + serviceName + " DynamoDB table")
}

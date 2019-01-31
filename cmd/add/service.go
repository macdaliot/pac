package add

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
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
	// createTestsDirectory(serviceName)
	// createTestFiles(config)
	commands.Run("npm i", str.Concat("./", serviceName))
  editHaProxyConfig(serviceName, config["projectName"])
}

func editHaProxyConfig(serviceName string, projectName string) {
  serviceConfig := str.Concat(`backend backend_`, serviceName, `
    mode http
    server `, serviceName, ` pac-`, projectName, `-`, serviceName, `
    timeout connect 5000
    timeout server 50000



`)
  proxyAcl := str.Concat(`

    acl path_`, serviceName, ` path_beg /api/`, serviceName, `
    use_backend backend_`, serviceName, ` if path_`, serviceName, `
`)
  files.Prepend("haproxy.cfg", []byte(serviceConfig))
  files.Append("haproxy.cfg", []byte(proxyAcl))
  logger.Info("Updated the local microservice proxy configuration")
}

func ensureRunningFromServicesDirectory() {
	workingDirectory := directories.GetWorking()
	if !strings.HasSuffix(workingDirectory, "svc") {
		errors.LogAndQuit("`pac add service [flags]` must be run from the svc directory")
	}
}

func createTemplateConfig(serviceName string) map[string]string {
	config := make(map[string]string)
	pacFile := readPacFile()
	config["projectName"] = pacFile.ProjectName
	config["serviceUrl"] = pacFile.ServiceUrl
	config["serviceName"] = serviceName
	return config
}

func createServiceDirectory(serviceName string) {
	workingDirectory := directories.GetWorking()
	var serviceDirectory string = str.Concat(workingDirectory, "/", serviceName)
	directories.Create(serviceDirectory)
}

func createTestsDirectory(serviceName string) {
	workingDirectory := directories.GetWorking()
	directories.Create(str.Concat(workingDirectory, "/", getPathForTestsFolder()))
}

func createTestFiles(config map[string]string) {
	/* add more tests as needed */
	service.CreateDefaultControllerSpecTs(getPathForTest("defaultController.spec.ts"), config)
	service.CreateDefaultServiceSpecTs(getPathForTest("defaultService.spec.ts"), config)
	service.CreateMockDynamoDbTs(getPathForTest("mockDynamoDb.ts"))
	service.CreateMockDefaultServiceTs(getPathForTest("mockDefaultService.ts"))
}

func getPathForTestsFolder() string {
	return str.Concat(_serviceName, _testPath)
}
func getPathForTest(filePath string) string {
	return str.Concat(getPathForTestsFolder(), "/", filePath)
}

func createServiceFiles(serviceName string, config map[string]string) {
	service.CreateDockerfile(str.Concat(serviceName, "/Dockerfile"))
	service.CreateDynamoConfigJSON(str.Concat(serviceName, "/dynamoConfig.json"), config)
	service.CreateLaunchSh(str.Concat(serviceName, "/launch.sh"), config)
	service.CreateJenkinsfile(str.Concat(serviceName, "/Jenkinsfile"))
	service.CreateBuildSh(str.Concat(serviceName, "/.build.sh"), serviceName)
	service.CreateTestSh(str.Concat(serviceName, "/.test.sh"), serviceName)
	service.CreateDeploySh(str.Concat(serviceName, "/.deploy.sh"), config)
	service.CreateAllTemplatedFiles(serviceName, config)
	// createServiceSource(serviceName, config)
	service.CreateFrontEndClient(str.Concat("../app/src/services/", strings.Title(serviceName), ".js"), config)
	logger.Info(str.Concat("Created ", serviceName, " Express microservice files"))
}

func createServiceSource(serviceName string, config map[string]string) {
	/* need a better way to construct folders */
	workingDirectory := directories.GetWorking()
	var serviceDirectory = str.Concat(workingDirectory, "/", serviceName)
	var serviceSourceDirectory = str.Concat(serviceDirectory, "/src")
	directories.Create(str.Concat(serviceSourceDirectory))

	directories.Create(str.Concat(serviceSourceDirectory, "/config"))

	directories.Create(str.Concat(serviceSourceDirectory, "/controllers"))
	service.CreateControllerInterfaceTs(str.Concat(serviceSourceDirectory, "/controllers", "/controller.interface.ts"))
	service.CreateDefaultControllerTs(str.Concat(serviceSourceDirectory, "/controllers", "/defaultController.ts"))

	directories.Create(str.Concat(serviceSourceDirectory, "/database"))
	service.CreateDbInterfaceTs(str.Concat(serviceSourceDirectory, "/database", "/db.interface.ts"))

	directories.Create(str.Concat(serviceSourceDirectory, "/middleware"))
	directories.Create(str.Concat(serviceSourceDirectory, "/middleware", "/logger"))
	service.CreateLoggerTs(str.Concat(serviceSourceDirectory, "/middleware", "/logger", "/logger.ts"))
	service.CreateLoggerMiddlewareTs(str.Concat(serviceSourceDirectory, "/middleware", "/logger", "/loggerMiddleware.ts"))

	directories.Create(str.Concat(serviceSourceDirectory, "/routes"))
	service.CreateRoutesTS(str.Concat(serviceSourceDirectory, "/routes", "/routes.ts"), config)

	directories.Create(str.Concat(serviceSourceDirectory, "/services"))
	service.CreateDefaultServiceTs(str.Concat(serviceSourceDirectory, "/services", "/defaultService.ts"))
	service.CreateServiceInterfaceTs(str.Concat(serviceSourceDirectory, "/services", "/service.interface.ts"))

	directories.Create(str.Concat(serviceSourceDirectory, "/utility"))
	service.CreateFunctions_indexTs(str.Concat(serviceSourceDirectory, "/utility", "/index.ts"))
	service.CreateFunctionsTs(str.Concat(serviceSourceDirectory, "/utility", "/functions.ts"))
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

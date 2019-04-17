package add

import (
	"path"
	"strings"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/add/service"
	"github.com/PyramidSystemsInc/pac/config"
)

// Service adds a new service to the application
func Service(serviceName string) {
	createServiceDirectory(serviceName)
	cfg := createTemplateConfig(serviceName)
	createServiceFiles(serviceName, cfg)

	commands.Run("npm i", path.Join("services/", serviceName))
	editHaProxyConfig(serviceName, cfg["projectName"])
	editIntegrationTestApiFeatures(serviceName)
	commands.Run("terraform init -input=false", path.Join("services/", "terraform"))
}

func createServiceDirectory(serviceName string) {
	serviceDirectory := str.Concat("services/", serviceName)
	directories.Create(serviceDirectory)
}

func createTemplateConfig(serviceName string) map[string]string {
	cfg := make(map[string]string)
	cfg["projectName"] = config.Get("projectName")
	cfg["serviceName"] = serviceName
	cfg["serviceNamePascal"] = strings.Title(serviceName)
	return cfg
}



func createServiceFiles(serviceName string, cfg map[string]string) {
	logger.Info("Create service files")
	service.CreateAllTemplatedFiles(cfg)
	logger.Info(str.Concat("Created ", serviceName, " Express microservice files"))
}

func editHaProxyConfig(serviceName string, projectName string) {
	haProxyConfigPath := "services/haproxy.cfg"
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
	files.Prepend(haProxyConfigPath, []byte(serviceConfig))
	files.Append(haProxyConfigPath, []byte(proxyAcl))
	logger.Info("Updated the local microservice proxy configuration")
}

func editIntegrationTestApiFeatures(serviceName string) {
  filePath := "integration-tests/src/test/resources/features/API.feature"
  lineToMatch := "      | endpoint  | status | token |"
  lineLength := 10
  serviceNameWithTrailingSpaces := serviceName
  for i := 0; i < lineLength - len(serviceName); i++ {
    serviceNameWithTrailingSpaces += " "
  }
  newLine := str.Concat("      | ", serviceNameWithTrailingSpaces, "| 200    | false |")
  files.AppendBelow(filePath, lineToMatch, newLine)

  lineToMatch = "      | endpoint  |"
  files.AppendBelow(filePath, lineToMatch, `      | ` + serviceNameWithTrailingSpaces + `|`)
  logger.Info(str.Concat("Edited the integration tests to test the new ", serviceName, " endpoints"))
}

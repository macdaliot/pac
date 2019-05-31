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
	editPackageJsonScripts(serviceName)
  editCfDeployScript(serviceName)
	commands.Run("terraform init -input=false", path.Join("services/", "terraform"))
}

func createServiceDirectory(serviceName string) {
	serviceDirectory := str.Concat("services/", serviceName)
	directories.Create(serviceDirectory)
}

func createTemplateConfig(serviceName string) map[string]string {
	cfg := make(map[string]string)
	cfg["projectName"] = config.Get("projectName")
	cfg["region"] = config.Get("region")
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

func editPackageJsonScripts(serviceName string) {
	filePath := "package.json"
	lineToMatch := "  \"scripts\": {"
	newLine := str.Concat("    \"start:", serviceName, "\": \"pushd app; npm i; popd; pushd core; npm i; popd; pushd domain; npm i; popd; cd services/" + serviceName + "; npm i\",");
	files.AppendBelow(filePath, lineToMatch, newLine)
	logger.Info(str.Concat("Edited the root package.json to include the new microservice start script"))
}

func editCfDeployScript(serviceName string) {
	filePath := "services/.cfDeploy.sh"
	lineToMatch := "#! /bin/bash"
	newLine := str.Concat("cf push -f manifest.yml -c \"npm run start:", serviceName, "\"")
	files.AppendBelow(filePath, lineToMatch, newLine)
	logger.Info(str.Concat("Edited the .cfDeploy.sh script to include the new microservice in its deploy to Cloud Foundry"))
}

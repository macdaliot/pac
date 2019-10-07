package add

import (
	"path"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/pac/cmd/add/authservice"
)

// Service adds a new service to the application
func AuthService() {
	serviceName := "auth"
	createServiceDirectory(serviceName)
	config := createTemplateConfig(serviceName)
	authservice.CreateAllTemplatedFiles(config)
	commands.Run("npm i", path.Join("services/", serviceName))
	editHaProxyConfig(serviceName, config["projectName"])
}

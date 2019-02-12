package add

import (
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/add/authservice"
)

// Service adds a new service to the application
func AuthService() {
	serviceName := "auth"
	createServiceDirectory(serviceName)
	config := createTemplateConfig(serviceName)
	authservice.CreateAllTemplatedFiles(serviceName, config)
	commands.Run("npm i", str.Concat("./", serviceName))
	editHaProxyConfig(serviceName, config["projectName"])
}

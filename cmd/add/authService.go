package add

import (
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/add/authservice"
)

// Service adds a new service to the application
func AuthService() {
	_serviceName = "auth"
	ensureRunningFromServicesDirectory()
	createServiceDirectory(_serviceName)
	config := createTemplateConfig(_serviceName)
	authservice.CreateAllTemplatedFiles(_serviceName, config)
	commands.Run("npm i", str.Concat("./", _serviceName))
	editHaProxyConfig(_serviceName, config["projectName"])
}

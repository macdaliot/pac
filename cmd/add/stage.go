package add

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/cmd/add/stages"
)

// Stage adds a new stage to a Jenkinsfile
func Stage(stageName string) {
	if stageName == "npm_install" {
		stages.AppendNPMInstall(stageName)
	} else if stageName == "test" {
		stages.AppendTest(stageName)
	} else if stageName == "deploy" {
		stages.AppendDeploy(stageName)
	} else {
		logger.Err("That stage name is unknown to PAC")
	}
}

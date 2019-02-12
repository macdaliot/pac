package setup

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/cmd/setup/frontEndFiles"
	"github.com/PyramidSystemsInc/pac/config"
)

func FrontEndFiles() {
	frontEndFiles.CreateAllTemplatedFiles("app", createConfig())
	logger.Info("Created ReactJS front-end")
}

func createConfig() map[string]string {
	cfg := make(map[string]string)
	cfg["projectName"] = config.Get("projectName")
	cfg["description"] = config.Get("description")
	return cfg
}

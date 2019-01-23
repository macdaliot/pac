package setup

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/cmd/setup/frontEndFiles"
)

func FrontEndFiles(projectDirectory string, projectName string, description string) {
	config := createConfig(projectName, description)
	frontEndFiles.CreateAllTemplatedFiles(str.Concat(projectDirectory, "/app"), config)
	logger.Info("Created ReactJS front-end")
}

func createConfig(projectName string, description string) map[string]string {
	config := make(map[string]string)
	config["projectName"] = projectName
	config["description"] = description
	return config
}

package deploy

import (
  "os"

  "github.com/PyramidSystemsInc/go/logger"
)

func SetEnvironmentVariables(projectName string) {
  os.Setenv("TF_IN_AUTOMATION", "NONEMPTYVALUE")
  os.Setenv("TF_VAR_project_name", projectName)
  logger.Info("Environment variables are set")
}

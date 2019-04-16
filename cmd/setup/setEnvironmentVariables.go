package setup

import (
  "os"

  "github.com/PyramidSystemsInc/go/logger"
)

func SetEnvironmentVariables(projectName string) {
  logger.Info("Environment variables are set")
  os.Setenv("TF_IN_AUTOMATION", "NONEMPTYVALUE")
  os.Setenv("TF_VAR_project_name", projectName)
}

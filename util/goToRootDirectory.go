package util

import (
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/pac/config"
)

// GoToRootDirectory - Changes the working directory to the root directory of the PAC project
func GoToRootDirectory() {
  directories.Change(config.GetRootDirectory())
}

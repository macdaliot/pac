package service

import (
  "runtime"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

func CreateTestSh(filePath string, serviceName string) {
  const template = `#! /bin/bash

npm run test
if [[ $(echo $?) -eq 0 ]]; then
  echo "All unit tests passed"
else
  echo "Stopping pipeline due to unit test failures"
  exit 2
fi
`
  files.CreateFromTemplate(filePath, template, nil)
  if runtime.GOOS == "windows" {
    files.ChangePermissions(str.Concat(".\\", serviceName, "\\.test.sh"), 0755)
  } else {
    files.ChangePermissions(str.Concat("./", serviceName, "/.test.sh"), 0755)
  }
}

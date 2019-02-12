package service

import (
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

func CreateTestSh(fileName string, serviceName string) {
  const template = `#! /bin/bash

npm run test
if [[ $(echo $?) -eq 0 ]]; then
  echo "All unit tests passed"
else
  echo "Stopping pipeline due to unit test failures"
  exit 2
fi
`
  files.CreateFromTemplate(fileName, template, nil)
  files.ChangePermissions(str.Concat("./", fileName), 0755)
}

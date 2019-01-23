package service

import (
  "runtime"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

func CreateBuildSh(filePath string, serviceName string) {
  const template = `#! /bin/bash

# Zip up Lambda function code
npm i
npx tsc
zip -r function dist/* lambda.js node_modules >> /dev/null
`
  files.CreateFromTemplate(filePath, template, nil)
  if runtime.GOOS == "windows" {
    files.ChangePermissions(str.Concat(".\\", serviceName, "\\.build.sh"), 0755)
  } else {
    files.ChangePermissions(str.Concat("./", serviceName, "/.build.sh"), 0755)
  }
}

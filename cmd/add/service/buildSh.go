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
echo $(sed -e 's/awsSdkConfig.local/awsSdkConfig.cloud/g' -e 's/app.listen(port);/module.exports = app;/g' dist/server.js) > dist/server.js
npx claudia generate-serverless-express-proxy --express-module dist/server
zip -r function dist/* lambda.js node_modules >> /dev/null
`
  files.CreateFromTemplate(filePath, template, nil)
  if runtime.GOOS == "windows" {
    files.ChangePermissions(str.Concat(".\\", serviceName, "\\.build.sh"), 0755)
  } else {
    files.ChangePermissions(str.Concat("./", serviceName, "/.build.sh"), 0755)
  }
}

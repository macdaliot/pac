package service

import (
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/str"
)

func CreateBuildSh(filePath string, serviceName string) {
  const template = `#! /bin/bash

# Zip up Lambda function code
npm i
npx tsc server.ts
echo $(sed -e 's/awsSdkConfig.local/awsSdkConfig.cloud/g' -e 's/app.listen(port);/module.exports = app;/g' server.js) > server.js
npx claudia generate-serverless-express-proxy --express-module server
zip -r function awsSdkConfig.js lambda.js server.js node_modules >> /dev/null
`
  files.CreateFromTemplate(filePath, template, nil)
  commands.Run(str.Concat("chmod 755 ", serviceName, "/.build.sh"), "")
}

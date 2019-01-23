package service

import (
  "github.com/PyramidSystemsInc/go/files"
)

func CreateLambdaJs(filePath string) {
  const template = `'use strict'
const awsServerlessExpress = require('aws-serverless-express');
const app = require('./dist/server');
const server = awsServerlessExpress.createServer(app);
 
exports.handler = (event, context) => awsServerlessExpress.proxy(server, event, context);
`
  files.CreateFromTemplate(filePath, template, nil)
}

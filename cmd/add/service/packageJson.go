package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreatePackageJSON creates a default package.json based on passed in config
func CreatePackageJSON(filePath string, config map[string]string) {
	const template = `{
  "name": "pac-{{.serviceName}}-service",
  "version": "0.0.1",
  "description": "{{.serviceName}} service (created by PAC)",
  "main": "server.js",
  "scripts": {
    "start": "nodemon",
    "test": "mocha --require ts-node/register tests/**/*.spec.ts"
  },
  "author": "Pyramid Systems, Inc.",
  "license": "ISC",
  "dependencies": {
    "aws-sdk": "^2.375.0",
    "aws-serverless-express": "^3.3.5",
    "body-parser": "^1.18.3",
    "claudia": "^5.3.0",
    "cors": "^2.8.5",
    "express": "^4.16.4",
    "lodash": "^4.17.11",
    "typescript": "^3.2.2",
    "uuid": "^3.3.2"
  },
  "devDependencies": {
    "@types/chai": "^4.1.7",
    "@types/express": "^4.16.0",
    "@types/lodash": "^4.14.119",
    "@types/uuid": "^3.4.4",
    "@types/node": "^10.12.12",
    "@types/mocha": "^5.2.5",
    "aws-sdk-mock": "^4.3.0",
    "chai": "^4.2.0",
    "mocha": "^5.2.0",
    "nodemon": "^1.18.9",
    "ts-node": "^7.0.1",
    "typescript": "^3.2.2"
    
  }
}
`
	files.CreateFromTemplate(filePath, template, config)
}

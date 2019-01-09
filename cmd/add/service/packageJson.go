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
    "start": "node server"
  },
  "author": "Pyramid Systems, Inc.",
  "license": "ISC",
  "dependencies": {
    "aws-sdk": "^2.375.0",
    "aws-serverless-express": "^3.3.5",
    "body-parser": "^1.18.3",
    "claudia": "^5.3.0",
    "express": "^4.16.4",
    "uuid": "^3.3.2"
  },
  "devDependencies": {
    "@types/node": "^10.12.12",
    "typescript": "^3.2.2"
  }
}
`
	files.CreateFromTemplate(filePath, template, config)
}

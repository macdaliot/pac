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
    "@types/node": "^10.12.12",
    "aws-sdk": "^2.375.0",
    "body-parser": "^1.18.3",
    "express": "^4.16.4",
    "mongodb": "^3.1.10",
    "uuid": "^3.3.2"
  },
  "devDependencies": {
    "typescript": "^3.2.2"
  }
}
`
	files.CreateFromTemplate(filePath, template, config)
}

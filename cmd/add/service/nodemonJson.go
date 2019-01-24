package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateNodemonJson creates a default package.json based on passed in config
func CreateNodemonJson(filePath string) {
	const template = `{
		"watch": [
			"src"
		],
		"ext": "ts",
		"ignore": [
			"src/**/*.spec.ts"
		],
		"exec": "ts-node ./src/local.ts"
	}
`
	files.CreateFromTemplate(filePath, template, nil)
}

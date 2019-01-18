package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateFunctions_indexTs(filePath string) {
	const template = `export * from './functions';`
	files.CreateFromTemplate(filePath, template, nil)
}

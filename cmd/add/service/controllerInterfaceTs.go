package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateControllerInterfaceTs(filePath string) {
	const template = `export interface Controller {
        serviceInstance;
    }`
	files.CreateFromTemplate(filePath, template, nil)
}

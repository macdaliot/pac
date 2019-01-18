package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateServiceInterfaceTs(filePath string) {
	const template = `export interface Service {
		db;
	}`
	files.CreateFromTemplate(filePath, template, nil)
}

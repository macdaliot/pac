package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateDbInterfaceTs(filePath string) {
	const template = `export interface Database {
		dbInstance;
		query(params: any);
		create(object: any);
		update(params: any, object: any);
	}`
	files.CreateFromTemplate(filePath, template, nil)
}

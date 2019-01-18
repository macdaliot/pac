package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateLoggerTs(filePath string) {
	const template = `export class Logger {
		info = text => console.log(text);
		error = error => console.error(error);
	}`
	files.CreateFromTemplate(filePath, template, nil)
}

package service

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(cfg map[string]string) {
	serviceName := cfg["serviceName"]

	// make sure we are at root dir
	os.Chdir(config.GetRootDirectory())

	// service templates
	options := files.TemplateOptions{
		Box:             packr.NewBox("./serviceTemplates"),
		TargetDirectory: path.Join("services/", serviceName),
		Config:          cfg}
	err := files.CreateTemplatedFiles(options)
	errors.QuitIfError(err)

	// domain templates
	options.Box = packr.NewBox("./domainTemplates")
	options.TargetDirectory = "domain"
	err = files.CreateTemplatedFiles(options)
	errors.QuitIfError(err)

	// manually update the index.ts
	// TODO: replace with another template perhaps
	logger.Info("Updating the index.ts for domain")

	files.EnsurePath(filepath.Dir("domain"))
	exportTemplate := fmt.Sprintf("\nexport * from './%s';", serviceName)
	exportModelTemplate := fmt.Sprintf("\nexport * from './%s/%s';", serviceName, serviceName)
	files.Append("./domain/index.ts", []byte(exportTemplate))
	files.Append("./domain/models.ts", []byte(exportModelTemplate))

	// terraform templates
	options.Box = packr.NewBox("./terraformTemplates")
	options.TargetDirectory = "./services/terraform"
	err = files.CreateTemplatedFiles(options)
	errors.QuitIfError(err)
}

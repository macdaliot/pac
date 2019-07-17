package service

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/gobuffalo/packr"
)

func CreateAllTemplatedFiles(cfg map[string]string) {
	serviceName := cfg["serviceName"]

	os.Chdir(config.GetRootDirectory())

	// service templates
	options := files.TemplateOptions{
		Box:             packr.NewBox("./serviceTemplates"),
		TargetDirectory: path.Join("services/", serviceName),
		Config:          cfg,
	}
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

	// Terraform templates
	environmentNames := strings.Split(config.Get("environments"), ",")
	for _, environmentName := range environmentNames {
		os.Chdir(config.GetRootDirectory())
		cfg["environmentName"] = environmentName
		options = files.TemplateOptions{
			Box:             packr.NewBox("../environment/environmentLambdaTemplates"),
			TargetDirectory: str.Concat("terraform/", environmentName),
			Config:          cfg,
		}
		err = files.CreateTemplatedFiles(options)
		errors.QuitIfError(err)
	}
}

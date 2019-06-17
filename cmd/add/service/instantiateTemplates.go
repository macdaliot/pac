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
	os.Chdir(path.Join("services/", serviceName))

	box := packr.NewBox("./serviceTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))
		template, _ := box.FindString(templatePath)
		if templatePath == "src/controller.ts" {
			files.CreateFromTemplate(str.Concat("src/", serviceName, ".controller.ts"), template, cfg)
		} else if templatePath == "src/controller.spec.ts" {
			files.CreateFromTemplate(str.Concat("src/", serviceName, ".controller.spec.ts"), template, cfg)
		} else {
			files.CreateFromTemplate(templatePath, template, cfg)
		}
	}
	os.Chdir(config.GetRootDirectory())
	os.Chdir("domain")

	logger.Info("Writing the domain templates")

	box = packr.NewBox("./domainTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)

		template, _ := box.FindString(templatePath)

		if strings.Contains(templatePath, "serviceNameFolder") {
			templatePath = strings.Replace(templatePath, "serviceNameFolder", serviceName, -1)
		}

		files.EnsurePath(filepath.Dir(templatePath))

		logger.Info("Replaced template path: " + templatePath)

		files.CreateFromTemplate(templatePath, template, cfg)
	}

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

	os.Chdir(config.GetRootDirectory())
	os.Chdir("./services/terraform")

	box = packr.NewBox("./terraformTemplates")
	for _, templatePath := range box.List() {
		logger.Info(templatePath)
		files.EnsurePath(filepath.Dir(templatePath))

		if strings.HasSuffix(templatePath, ".zip") {
			fileContents, _ := box.Find(templatePath)
			files.CreateFromBinary(templatePath, fileContents)
		} else {
			template, _ := box.FindString(templatePath)
			files.CreateFromTemplate(templatePath, template, cfg)
		}
	}

	os.Rename("./lambda.tf", "./lambda_"+serviceName+".tf")

	os.Chdir(config.GetRootDirectory())
}

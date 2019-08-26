package add

import (
	"strconv"
	"strings"

	"github.com/gobuffalo/packr"

	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
)

// Environment adds a new environment to the application
func Environment(environmentName string) {
	config.AppendToCommaSeparatedValue("environments", environmentName)
	createEnvironmentDirectory(environmentName)
	cfg := createAddEnvironmentConfig(environmentName)
	createBaseTerraformTemplates(cfg)
	createLambdaTerraformTemplates(cfg)
}

func createEnvironmentDirectory(environmentName string) {
	environmentDirectory := str.Concat("terraform/", environmentName)
	directories.Create(environmentDirectory)
}

func createAddEnvironmentConfig(environmentName string) map[string]string {
	cfg := make(map[string]string)
	cfg["env"] = config.Get("env")
	cfg["endUserIP"] = config.Get("endUserIP")
	cfg["environmentAbbr"] = environmentName[0:3]
	cfg["environmentName"] = environmentName
	cfg["hostedZone"] = config.Get("hostedZone")
	cfg["projectFqdn"] = config.Get("projectFqdn")
	cfg["projectName"] = config.Get("projectName")
	cfg["region"] = config.Get("region")
	environmentCount := strconv.Itoa(len(strings.Split(config.Get("environments"), ",")))
	cfg["vpcCidrBlock"] = str.Concat("10.", environmentCount, ".0.0/16")
	return cfg
}

func createBaseTerraformTemplates(cfg map[string]string) {
	options := files.TemplateOptions{
		Box:             packr.NewBox("./environment/environmentBaseTemplates"),
		TargetDirectory: str.Concat("terraform/", cfg["environmentName"]),
		Config:          cfg,
	}
	err := files.CreateTemplatedFiles(options)
	errors.QuitIfError(err)
}

func createLambdaTerraformTemplates(cfg map[string]string) {
	allServiceNames := config.Get("services")
	// If services exist,
	if len(allServiceNames) > 0 {
		serviceNames := strings.Split(allServiceNames, ",")
		// For each service, stamp out the Lambda templates box in the new environment
		for _, serviceName := range serviceNames {
			cfg["serviceName"] = serviceName
			options := files.TemplateOptions{
				Box:             packr.NewBox("./environment/environmentLambdaTemplates"),
				TargetDirectory: str.Concat("terraform/", cfg["environmentName"]),
				Config:          cfg,
			}
			err := files.CreateTemplatedFiles(options)
			errors.QuitIfError(err)
		}
	}
}

package setup

import (
	"os"

	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/aws/elbv2"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
)

func ValidateInputs(projectName, frontEnd, backEnd, database, env string) {
	validateProjectName(projectName)
	// validateDockerAndPorts()
	validateStack(frontEnd, backEnd, database)
	validateALBUniqueName(projectName)
	validateEnv(env)
	// TODO: validateGitHubRepoUniqueName
	// TODO: validateHostedZoneIsRegistered
}

// validateEnv checks the given env value and checks it against a list of valid values.
// If the value is not in the list validation fails.
func validateEnv(env string) {
	if env != "dev" {
		if env != "integration" {
			logger.Err("Project environment must be either 'dev' or 'integration'")
			os.Exit(2)
		}
	}
}

func validateProjectName(projectName string) {
	if !str.IsAllLowercaseCharacters(projectName) {
		logger.Err("Project names must be lowercase with no special characters")
		os.Exit(2)
	}
}

func validateDockerAndPorts() {
	_, stderr := commands.Run("docker run --name pac-porttest -p 3000:3000 -p 8001:8000 hello-world", "")
	if stderr != nil {
		logger.Err("Docker must be installed and ports 3000 and 8001 must be open in order to run PAC setup")
		errors.QuitIfError(stderr)
	} else {
		_, stderr := commands.Run("docker rm pac-porttest", "")
		errors.QuitIfError(stderr)
	}
}

func validateStack(frontEnd string, backEnd string, database string) {
	if frontEnd != "ReactJS" || backEnd != "Express" || database != "DynamoDB" {
		errors.LogAndQuit("PAC only supports the following stack at this time: ReactJS, Express, DynamoDB")
	}
}

func validateALBUniqueName(projectName string) {
	region := "{{ .region }}"
	awsSession := aws.CreateAwsSession(region)
	if elbv2.Exists(str.Concat("pac-", projectName, "-i"), awsSession) {
		errors.LogAndQuit("The project name supplied matches resources already provisioned in AWS")
	}
}

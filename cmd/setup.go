package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/PyramidSystemsInc/go/aws/sts"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/cmd/setup"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup new project templates",
	Long: `Generate new project templates with PAC (The default stack is a ReactJS front-end,
NodeJS/Express back-end, and DynamoDB database)`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetLogLevel("info")

		// Get the values provided by command line flags -OR- the default values if not provided
		projectName := getProjectName(cmd)
		description := getDescription(cmd)
		frontEnd := getFrontEnd(cmd)
		backEnd := getBackEnd(cmd)
		database := getDatabase(cmd)
		env := getEnv(cmd)
		awsRegion := getAWSRegion(cmd)
		warnExtraArgumentsAreIgnored(args)

		// Perform various checks to ensure we should proceed
		setup.ValidateInputs(projectName, frontEnd, backEnd, database, env)

		// Create project directory
		setup.CreateRootProjectDirectory(projectName)

		// Copy configuration file to project directory
		input, err := ioutil.ReadFile("C:\\Users\\MMcNairy\\go\\src\\github.com\\PyramidSystemsInc\\pac\\.pac.json")
		if err != nil {
			errors.QuitIfError(err)

			return
		}

		config.GoToRootProjectDirectory(projectName)

		destinationFile := ".pac.json"

		err = ioutil.WriteFile(destinationFile, input, 0644)

		if err != nil {
			fmt.Println("Error creating", destinationFile)
			errors.QuitIfError(err)

			return
		}

		// Create encryption key (used to secure Terraform state) which is needed for the Terraform templates
		encryptionKeyID := setup.CreateEncryptionKey(projectName)

		// Get AWS account ID, used to form some ARNs in Terraform
		awsAccountID := sts.GetAccountID()

		// Set configuration values in the .pac file in the new project directory
		config.Set("projectName", projectName)
		config.Set("description", description)
		config.Set("region", awsRegion)
		config.Set("awsID", awsAccountID)
		config.Set("encryptionKeyID", encryptionKeyID)
		config.Set("gitAuth", gitAuth)
		config.Set("env", env)

		// Copy template files from ./cmd/setup/templates over to the new project
		setup.Templates()

		// Create a GitHub repository for the project
		setup.GitRepository()
	},
}

func init() {
	RootCmd.AddCommand(setupCmd)
	setupCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "project name (required)")
	setupCmd.MarkPersistentFlagRequired("name")
	setupCmd.PersistentFlags().StringVar(&description, "description", "Project created by PAC", "short description of the project")
	setupCmd.PersistentFlags().StringVarP(&frontEnd, "front", "f", "ReactJS", "front-end framework/library")
	setupCmd.PersistentFlags().StringVarP(&backEnd, "back", "b", "Express", "back-end framework/library")
	setupCmd.PersistentFlags().StringVarP(&database, "database", "d", "DynamoDB", "database type")
	setupCmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "environment name")
	setupCmd.PersistentFlags().StringVarP(&awsRegion, "awsregion", "w", "us-east-2", "AWS Region")
}

// TODO: pull from systems manager parameter store
var gitAuth = "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"

func warnExtraArgumentsAreIgnored(args []string) {
	if len(args) > 0 {
		logger.Warn("Arguments were provided, but all arguments after 'setup' and before the flags are ignored")
	}
}

var projectName string

func getProjectName(cmd *cobra.Command) string {
	projectName, err := cmd.Flags().GetString("name")
	errors.QuitIfError(err)
	return projectName
}

var description string

func getDescription(cmd *cobra.Command) string {
	description, err := cmd.Flags().GetString("description")
	errors.QuitIfError(err)
	return description
}

var frontEnd string

func getFrontEnd(cmd *cobra.Command) string {
	frontEnd, err := cmd.Flags().GetString("front")
	errors.QuitIfError(err)
	return frontEnd
}

var backEnd string

func getBackEnd(cmd *cobra.Command) string {
	backEnd, err := cmd.Flags().GetString("back")
	errors.QuitIfError(err)
	return backEnd
}

var database string

func getDatabase(cmd *cobra.Command) string {
	database, err := cmd.Flags().GetString("database")
	errors.QuitIfError(err)
	return database
}

var env string

func getEnv(cmd *cobra.Command) string {
	env, err := cmd.Flags().GetString("env")
	errors.QuitIfError(err)
	return env
}

var awsRegion string

func getAWSRegion(cmd *cobra.Command) string {
	awsRegion, err := cmd.Flags().GetString("awsregion")
	errors.QuitIfError(err)
	return awsRegion
}

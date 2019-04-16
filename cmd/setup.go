package cmd

import (
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
  "github.com/PyramidSystemsInc/go/terraform"
  "github.com/PyramidSystemsInc/pac/cmd/setup"
  "github.com/PyramidSystemsInc/pac/config"
  "github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
  Use:   "setup",
  Short: "Setup a new project",
  Long: `Generate a new project with PAC (The default stack is a ReactJS front-end,
NodeJS/Express back-end, and DynamoDB database)`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    awsRegion := "us-east-2"

    // Get the values provided by command line flags -OR- the default values if not provided
    projectName := getProjectName(cmd)
    description := getDescription(cmd)
    frontEnd := getFrontEnd(cmd)
    backEnd := getBackEnd(cmd)
    database := getDatabase(cmd)
    warnExtraArgumentsAreIgnored(args)
    setup.ValidateInputs(projectName, frontEnd, backEnd, database)

    // Perform various checks to ensure we should proceed
    terraform.VerifyInstallation()

    // Set environment variables
    setup.SetEnvironmentVariables(projectName)

    // Create an encrypted S3 bucket where Terraform can store state
    projectFqdn, encryptionKeyID := setup.TerraformS3Bucket(projectName)

    // Copy template files from ./cmd/setup/templates over to the new project's directory
    setup.Templates(projectName, description, gitAuth, awsRegion, encryptionKeyID)

    // Call on Terraform to create the AWS infrastructure
    setup.Infrastructure()

    // Creates a GitHub repository and sets up a webhook to queue a Jenkins build every time a push is made to GitHub
    jenkinsUrl := str.Concat("jenkins.", projectFqdn, ":8080")
    setup.GitRepository(projectName)
    setup.GitHubWebhook(projectName, gitAuth, jenkinsUrl)

    // Set configuration values in the .pac file in the new project directory
    config.Set("encryptionKeyID", encryptionKeyID)
    config.Set("jenkinsUrl", jenkinsUrl)
    config.Set("projectFqdn", projectFqdn)

    // TODO: Add as a setup step
    // setup.AutomateJenkins()
  },
}

func init() {
  RootCmd.AddCommand(setupCmd)
  setupCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "project name (required)")
  setupCmd.MarkPersistentFlagRequired("name")
  setupCmd.PersistentFlags().StringVar(&description, "description", "Project created by PAC", "short description of the project")
  setupCmd.PersistentFlags().StringVar(&hostedZone, "hostedZone", "pac.pyramidchallenges.com", "Existing AWS hosted zone FQDN (i.e. pac.pyramidchallenges.com)")
  setupCmd.PersistentFlags().StringVarP(&frontEnd, "front", "f", "ReactJS", "front-end framework/library")
  setupCmd.PersistentFlags().StringVarP(&backEnd, "back", "b", "Express", "back-end framework/library")
  setupCmd.PersistentFlags().StringVarP(&database, "database", "d", "DynamoDB", "database type")
}

func warnExtraArgumentsAreIgnored(args []string) {
  if len(args) > 0 {
    logger.Warn("Arguments were provided, but all arguments after 'setup' and before the flags are ignored")
  }
}

// TODO: pull from systems manager parameter store
var gitAuth = "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"

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

var hostedZone string

func getHostedZone(cmd *cobra.Command) string {
  hostedZone, err := cmd.Flags().GetString("hostedZone")
  errors.QuitIfError(err)
  return hostedZone
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

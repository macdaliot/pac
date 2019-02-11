package cmd

import (
  "github.com/PyramidSystemsInc/pac/cmd/setup"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var projectDirectory string

var setupCmd = &cobra.Command{
  Use:   "setup",
  Short: "Setup a new project",
  Long: `Generate a new project with PAC (The default stack is a ReactJS front-end,
NodeJS/Express back-end, and DynamoDB database)`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    projectName := getProjectName(cmd)
    description := getDescription(cmd)
    hostedZone := getHostedZone(cmd)
    frontEnd := getFrontEnd(cmd)
    backEnd := getBackEnd(cmd)
    database := getDatabase(cmd)
    warnExtraArgumentsAreIgnored(args)
    setup.ValidateInputs(projectName, frontEnd, backEnd, database)
    projectDirectory := setup.ProjectStructure(projectName, description, gitAuth)
    projectFqdn := setup.Route53HostedZone(projectName, hostedZone)
    setup.S3Buckets(projectName, projectFqdn)
    setup.ElasticLoadBalancer(projectName, projectFqdn)
    setup.Jenkins(projectName, projectFqdn)
    setup.SonarQube(projectName, projectFqdn)
    setup.Selenium(projectName, projectFqdn)
    setup.FrontEndFiles(projectDirectory, projectName, description)
    setup.HaProxy(projectDirectory, projectName)
    setup.GitRepository(projectName, gitAuth, projectDirectory)
    setup.GitHubWebhook(projectName)
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

var gitAuth string = "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"

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

var path string

package cmd

import (
  "github.com/PyramidSystemsInc/pac/cmd/create"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var projectDirectory string

var createCmd = &cobra.Command{
  Use:   "create",
  Short: "Create a new project",
  Long: `Generate a new project with PAC (The default stack is a ReactJS front-end,
NodeJS/Express back-end, and DynamoDB database)`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    warnArgumentsAreIgnored(args)
    validateStack(cmd)
    projectName := getProjectName(cmd)
    description := getDescription(cmd)
    gitUser := getGitUser(cmd)
    gitPass := getGitPass(cmd)
    projectDirectory := create.ProjectStructure(projectName, description)
    create.FrontEndFiles(projectDirectory, projectName, description)
    create.DynamoDb(projectName)
    create.GitRepository(projectName, gitUser, gitPass, projectDirectory)
  },
}

func init() {
  RootCmd.AddCommand(createCmd)
  createCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "project name (required)")
  createCmd.MarkFlagRequired("name")
  createCmd.PersistentFlags().StringVar(&description, "description", "Project created by PAC", "short description of the project")
  createCmd.PersistentFlags().StringVar(&path, "path", ".", "location where to create the project (Note: PAC creates the project's folder)")
  createCmd.PersistentFlags().StringVarP(&frontEnd, "front", "f", "ReactJS", "front-end framework/library")
  createCmd.PersistentFlags().StringVarP(&backEnd, "back", "b", "Express", "back-end framework/library")
  createCmd.PersistentFlags().StringVarP(&database, "database", "d", "DynamoDB", "database type")
  createCmd.PersistentFlags().StringVar(&gitUser, "gitUser", "", "GitHub user name (must be used with --gitPass)")
  createCmd.PersistentFlags().StringVar(&gitPass, "gitPass", "", "GitHub password (must be used with --gitUser)")
  //createCmd.PersistentFlags().StringVarP(&pages, "pages", "p", "", "pages to be created on the front-end")
  //createCmd.PersistentFlags().StringVarP(&services, "services", "s", "", "services to be created on the back-end")
}

func warnArgumentsAreIgnored(args []string) {
  if len(args) > 0 {
    logger.Warn("Arguments were provided, but all arguments after 'create' and before the flags are ignored")
  }
}

func validateStack(cmd *cobra.Command) {
  frontEnd := getFrontEnd(cmd)
  backEnd := getBackEnd(cmd)
  database := getDatabase(cmd)
  if frontEnd != "ReactJS" || backEnd != "Express" || database != "DynamoDB" {
    errors.LogAndQuit("Sorry, PAC only supports the following stack at this time: ReactJS, Express, DynamoDB. Check back soon")
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

var gitUser string

func getGitUser(cmd *cobra.Command) string {
  gitUser, err := cmd.Flags().GetString("gitUser")
  errors.QuitIfError(err)
  return gitUser
}

var gitPass string

func getGitPass(cmd *cobra.Command) string {
  gitPass, err := cmd.Flags().GetString("gitPass")
  errors.QuitIfError(err)
  return gitPass
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

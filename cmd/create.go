package cmd

import (
  "github.com/PyramidSystemsInc/pac/cmd/create"
  "github.com/PyramidSystemsInc/go/commands"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var projectDirectory string

var createCmd = &cobra.Command{
  Use:   "create",
  Short: "Create a new project",
  Long: `Generate a new project with PAC (The default stack is a ReactJS front-end,
NodeJS/Express back-end, and MongoDB database)`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    warnArgumentsAreIgnored(args)
    validateStack(cmd)
    projectName := getProjectName(cmd)
    description := getDescription(cmd)
    projectDirectory := createAllDirectories(projectName)
    create.FrontEndFiles(projectDirectory, projectName, description)
    commands.Run("docker run -d --name pac-" + projectName + "-db -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo", "")
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
  createCmd.PersistentFlags().StringVarP(&database, "database", "d", "MongoDB", "database type")
  createCmd.PersistentFlags().StringVar(&gitUserName, "gitUser", "", "GitHub user name (must be used with --gitPass)")
  createCmd.PersistentFlags().StringVar(&gitPassword, "gitPass", "", "GitHub password (must be used with --gitUser)")
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
  if frontEnd != "ReactJS" || backEnd != "Express" || database != "MongoDB" {
    errors.LogAndQuit("Sorry, PAC only supports the following stack at this time: ReactJS, Express, MongoDB. Check back soon")
  }
}

func createAllDirectories(projectName string) string {
  projectDirectory := createProjectDirectory(projectName)
  directories.Create(projectDirectory + "/app/src/assets/png")
  directories.Create(projectDirectory + "/app/src/components/Header")
  directories.Create(projectDirectory + "/app/src/components/Sidebar/parts/Button")
  directories.Create(projectDirectory + "/app/src/components/pages/NotFound")
  directories.Create(projectDirectory + "/app/src/routes")
  directories.Create(projectDirectory + "/app/src/scss")
  directories.Create(projectDirectory + "/svc")
  return projectDirectory
}

func createProjectDirectory(projectName string) string {
  workingDirectory := directories.GetWorking()
  projectDirectory := workingDirectory + "/" + projectName
  directories.Create(projectDirectory)
  return projectDirectory
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

var gitUserName string
var gitPassword string
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

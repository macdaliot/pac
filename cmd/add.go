package cmd

import (
  "os"

  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/pac/cmd/add"
  "github.com/PyramidSystemsInc/pac/config"
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add to an existing project",
  Long:  `Generate skeleton code for a new front-end page or back-end service`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    addType := validateAddTypeArgument(args)
    os.Chdir(config.GetRootDirectory())
    addFiles(cmd, addType)
  },
}

func init() {
  RootCmd.AddCommand(addCmd)
  addCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the new service (required)")
}

func validateAddTypeArgument(args []string) string {
  if len(args) == 1 {
    if args[0] == "service" || args[0] == "authService" {
      return args[0]
    } else {
      errors.LogAndQuit("The type was set to an invalid value. The valid types are 'service' and 'authService'")
    }
  } else if len(args) == 0 {
    errors.LogAndQuit("A type must be specifed after the 'add' command. The valid types are 'service' and 'authService'")
  } else if len(args) > 1 {
    errors.LogAndQuit("Only one type may be passed for each 'add' command")
  }
  return ""
}

func addFiles(cmd *cobra.Command, addType string) {
  if addType == "service" {
    addService(cmd)
  } else if addType == "authService" {
    addAuthService(cmd)
  } else {
    errors.LogAndQuit("The type was set to an invalid value. The valid types are 'service' and 'authService'")
  }
}

func addService(cmd *cobra.Command) {
  name := getName(cmd)
  add.Service(name)
}

func addAuthService(cmd *cobra.Command) {
  add.AuthService()
}

var name string

func getName(cmd *cobra.Command) string {
  name, err := cmd.Flags().GetString("name")
  errors.QuitIfError(err)
  return name
}

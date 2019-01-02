package cmd

import (
  "github.com/PyramidSystemsInc/pac/cmd/automate"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var automateCmd = &cobra.Command{
  Use:   "automate",
  Short: "Instruct pac to handle an aspect of the project by itself",
  Long: `Instruct pac to handle an aspect of the project by itself`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    validateAutomateTypeArgument(args)
    automate.Jenkins()
  },
}

func init() {
  RootCmd.AddCommand(automateCmd)
}

func validateAutomateTypeArgument(args []string) {
  if len(args) == 1 {
    if args[0] != "jenkins" {
      errors.LogAndQuit("The type was set to an invalid value. The valid types are 'jenkins'")
    }
  } else if len(args) == 0 {
    errors.LogAndQuit("A type must be specifed after the 'automate' command. The valid types are 'jenkins'")
  } else if len(args) > 1 {
    errors.LogAndQuit("Only one type may be passed for each 'automate' command")
  }
}

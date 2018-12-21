package cmd

import (
  "fmt"
  "github.com/PyramidSystemsInc/pac/cmd/add"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add to an existing project",
  Long: `Generate skeleton code for a new front-end page or back-end service`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    addType := validateAddTypeArgument(args)
    addFiles(cmd, addType)
  },
}

func init() {
  RootCmd.AddCommand(addCmd)
  addCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the new page/service (required)")
  addCmd.PersistentFlags().StringVar(&color, "color", "#4285f4", "either hex code or CSS color for a new page (only for 'page' type)")
  addCmd.PersistentFlags().StringVar(&icon, "icon", "build", "either FontAwesome class or Material icon name for a new page (only for 'page' type)")
  addCmd.PersistentFlags().BoolVar(&dividerAfter, "dividerAfter", false, "adds a divider after the sidebar button (only for 'page' type)")
  addCmd.PersistentFlags().StringVar(&uri, "uri", "/<page-name>", "URI for a new page (only for 'page' type)")
  addCmd.PersistentFlags().BoolVar(&getMethod, "get", false, "creates an HTTP GET method stub (only for 'service' type)")
  addCmd.PersistentFlags().BoolVar(&postMethod, "post", false, "creates an HTTP POST method stub (only for 'service' type)")
  addCmd.PersistentFlags().BoolVar(&putMethod, "put", false, "creates an HTTP PUT method stub (only for 'service' type)")
  addCmd.PersistentFlags().BoolVar(&deleteMethod, "delete", false, "creates an HTTP DELETE method stub (only for 'service' type)")
}

func validateAddTypeArgument(args []string) string {
  if len(args) == 1 {
    if args[0] == "page" || args[0] == "service" {
      return args[0]
    } else {
      errors.LogAndQuit("The type was set to an invalid value. The valid types are 'page' and 'service'")
    }
  } else if len(args) == 0 {
    errors.LogAndQuit("A type must be specifed after the 'add' command. The valid types are 'page' and 'service'")
  } else if len(args) > 1 {
    errors.LogAndQuit("Only one type may be passed for each 'add' command")
  }
  return ""
}

func addFiles(cmd *cobra.Command, addType string) {
  if addType == "page" {
    addPage(cmd)
  } else if addType == "service" {
    addService(cmd)
  } else {
    errors.LogAndQuit("The type was set to an invalid value. The valid types are 'page' and 'service'")
  }
}

func addPage(cmd *cobra.Command) {
  fmt.Println(cmd.Flags())
}

func addService(cmd *cobra.Command) {
  name := getName(cmd)
  add.Service(name)
}

var name string

func getName(cmd *cobra.Command) string {
  name, err := cmd.Flags().GetString("name")
  errors.QuitIfError(err)
  return name
}

var color string
var dividerAfter bool
var icon string
var uri string
var getMethod bool
var postMethod bool
var putMethod bool
var deleteMethod bool

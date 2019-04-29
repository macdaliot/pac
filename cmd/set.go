package cmd

import (
  "github.com/PyramidSystemsInc/pac/config"
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
  Use:   "set",
  Short: "Saves a property/value combination in the PAC project configuration",
  Long:  `Saves a property/value combination in the PAC project configuration`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    property := getConfigPropertyNameToSave(cmd)
    value := getConfigValueToSave(cmd)
    config.Set(property, value)
    logger.Info("Configuration value saved")
  },
}

func init() {
  RootCmd.AddCommand(setCmd)
  setCmd.PersistentFlags().StringVarP(&configPropertyNameToSave, "property", "p", "", "Name of the property you want to create/overwrite")
  setCmd.MarkPersistentFlagRequired("property")
  setCmd.PersistentFlags().StringVarP(&configValueToSave, "value", "v", "", "Value you want to create/overwrite")
  setCmd.MarkPersistentFlagRequired("value")
}

var configPropertyNameToSave string

func getConfigPropertyNameToSave(cmd *cobra.Command) string {
  property, err := cmd.Flags().GetString("property")
  errors.QuitIfError(err)
  return property
}

var configValueToSave string

func getConfigValueToSave(cmd *cobra.Command) string {
  value, err := cmd.Flags().GetString("value")
  errors.QuitIfError(err)
  return value
}

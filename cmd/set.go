package cmd

import (
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

const (
	prop string = "property"
	val  string = "value"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Saves a property/value combination in the PAC project configuration",
	Long:  `Saves a property/value combination in the PAC project configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		property := getConfigPropertyNameToSave(cmd)
		value := getConfigValueToSave(cmd)
		config.Set(property, value)
		logger.Info("Configuration value saved")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().StringVarP(&configPropertyNameToSave, prop, "p", "", "Name of the property you want to create/overwrite")
	setCmd.MarkPersistentFlagRequired(prop)
	setCmd.PersistentFlags().StringVarP(&configValueToSave, val, "v", "", "Value you want to create/overwrite")
	setCmd.MarkPersistentFlagRequired(val)
}

var configPropertyNameToSave string

func getConfigPropertyNameToSave(cmd *cobra.Command) string {
	property, err := cmd.Flags().GetString(prop)
	errors.QuitIfError(err)
	return property
}

var configValueToSave string

func getConfigValueToSave(cmd *cobra.Command) string {
	value, err := cmd.Flags().GetString(val)
	errors.QuitIfError(err)
	return value
}

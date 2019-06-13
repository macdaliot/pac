package cmd

import (
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the value from the PAC project configuration",
	Long:  `Retrieves the value from the PAC project configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		property := getConfigPropertyName(cmd)
		if property != "" {
			logger.Info(config.Get(property))
		} else {
			logger.Info(config.Get(args[0]))
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&configPropertyName, "property", "p", "", "Name of the property you want retrieved")
}

var configPropertyName string

func getConfigPropertyName(cmd *cobra.Command) string {
	property, err := cmd.Flags().GetString("property")
	errors.QuitIfError(err)
	return property
}

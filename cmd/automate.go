package cmd

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/cmd/setup"
	"github.com/PyramidSystemsInc/pac/util"
	"github.com/spf13/cobra"
)

var automateCmd = &cobra.Command{
	Use:   "automate",
	Short: "Instruct pac to handle an aspect of the project by itself",
	Long:  `Instruct pac to handle an aspect of the project by itself`,
	Run: func(cmd *cobra.Command, args []string) {
		pipelineInitializerMap := make(map[string]func())
		pipelineInitializerMap["jenkins"] = setup.AutomateJenkins
		pipelineInitializerMap["azure"] = setup.AutomateAzure

		logger.SetLogLevel("info")
		if util.ValidatePipelineType(args) {
			pipelineInitializerMap[args[0]]()
		}
	},
}

func init() {

	RootCmd.AddCommand(automateCmd)
}

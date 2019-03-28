package cmd

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/pac/cmd/cleanCloud"
	"github.com/spf13/cobra"
)

var cleanCloudCmd = &cobra.Command{
	Use:   "clean-cloud",
	Short: "Destroy all resources related to a PAC project",
	Long:  `Destroy all resources related to a PAC project (does not delete Route53 or Cloudfront resources)`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetLogLevel("info")
		cleanCloud.DeleteAllResources()
	},
}

func init() {
	RootCmd.AddCommand(cleanCloudCmd)
}

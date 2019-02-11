package cmd

import (
	"github.com/PyramidSystemsInc/go/errors"
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
    projectName := getProjectToClean(cmd)
    cleanCloud.DeleteAllResources(projectName)
	},
}

func init() {
	RootCmd.AddCommand(cleanCloudCmd)
	cleanCloudCmd.PersistentFlags().StringVarP(&projectToClean, "name", "n", "", "name of the new page/service/stage (required)")
  cleanCloudCmd.MarkPersistentFlagRequired("name")
}

var projectToClean string

func getProjectToClean(cmd *cobra.Command) string {
	name, err := cmd.Flags().GetString("name")
	errors.QuitIfError(err)
	return name
}

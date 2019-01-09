package cmd

import (
  "github.com/PyramidSystemsInc/go/errors"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/pac/cmd/create"
  "github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
  Use:   "create",
  Short: "Create a pipeline",
  Long: `Create a new pipeline in Jenkins`,
  Run: func(cmd *cobra.Command, args []string) {
    logger.SetLogLevel("info")
    validateCreateTypeArgument(args)
    resourceName := getResourceName(cmd)
    branches := getBranches(cmd)
    resourceDescription := getResourceDescription(cmd)
    create.Jenkinsfile(resourceName)
    create.Pipeline(resourceName, branches, resourceDescription)
  },
}

func init() {
  RootCmd.AddCommand(createCmd)
  createCmd.PersistentFlags().StringVarP(&resourceName, "name", "n", "", "pipeline name (required)")
  createCmd.MarkFlagRequired("name")
  createCmd.PersistentFlags().StringVarP(&branches, "branches", "b", "master", "branch names separated by commas, that when pushed to, will trigger a build on this pipeline")
  createCmd.PersistentFlags().StringVar(&jenkinsUrl, "jenkinsUrl", "", "URL for Jenkins (including port number if necessary)")
  createCmd.MarkFlagRequired("jenkinsUrl")
  createCmd.PersistentFlags().StringVar(&resourceDescription, "description", "Pipeline created by PAC", "short description of the pipeline")
}

func validateCreateTypeArgument(args []string) {
  if len(args) == 1 {
    if args[0] != "pipeline" {
      errors.LogAndQuit("The type was set to an invalid value. The valid types are 'pipeline'")
    }
  } else if len(args) == 0 {
    errors.LogAndQuit("A type must be specifed after the 'create' command. The valid types are 'pipeline' (i.e. 'pac create pipeline --name sample'")
  } else if len(args) > 1 {
    errors.LogAndQuit("Only one type may be passed for each 'create' command")
  }
}

var resourceName string

func getResourceName(cmd *cobra.Command) string {
  resourceName, err := cmd.Flags().GetString("name")
  errors.QuitIfError(err)
  return resourceName
}

var branches string

func getBranches(cmd *cobra.Command) string {
  branches, err := cmd.Flags().GetString("branches")
  errors.QuitIfError(err)
  return branches
}

var jenkinsUrl string

func getJenkinsUrl(cmd *cobra.Command) string {
  jenkinsUrl, err := cmd.Flags().GetString("jenkinsUrl")
  errors.QuitIfError(err)
  return jenkinsUrl
}

var resourceDescription string

func getResourceDescription(cmd *cobra.Command) string {
  resourceDescription, err := cmd.Flags().GetString("description")
  errors.QuitIfError(err)
  return resourceDescription
}

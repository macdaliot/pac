package cmd

import (
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/go/terraform"
	"github.com/PyramidSystemsInc/pac/cmd/deploy"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the cloud resources necessary to run the project",
	Long: `Provision cloud resources using Terraform.
This command is to be run after templates are generated with 'pac setup'`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetLogLevel("info")

		// Get the values from the pac config file
		projectName := config.Get("projectName")
		encryptionKeyID := config.Get("encryptionKeyID")
		gitAuth := config.Get("gitAuth")
		warnExtraArgumentsAreIgnored(args)

		// Perform various checks to ensure we should proceed
		terraform.VerifyInstallation()

		// Set environment variables
		deploy.SetEnvironmentVariables(projectName)

		// Create an encrypted S3 bucket where Terraform can store state
		terraformS3Bucket, projectFqdn := deploy.TerraformS3Bucket(projectName, encryptionKeyID)

		// Sets up a webhook to queue a Jenkins build every time a push is made to GitHub
		jenkinsURL := str.Concat("jenkins.", projectFqdn)
		deploy.GitHubWebhook(projectName, gitAuth, jenkinsURL)

		// Call on Terraform to create the infrastructure in the cloud
		deploy.Infrastructure("dns")
		deploy.Infrastructure(".") // management vpc
		deploy.Infrastructure("application")

		// Set configuration values in the .pac file in the new project directory
		config.Set("jenkinsURL", jenkinsURL)
		config.Set("projectFqdn", projectFqdn)
		config.Set("terraformS3Bucket", terraformS3Bucket)
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

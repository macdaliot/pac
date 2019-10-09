package cmd

import (
	"strings"

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
		// Get the values from the pac config file
		projectFqdn := config.Get("projectFqdn")
		projectName := config.Get("projectName")
		gitAuth := config.Get("gitAuth")

		// Perform various checks to ensure we should proceed
		terraform.VerifyInstallation()

		// Set environment variables
		deploy.SetEnvironmentVariables(projectName)

		// Sets up a webhook to queue a Jenkins build every time a push is made to GitHub
		jenkinsURL := strings.Join([]string{"jenkins.", projectFqdn}, "")
		deploy.GitHubWebhook(projectName, gitAuth, jenkinsURL)

		// Call on Terraform to create the infrastructure in the cloud
		// If the hosted zone fqdn is not already hosted by AWS then the it is considered "pristine"
		if config.Get("dnsPristine") == "true" {
			deploy.Infrastructure("dns_pristine")
		} else {
			deploy.Infrastructure("dns")
		}
		deploy.Infrastructure("management")

		// Set configuration values in the .pac file in the new project directory
		config.Set("jenkinsURL", jenkinsURL)
		config.Set("projectFqdn", projectFqdn)
		config.Set("terraformS3Bucket", terraformS3Bucket)
	},
}

func init() {
	RootCmd.AddCommand(deployCmd)
}

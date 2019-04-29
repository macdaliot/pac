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
    warnExtraArgumentsAreIgnored(args)

    // Perform various checks to ensure we should proceed
    terraform.VerifyInstallation()

    // Set environment variables
    deploy.SetEnvironmentVariables(projectName)

    // Create an encrypted S3 bucket where Terraform can store state
    terraformS3Bucket, projectFqdn := deploy.TerraformS3Bucket(projectName, encryptionKeyID)

    // Creates a GitHub repository and sets up a webhook to queue a Jenkins build every time a push is made to GitHub
    jenkinsUrl := str.Concat("jenkins.", projectFqdn, ":8080")
    deploy.GitRepository(projectName, gitAuth)
    deploy.GitHubWebhook(projectName, gitAuth, jenkinsUrl)

    // Call on Terraform to create the infrastructure in the cloud
    deploy.Infrastructure()

    // Set configuration values in the .pac file in the new project directory
    config.Set("gitAuth", gitAuth)
    config.Set("jenkinsUrl", jenkinsUrl)
    config.Set("projectFqdn", projectFqdn)
    config.Set("terraformS3Bucket", terraformS3Bucket)
  },
}

func init() {
  RootCmd.AddCommand(deployCmd)
}

// TODO: pull from systems manager parameter store
var gitAuth = "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"

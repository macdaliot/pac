package cmd

import (
	"github.com/PyramidSystemsInc/go/aws/sts"
	"github.com/PyramidSystemsInc/go/aws/util"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/pac/cmd/setup"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/PyramidSystemsInc/pac/go/terraform"
	"github.com/spf13/cobra"
)

var supportedFlags = []string{
	"back",
	"database",
	"description",
	"pristine",
	"env",
	"front",
	"name",
	"awsregion",
	"hostedzone"}

// Used to store values for the supported flags
var supportedFlagValues = make(map[string]string)

// Used to convert the supported flags into the flag names used for the config file
var flagToConfig = map[string]string{
	"back":       "backEnd",
	"pristine":   "dnsPristine",
	"front":      "frontEnd",
	"name":       "projectName",
	"awsregion":  "region",
	"hostedzone": "hostedZone",
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup new project templates",
	Long: `Generate new project templates with PAC (The default stack is a ReactJS front-end,
	NodeJS/Express back-end, and DynamoDB database)`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the values provided by command line flags -OR- the default values if not provided
		for _, val := range supportedFlags {
			supportedFlagValues[val] = getFlagValue(cmd, val)
		}

		// Perform various checks to ensure we should proceed
		setup.ValidateInputs(supportedFlagValues["name"], supportedFlagValues["front"], supportedFlagValues["back"], supportedFlagValues["database"], supportedFlagValues["env"])

		setup.CreateRootProjectDirectory(supportedFlagValues["name"])
		setup.GoToRootProjectDirectory(supportedFlagValues["name"])
		config.Create()

		// Set configuration values in the .pac.json file in the new project directory
		for currFlag, val := range supportedFlagValues {
			if configKey, exists := flagToConfig[currFlag]; exists {
				config.Set(configKey, val)
			} else {
				config.Set(currFlag, val)
			}
		}

		config.Set("gitAuth", gitAuth)
		config.Set("terraformAWSVersion", terraform.AWSVersion)
		config.Set("terraformTemplateVersion", terraform.TemplateVersion)

		// Create encryption key (used to secure Terraform state) which is needed for the Terraform templates
		encryptionKeyID := setup.CreateEncryptionKey()
		config.Set("encryptionKeyID", encryptionKeyID)

		// Read AWS Account ID from System Manager
		awsAccountID := sts.GetAccountID()
		config.Set("awsID", awsAccountID)

		// Find the first available VPC CIDR blocks and save them to the configuration
		freeVpcCidrBlocks := setup.FindAvailableVpcCidrBlocks()
		config.Set("awsManagementVpcCidrBlock", freeVpcCidrBlocks[0])
		config.Set("awsApplicationVpcCidrBlock", freeVpcCidrBlocks[1])

		// Get the public address of the end-user for security groups to give access to network resources to user
		endUserIP := util.GetPublicIP()
		config.Set("endUserIP", endUserIP)

		// Copy template files from ./cmd/setup/templates over to the new project
		// This needs to happen before CopyBinaries so it has somewhere to copy to
		setup.Templates()

		// Create a GitHub repository for the project
		setup.GitRepository()
	},
}

func init() {
	RootCmd.AddCommand(setupCmd)
	setupCmd.PersistentFlags().StringP("name", "n", "", "project name (required)")
	setupCmd.MarkPersistentFlagRequired("name")
	setupCmd.PersistentFlags().StringP("description", "i", "Project created by PAC", "short description of the project")
	setupCmd.PersistentFlags().StringP("front", "f", "ReactJS", "front-end framework/library")
	setupCmd.PersistentFlags().StringP("back", "b", "Express", "back-end framework/library")
	setupCmd.PersistentFlags().StringP("database", "d", "DynamoDB", "database type")
	setupCmd.PersistentFlags().StringP("env", "e", "dev", "environment name")
	setupCmd.PersistentFlags().StringP("awsregion", "w", "us-east-2", "AWS Region")
	setupCmd.PersistentFlags().StringP("pristine", "p", "false", "Hosted zone doesn't already exist.")
	setupCmd.PersistentFlags().StringP("hostedzone", "z", "pac.pyramidchallenges.com", "DNS zone to add records to.")
}

// TODO: pull from systems manager parameter store
var gitAuth = "amRpZWRlcmlrc0Bwc2ktaXQuY29tOkRpZWRyZV4yMDE4"

var flagValue string

func getFlagValue(cmd *cobra.Command, flagString string) string {
	flagValue, err := cmd.Flags().GetString(flagString)
	errors.QuitIfError(err)
	return flagValue
}

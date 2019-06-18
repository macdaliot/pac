package cmd

import (
	"os"
	"strings"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/pac/cmd/add"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

var addResourceMap = map[string]func(*cobra.Command){
	"service" : addService,
	"authService" : addAuthService,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add to an existing project",
	Long:  `Generate skeleton code for a new front-end page or back-end service`,
	Run: func(cmd *cobra.Command, args []string) {
		addType := validateAddTypeArgument(args)
		os.Chdir(config.GetRootDirectory())
		addResourceMap[addType](cmd)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the new service (required)")
}

func validateAddTypeArgument(args []string) string {
	if len(args) == 1 {
		if _,ok := addResourceMap[args[0]]; ok {
			return args[0]
		} else {
			errors.LogAndQuit("The type was set to an invalid value. The valid types are " + getSupportedFlags(addResourceMap))
		}
	} else if len(args) == 0 {
		errors.LogAndQuit("A type must be specified after the 'add' command. The valid types are " + getSupportedFlags(addResourceMap))
	} else if len(args) > 1 {
		errors.LogAndQuit("Only one type may be passed for each 'add' command")
	}
	return ""
}

func addService(cmd *cobra.Command) {
	name := getName(cmd)
	add.Service(name)
}

func addAuthService(cmd *cobra.Command) {
	add.AuthService()
}

var name string

func getName(cmd *cobra.Command) string {
	name, err := cmd.Flags().GetString("name")
	errors.QuitIfError(err)
	return name
}

func getSupportedFlags(m map[string]func(*cobra.Command)) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return "'" + strings.Join(keys, "', '") + "'"
}

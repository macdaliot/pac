package cmd

import (
	"strings"

	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/util"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs all project dependencies",
	Long: `Finds all package.json files within a project and runs 'npm install'
on each directory`,
	Run: func(cmd *cobra.Command, args []string) {
		// Run `npm install` for each Node project directory within the PAC project
		util.GoToRootDirectory()
		packageJsonLocations, err := commands.Run("find . -name package.json", "")
		errors.QuitIfError(err)
		str.RunForEachLine(performNpmInstall, packageJsonLocations)
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
}

func performNpmInstall(packageJsonLocation string) {
	if !strings.Contains(packageJsonLocation, "node_modules") {
		subProjectDirectory := strings.TrimSuffix(packageJsonLocation, "/package.json")
		_, err := commands.Run("npm i", subProjectDirectory)
		errors.QuitIfError(err)
		logger.Info(str.Concat("Installed node modules at ", subProjectDirectory))
	}
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "pac",
	Short: "Pyramid Application Constructor",
	Long: `The Pyramid App Constructor (PAC) is a toolkit to help jumpstart the
application development process, specifically designed for compressed
time-duration events like hackathons and tech challenges. PAC will generate
scaffolding composed of reusable components, templates, and pipelines to help
accelerate development velocity, while ensuring security and quality discipline,
to achieve acceptable software hygiene. PAC is an evolving toolkit, and
currently supports a slightly adjusted MERN stack (DynamoDB, Express, React,
Node). It leverages Jenkins for pipelines, Auth0 for authentication, AWS as the
cloud platform, and is supported by relevant open source libraries`,
	Version: "1.0",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logLevel := strings.ToUpper(getLogLevel(cmd))

		input, validInput := logger.ParseLevel(logLevel)

		if validInput {
			logger.SetLogLevel(input)
		} else if logLevel == "" {
			logger.SetLogLevel(logger.INFO)
		} else {
			logger.SetLogLevel(logger.INFO)
			logger.Warn("Invalid log level input. Valid inputs are: " + getValidLogFlags(logger.Levels) + ". Defaulting to 'info' log level")
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "", "log level (info, warning, error)")

	// cobra.OnInitialize(initConfig)
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pac.yaml)")
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".pac.json")   // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
*/

var logLevel string

func getLogLevel(cmd *cobra.Command) string {
	logLevel, err := cmd.Flags().GetString("loglevel")
	errors.QuitIfError(err)
	return logLevel
}

func getValidLogFlags(m map[logger.LogLevel]string) string {
	values := make([]string, 0)
	for _, v := range m {
		values = append(values, strings.ToLower(v))
	}
	return "'" + strings.Join(values, "', '") + "'"
}

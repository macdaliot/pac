package cmd

import (
	"io/ioutil"
	"github.com/PyramidSystemsInc/go/aws"
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/docker"
	"github.com/PyramidSystemsInc/go/errors"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
	"github.com/PyramidSystemsInc/pac/config"
	"github.com/spf13/cobra"
)

// TODO: Attempting to run a container that is already running does not fail, but fails silently (currently the quickest way to get this done)

var runCmd = &cobra.Command{
	Use: "run",
	Short: "Runs the front-end and the micro-services of a PAC project, locally",
	Long: `Runs the front-end and the micro-services of a PAC project, locally`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetLogLevel("info")
		projectName := findProjectName()
		createDockerNetworkIfNeeded(projectName)
		docker.CleanContainers("pac-" + projectName + "-db-local")
		runDatabaseContainer(projectName)
		runMicroserviceContainers(projectName)
		runReverseProxyContainer(projectName)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func findProjectName() string {
	projectName := config.Get("projectName")
	const emptyString string = ""
	if projectName == emptyString {
		errors.LogAndQuit("Error finding the project name. Are you inside a PAC project?")
	}
	return projectName
}

func createDockerNetworkIfNeeded(projectName string) {
	networkName := str.Concat("pac-", projectName)
	if !docker.DoesNetworkExist(networkName) {
		docker.CreateNetwork(networkName)
	}
}

func runDatabaseContainer(projectName string) {
	docker.RunContainer("pac-" + projectName + "-db-local", "pac-" + projectName, []int{8001}, []int{8000}, "", "", []string{}, "amazon/dynamodb-local", "-jar DynamoDBLocal.jar -sharedDb")
	logger.Info("The database is running")
}

func runMicroserviceContainers(projectName string) {
	files, err := ioutil.ReadDir(config.GetRootDirectory() + "/services")
	errors.QuitIfError(err)
	for _, file := range files {
		fileName := file.Name()
		if fileName != "terraform" && file.IsDir() {
			runMicroserviceContainer(fileName, projectName)
		}
	}
}

func runMicroserviceContainer(serviceName string, projectName string) {
	awsAccessKey, err := aws.GetAccessKey()
	errors.LogIfError(err)
	awsSecretKey, err := aws.GetSecretKey()
	errors.LogIfError(err)
	serviceMountDirectory := config.GetRootDirectory() + "/services/" + serviceName
  commands.Run("npm run generate:templates", serviceMountDirectory)
  commands.Run("npx tsc", serviceMountDirectory)
	docker.RunContainer("pac-" + projectName + "-" + serviceName, "pac-" + projectName, []int{}, []int{}, serviceMountDirectory + ":/usr/src/app", "/usr/src/app", []string{"AWS_ACCESS_KEY_ID=" + awsAccessKey, "AWS_SECRET_ACCESS_KEY=" + awsSecretKey}, "node:8", "node dist/services/"+ serviceName + "/src/local")
	logger.Info("The " + serviceName + " microservice is running")
}

// docker run --name pac-proxy-local --network pac-$PROJECT_NAME -p $HAPROXY_PORT:$HAPROXY_PORT -v $HAPROXY_CONFIG_PATH:/usr/local/etc/haproxy:ro -d haproxy
func runReverseProxyContainer(projectName string) {
	proxyMountDirectory := config.GetRootDirectory() + "/services"
	docker.RunContainer("pac-" + projectName + "-proxy-local", "pac-" + projectName, []int{3000}, []int{3000}, proxyMountDirectory + ":/usr/local/etc/haproxy:ro", "", []string{}, "haproxy", "")
	logger.Info("The reverse proxy is running")
}

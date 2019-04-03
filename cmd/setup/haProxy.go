package setup

import (
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
)

func HaProxy(projectName string) {
	createDockerNetwork(projectName)
	runHaProxyContainer(projectName)
	logger.Info("Created local proxy for future services in Docker")
}

func createDockerNetwork(projectName string) {
	commands.Run(str.Concat("docker network create pac-", projectName), "")
}

func runHaProxyContainer(projectName string) {
	cmd := str.Concat("docker run --name pac-proxy-local --network pac-", projectName, " -p 3000:3000 -v ", directories.GetWorking(), "/", projectName, "/services:/usr/local/etc/haproxy:ro -d haproxy")
	commands.Run(cmd, "")
}

package setup

import (
	"github.com/PyramidSystemsInc/go/commands"
	"github.com/PyramidSystemsInc/go/directories"
	"github.com/PyramidSystemsInc/go/files"
	"github.com/PyramidSystemsInc/go/logger"
	"github.com/PyramidSystemsInc/go/str"
)

func HaProxy(projectName string) {
	createHaProxyConfig(projectName)
	createDockerNetwork(projectName)
	runHaProxyContainer(projectName)
	logger.Info("Created local proxy for future services in Docker")
}

func createHaProxyConfig(projectName string) {
	config := make(map[string]string)
	config["projectName"] = projectName
	const template = `frontend service_gateway
    bind 0.0.0.0:3000
    mode http
    option tcplog
    timeout client 50000
`
	filePath := "svc/haproxy.cfg"
	files.CreateFromTemplate(filePath, template, config)
}

func createDockerNetwork(projectName string) {
	commands.Run(str.Concat("docker network create pac-", projectName), "")
}

func runHaProxyContainer(projectName string) {
	cmd := str.Concat("docker run --name pac-proxy-local --network pac-", projectName, " -p 3000:3000 -v ", directories.GetWorking(), "/", projectName, "/svc:/usr/local/etc/haproxy:ro -d haproxy")
	commands.Run(cmd, "")
}

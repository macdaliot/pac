package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateAwsSdkConfig creates a awsSdkConfig.js file based on the configuration passed in
func CreateAwsSdkConfigJs(filePath string, config map[string]string) {
	const template = `const cloud = {
  region: 'us-east-2'
};

const local = {
  region: 'local',
  endpoint: 'http://pac-db-local:8000'
};

module.exports = {
  cloud: cloud,
  local: local
};
`
	files.CreateFromTemplate(filePath, template, config)
}

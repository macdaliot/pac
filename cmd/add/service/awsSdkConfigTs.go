package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

// CreateAwsSdkConfig creates a awsSdkConfig.js file based on the configuration passed in
func CreateAwsSdkConfigTs(filePath string) {
	const template = `export const cloud = {
    region: 'us-east-2'
  };
  
  export const local = {
    region: 'local',
    endpoint: 'http://pac-db-local:8000'
  };
`
	files.CreateFromTemplate(filePath, template, nil)
}

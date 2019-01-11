package frontEndFiles

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateJenkinsfile(filePath string, config map[string]string) {
  const template = `pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'cd app; npm i'
        sh 'cd app; npm run build'
      }
    }
    stage('Deploy') {
      steps {
        script {
          sh '''#!/bin/bash
            bucket_found=$(aws s3 ls --region us-east-2 | grep {{.projectName}}.pyramidchallenges.com)
            if [ ${#bucket_found} -eq 0 ]; then
              aws s3 mb s3://{{.projectName}}.pyramidchallenges.com --region us-east-2
              aws s3 website s3://{{.projectName}}.pyramidchallenges.com --index-document index.html --error-document index.html
            fi
            aws s3 sync app/dist s3://{{.projectName}}.pyramidchallenges.com --acl public-read
          '''
        }
      }
    }
    stage('Test') {
      steps {
        echo 'Testing...'
      }
    }
  }
}
`
  files.CreateFromTemplate(filePath, template, config)
}

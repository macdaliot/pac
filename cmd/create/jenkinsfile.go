package create

import (
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
)

func Jenkinsfile(pipelineName string) {
  pacFile := readPacFile()
  config := make(map[string]string)
  config["projectName"] = pacFile.ProjectName
  config["pipelineName"] = pipelineName
  directories.Create("./jenkins/" + pipelineName)
  const template = `pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'cd app; npm i'
        sh 'cd app; npm run build'
      }
    }
    stage('Test') {
      steps {
        echo 'Testing...'
      }
    }
    stage('Integration') {
      steps {
        echo 'Integrating...'
      }
    }
    stage('Inspect') {
      steps {
        echo 'Inspecting...'
      }
    }
    stage('Deploy') {
      steps {
        script {
          sh '''#!/bin/bash
            bucket_found=$(aws s3 ls --region us-east-2 | grep {{.pipelineName}}.{{.projectName}}.pyramidchallenges.com)
            if [ ${#bucket_found} -eq 0 ]; then
              aws s3 mb s3://{{.pipelineName}}.{{.projectName}}.pyramidchallenges.com --region us-east-2
              aws s3 website s3://{{.pipelineName}}.{{.projectName}}.pyramidchallenges.com --index-document index.html --error-document index.html
            fi
            aws s3 sync app/dist s3://{{.pipelineName}}.{{.projectName}}.pyramidchallenges.com --acl public-read
          '''
        }
      }
    }
  }
}
`
  files.CreateFromTemplate("./jenkins/" + pipelineName + "/Jenkinsfile", template, config)
  logger.Info("Created skeleton Jenkinsfile")
}

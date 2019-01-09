package create

import (
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/files"
  "github.com/PyramidSystemsInc/go/logger"
  "github.com/PyramidSystemsInc/go/str"
)

func Jenkinsfile(pipelineName string) {
  pacFile := readPacFile()
  config := make(map[string]string)
  config["projectName"] = pacFile.ProjectName
  config["pipelineName"] = pipelineName
  directories.Create(str.Concat("./jenkins/", pipelineName))
  template := `pipeline {
  agent any
  stages {
    stage('Init') {
      steps {
        echo 'Jenkins Pipeline present.'
      }
    }
  }
}
`
  files.CreateFromTemplate(str.Concat("./jenkins/", pipelineName, "/Jenkinsfile"), template, config)
  template = `pipeline {
  agent any
  stages {
    stage('Init') {
      steps {
        echo 'Jenkins Pipeline present.'
      }
    }
/* NPM INSTALL */
/* WEB APP DEPLOY */
/* WEB APP TEST */
  }
}
`
  files.CreateFromTemplate(str.Concat("./jenkins/", pipelineName, "/.Jenkinsfile"), template, config)
  logger.Info("Created skeleton Jenkinsfile")
}

package stages

import (
  "encoding/json"
  "io/ioutil"
  "regexp"
  "strings"
  "github.com/PyramidSystemsInc/go/directories"
  "github.com/PyramidSystemsInc/go/errors"
)

type PacFile struct {
  ProjectName  string  `json:"projectName"`
  GitAuth      string  `json:"gitAuth"`
  JenkinsUrl   string  `json:"jenkinsUrl"`
}

func AppendDeploy(filePath string) {
  template := `stage('Deploy') {
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
  }`
  jenkinsfileData, err := ioutil.ReadFile(".Jenkinsfile")
  errors.QuitIfError(err)
  template = replaceValuesInTemplate(template)
  newJenkinsfile := strings.Replace(string(jenkinsfileData), "/* WEB APP DEPLOY */", template, 1)
  ioutil.WriteFile(".Jenkinsfile", []byte(newJenkinsfile), 0644)
  commentRegEx := regexp.MustCompile(`\/\*.*\*\/`)
  newJenkinsfile = commentRegEx.ReplaceAllString(newJenkinsfile, "")
  ioutil.WriteFile("Jenkinsfile", []byte(newJenkinsfile), 0644)
}

func replaceValuesInTemplate(template string) string {
  replaceAll := -1
  template = strings.Replace(template, "{{.projectName}}", readPacFile().ProjectName, replaceAll)
  workingDirectory := directories.GetWorking()
  pipelineName := workingDirectory[strings.LastIndex(workingDirectory, "/") + 1:len(workingDirectory)]
  template = strings.Replace(template, "{{.pipelineName}}", pipelineName, replaceAll)
  return template
}

func readPacFile() PacFile {
  var pacFile PacFile
  pacFileData, err := ioutil.ReadFile("../../.pac")
  errors.QuitIfError(err)
  json.Unmarshal(pacFileData, &pacFile)
  return pacFile
}


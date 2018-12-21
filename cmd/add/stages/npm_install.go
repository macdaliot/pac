package stages

import (
  "io/ioutil"
  "regexp"
  "strings"
  "github.com/PyramidSystemsInc/go/errors"
)

func AppendNPMInstall(filePath string) {
  const template = `stage('Build') {
    steps {
      sh 'cd app; npm i'
      sh 'cd app; npm run build'
    }
  }`
  jenkinsfileData, err := ioutil.ReadFile(".Jenkinsfile")
  errors.QuitIfError(err)
  newJenkinsfile := strings.Replace(string(jenkinsfileData), "/* NPM INSTALL */", template, 1)
  ioutil.WriteFile(".Jenkinsfile", []byte(newJenkinsfile), 0644)
  commentRegEx := regexp.MustCompile(`\/\*.*\*\/`)
  newJenkinsfile = commentRegEx.ReplaceAllString(newJenkinsfile, "")
  ioutil.WriteFile("Jenkinsfile", []byte(newJenkinsfile), 0644)
}

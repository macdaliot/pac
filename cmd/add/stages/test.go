package stages

import (
  "io/ioutil"
  "regexp"
  "strings"
  "github.com/PyramidSystemsInc/go/errors"
)

func AppendTest(filePath string) {
  const template = `stage('Test') {
    steps {
      echo 'Testing...'
    }
  }`
  jenkinsfileData, err := ioutil.ReadFile(".Jenkinsfile")
  errors.QuitIfError(err)
  newJenkinsfile := strings.Replace(string(jenkinsfileData), "/* WEB APP TEST */", template, 1)
  ioutil.WriteFile(".Jenkinsfile", []byte(newJenkinsfile), 0644)
  commentRegEx := regexp.MustCompile(`\/\*.*\*\/`)
  newJenkinsfile = commentRegEx.ReplaceAllString(newJenkinsfile, "")
  ioutil.WriteFile("Jenkinsfile", []byte(newJenkinsfile), 0644)
}

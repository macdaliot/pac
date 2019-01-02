package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateJenkinsfile(filePath string) {
	const template = `
JOB = "${env.JOB_NAME}"
pipeline {
  agent any
  stages {
    stage('Magic') {
      steps {
        sh '''#!/bin/bash
          cd svc/'''+JOB+'''
          ./magic.sh
        '''
      }
    }
  }
}
`
	files.CreateFromTemplate(filePath, template, nil)
}

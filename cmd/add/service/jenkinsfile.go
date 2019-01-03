package service

import (
	"github.com/PyramidSystemsInc/go/files"
)

func CreateJenkinsfile(filePath string) {
	const template = `JOB = "${env.JOB_NAME}"
pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''#!/bin/bash
          cd svc/'''+JOB+'''
          ./.build.sh
        '''
      }
    }
    stage('Test') {
      steps {
        echo "Test"
      }
    }
    stage('Integrate') {
      steps {
        echo "Integrate"
      }
    }
    stage('Inspect') {
      steps {
        echo "Inspect"
      }
    }
    stage('Deploy') {
      steps {
        sh '''#!/bin/bash
          cd svc/'''+JOB+'''
          ./.deploy.sh
        '''
      }
    }
  }
}
`
	files.CreateFromTemplate(filePath, template, nil)
}

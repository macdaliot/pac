properties([
  buildDiscarder(
    logRotator(
      daysToKeepStr: '1',
      numToKeepStr: '5'
    )
  )
])

node {
  stage("SCM Checkout") {
    checkout scm
  }
  stage("Build") {
    script {
      sh "go build"
    }
  }
  stage("Integration Tests") {
  }
  stage("Publish Binaries") {
  }
}
